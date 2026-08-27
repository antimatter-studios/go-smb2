package smb2

import (
	"encoding/binary"
	"testing"
)

// The QUERY_DIRECTORY decoders validate a server-supplied length before
// slicing by it. That validation overflows.
//
// IsInvalid computed 64+FileNameLength in uint32, so a FileNameLength
// near the top of the range wrapped the sum to a small number, the
// comparison passed, and the decoder went on to slice a buffer by a
// length it had just failed to reject.
//
// A directory listing against a hostile or man-in-the-middle SMB server
// crashes the process on it. GO-2026-5051.
func TestFileDirectoryInformationDecoderRejectsOverflowingNameLength(t *testing.T) {
	for _, nameLen := range []uint32{
		0xFFFFFFFF, // 64 + this wraps to 63
		0xFFFFFFC0, // wraps to exactly 0
		0xFFFFFFC1, // wraps to 1
		1 << 31,
		0x7FFFFFFF,
	} {
		// A minimal entry: the 64-byte fixed part and no name, claiming
		// a name far larger than the buffer.
		buf := make([]byte, 64)
		binary.LittleEndian.PutUint32(buf[60:64], nameLen)

		d := FileDirectoryInformationDecoder(buf)
		if !d.IsInvalid() {
			t.Errorf("FileNameLength %#x in a %d-byte buffer was accepted as valid",
				nameLen, len(buf))
		}
	}
}

// A well-formed entry must still be accepted — the bound must reject
// what does not fit, not everything.
func TestFileDirectoryInformationDecoderAcceptsAWellFormedEntry(t *testing.T) {
	const name = 8
	buf := make([]byte, 64+name)
	binary.LittleEndian.PutUint32(buf[60:64], name)

	if d := FileDirectoryInformationDecoder(buf); d.IsInvalid() {
		t.Errorf("a %d-byte buffer with a %d-byte name was rejected", len(buf), name)
	}

	// One byte short of what it claims must still be rejected.
	short := make([]byte, 64+name-1)
	binary.LittleEndian.PutUint32(short[60:64], name)
	if d := FileDirectoryInformationDecoder(short); !d.IsInvalid() {
		t.Errorf("a buffer one byte short of its declared name was accepted")
	}
}

// The same arithmetic appears in three other decoders, all reachable
// from a server response.
func TestOtherDecodersRejectOverflowingLengths(t *testing.T) {
	t.Run("SrvRequestResumeKeyResponse", func(t *testing.T) {
		buf := make([]byte, 28)
		binary.LittleEndian.PutUint32(buf[24:28], 0xFFFFFFFF)
		if d := SrvRequestResumeKeyResponseDecoder(buf); !d.IsInvalid() {
			t.Error("an overflowing ContextLength was accepted")
		}
	})

	t.Run("FileQuotaInformation", func(t *testing.T) {
		buf := make([]byte, 40)
		binary.LittleEndian.PutUint32(buf[4:8], 0xFFFFFFFF)
		if d := FileQuotaInformationDecoder(buf); !d.IsInvalid() {
			t.Error("an overflowing SidLength was accepted")
		}
	})
}

// IsInvalid is the guard that makes a buffer safe to read, so it has to be
// safe to call on a buffer of any length. These three measured a
// variable-length field before checking the fixed part was present to measure
// it from, which panics inside the one function whose job is to prevent one.
//
// The resume-key decoder is the reachable case: client.go hands it the ioctl
// output buffer as it arrives from the server, during a server-side copy.
func TestIsInvalidOnTruncatedBuffers(t *testing.T) {
	tests := []struct {
		name      string
		fixed     int
		isInvalid func(b []byte) bool
	}{
		{
			name:      "SrvRequestResumeKeyResponse",
			fixed:     28,
			isInvalid: func(b []byte) bool { return SrvRequestResumeKeyResponseDecoder(b).IsInvalid() },
		},
		{
			name:      "FileDirectoryInformation",
			fixed:     64,
			isInvalid: func(b []byte) bool { return FileDirectoryInformationDecoder(b).IsInvalid() },
		},
		{
			name:      "FileQuotaInformation",
			fixed:     40,
			isInvalid: func(b []byte) bool { return FileQuotaInformationDecoder(b).IsInvalid() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for n := 0; n < tt.fixed; n++ {
				b := make([]byte, n)
				// 0xff so a length field read out of bounds would also be the
				// largest possible value rather than a quiet zero.
				for i := range b {
					b[i] = 0xff
				}
				if !tt.isInvalid(b) {
					t.Errorf("%d bytes accepted, fixed part is %d", n, tt.fixed)
				}
			}

			if tt.isInvalid(make([]byte, tt.fixed)) {
				t.Errorf("%d bytes rejected, which is the whole fixed part", tt.fixed)
			}
		})
	}
}
