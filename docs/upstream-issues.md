# Upstream issue triage

`hirochachacha/go-smb2` has had no commit since July 2022 and 35 open issues.
This fork works through them: fix what can be fixed, say plainly what cannot,
and offer every fix back upstream as a pull request whether or not it is taken.

Several fixes come from [`cloudsoda/go-smb2`](https://github.com/cloudsoda/go-smb2),
the actively maintained fork. It carries the same BSD-2-Clause licence and the
same copyright holder, so porting is a matter of attribution, which each commit
records.

Status values: **fixed** here, **offered** upstream, **blocked** on something we
do not have, **answered** where the issue is a question rather than a defect, and
**wontfix** with a reason.

| Issue | Summary | Status | Notes |
|---|---|---|---|
| [#22](https://github.com/hirochachacha/go-smb2/issues/22) | Access denied against Samba 4.10 | todo | Needs a repro against an old Samba. Possibly the stricter-server compatibility work. |
| [#28](https://github.com/hirochachacha/go-smb2/issues/28) | Expose NTLMSSP target info, even on auth failure | todo | Has an unmerged PR (#29). Useful for scanners. |
| [#36](https://github.com/hirochachacha/go-smb2/issues/36) | Server OS/version detection | todo | NTLM version is available; OS detection is not, and the thread already establishes that. |
| [#40](https://github.com/hirochachacha/go-smb2/issues/40) | `ListSharenames` returns the wrong shares | todo | Root cause is building the UNC path from `RemoteAddr()` rather than the dialed hostname. |
| [#48](https://github.com/hirochachacha/go-smb2/issues/48) | DFS namespace unsupported | todo | Large. Unmerged PR #79 implements it. |
| [#49](https://github.com/hirochachacha/go-smb2/issues/49) | `invalid negotiate flags` | todo | Rejects servers whose NTLM flags are legal but unusual. Hits rclone users. |
| [#50](https://github.com/hirochachacha/go-smb2/issues/50) | `RemoveAll` reports `signing required` | todo | Needs a repro. |
| [#51](https://github.com/hirochachacha/go-smb2/issues/51) | Null session | todo | Same fix as #81. |
| [#52](https://github.com/hirochachacha/go-smb2/issues/52) | Slow transfers | todo | No numbers in the issue. Buffer reuse upstream of us is the likely answer. |
| [#60](https://github.com/hirochachacha/go-smb2/issues/60) | Create fails in a folder with spaces | todo | Reporter's own follow-up points at permissions, not spaces. Needs a test either way. |
| [#62](https://github.com/hirochachacha/go-smb2/issues/62) | SMB3 encryption | todo | Question. Encryption is implemented; answer and close. |
| [#63](https://github.com/hirochachacha/go-smb2/issues/63) | Kerberos authentication | todo | Large. |
| [#66](https://github.com/hirochachacha/go-smb2/issues/66) | Negative counts returned on read/write errors | todo | Violates `io.Reader`/`io.Writer`, panics `bufio`. |
| [#67](https://github.com/hirochachacha/go-smb2/issues/67) | Cannot log in to a macOS SMB server | todo | Testable: we have a Mac. |
| [#68](https://github.com/hirochachacha/go-smb2/issues/68) | Idle connection reset by peer | todo | Needs SMB2 Echo as a keepalive. |
| [#69](https://github.com/hirochachacha/go-smb2/issues/69) | DFS referrals | todo | Same work as #48. |
| [#70](https://github.com/hirochachacha/go-smb2/issues/70) | Banner grabbing | todo | Same ground as #36. |
| [#71](https://github.com/hirochachacha/go-smb2/issues/71) | Empty domain cannot be expressed | todo | Reporter has Wireshark evidence: an extra `\` is always prefixed. |
| [#72](https://github.com/hirochachacha/go-smb2/issues/72) | `Statfs.BlockSize()` is wrong | todo | Sizes come out halved. Reporter names the fix. |
| [#73](https://github.com/hirochachacha/go-smb2/issues/73) | Azure AD joined domain auth | todo | Likely blocked: needs an Azure AD tenant. |
| [#74](https://github.com/hirochachacha/go-smb2/issues/74) | FIPS compliance, MD5 usage | todo | MD5 is mandated by NTLM itself. Probably an answer, not a fix. |
| [#76](https://github.com/hirochachacha/go-smb2/issues/76) | How to donate | wontfix | Not a defect, and not ours to answer. |
| [#77](https://github.com/hirochachacha/go-smb2/issues/77) | Panic in NTLM `client.go` | todo | Same class as #98 and #101. |
| [#81](https://github.com/hirochachacha/go-smb2/issues/81) | Anonymous login | todo | Same fix as #51. |
| [#83](https://github.com/hirochachacha/go-smb2/issues/83) | Server-side library | todo | Scope question, not a defect. |
| [#87](https://github.com/hirochachacha/go-smb2/issues/87) | `Glob` swallows I/O errors | todo | Deliberate in the code, but it hides misconfiguration. |
| [#88](https://github.com/hirochachacha/go-smb2/issues/88) | File watching | todo | SMB2 CHANGE_NOTIFY exists; the library does not expose it. |
| [#89](https://github.com/hirochachacha/go-smb2/issues/89) | `STATUS_PENDING` mishandled | todo | Aborts long transfers. |
| [#90](https://github.com/hirochachacha/go-smb2/issues/90) | Crash in `encrypt` against Azure Files | todo | Same class as #98 and #101. |
| [#93](https://github.com/hirochachacha/go-smb2/issues/93) | Cannot set a minimum dialect | todo | Only an exact dialect can be pinned, and the constants are unexported. |
| [#94](https://github.com/hirochachacha/go-smb2/issues/94) | Session ID handling under SMB 2.0.2 | todo | Reporter supplies a working patch. |
| [#97](https://github.com/hirochachacha/go-smb2/issues/97) | Only `IPC$` visible on AWS FSx | todo | Same root cause as #40. |
| [#98](https://github.com/hirochachacha/go-smb2/issues/98) | Malformed packet crashes the process | todo | Four zero bytes are enough. |
| [#99](https://github.com/hirochachacha/go-smb2/issues/99) | Is the project maintained | todo | Answerable. |
| [#101](https://github.com/hirochachacha/go-smb2/issues/101) | Panic in `PacketCodec.SessionId` | todo | Unmerged PR #102 has the fix. |

## Things we cannot test

Kept here rather than left implicit, because most of them are solvable with
setup rather than genuinely out of reach.

| Constraint | Blocks | Way around it |
|---|---|---|
| No Azure AD tenant | #73 | A free tier tenant would do it. |
| No Azure Files share | #90 | The crash itself is reproducible from a crafted response. |
| No AWS FSx | #97 | Root cause is reachable without it, from #40. |
| No Windows DFS namespace | #48, #69 | Samba can host a DFS root. |
| No old Samba (4.10) | #22 | A pinned container image. |
| No macOS SMB server in CI | #67 | Reproducible by hand on a Mac; macOS runners exist. |
