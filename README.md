# panicfixer
panicfixer is a tool to check your system for any problem with the installed packages which may cause issues in the next patch,
panicfixer also fixed this issue, but not by default to avoid further issues to your system.
Simply fix the issues if you don't see any bigger issue.

## Install

```bash
go get -u github.com/hhemied/panicfixer
```
Or

```
git clone https://github.com/hhemied/panicfixer.git
cd panicfixer 
sudo ./rhelpatchfoix
```
## Examples
panicfixer has 2 subcommands

```
⌦  sudo ./panicfixer
Checks installed package for future update for any problems may happen and fix them

This tool is working only on RedHat based OS like [RHEL, CentOS, Fedora]

Usage:
  panicfixer [command]

Available Commands:
  check       checking errors
  fix         Fix Errors
  help        Help about any command

Flags:
  -h, --help   help for panicfixer

Use "panicfixer [command] --help" for more information about a command.
```

### Checking Errors
```
Checking your system...
Found Issues In:
+---+-----------------------------------+
| # | PACKAGE NAME                      |
+---+-----------------------------------+
| 1 | erlang-sd_notify-0.1-1.el7.x86_64 |
| 2 | erlang-erts-R16B-03.18.el7.x86_64 |
| 3 | erlang-erts-21.3-1.el7.x86_64     |
+---+-----------------------------------+
```

### Fixing Errors
```
⌦  sudo ./panicfixer fix
⌦  sudo ./panicfixer fix
Deleting erlang-ssl-R16B-03.18.el7.x86_64
Deleting erlang-erts-R16B-03.18.el7.x86_64
Deleting erlang-sd_notify-0.1-1.el7.x86_64
Deleting erlang-sasl-R16B-03.18.el7.x86_64
Deleting erlang-crypto-R16B-03.18.el7.x86_64
Deleting erlang-snmp-R16B-03.18.el7.x86_64
Deleting erlang-asn1-R16B-03.18.el7.x86_64
Deleting erlang-tools-R16B-03.18.el7.x86_64
Deleting erlang-kernel-R16B-03.18.el7.x86_64
Deleting erlang-hipe-R16B-03.18.el7.x86_64
Deleting erlang-otp_mibs-R16B-03.18.el7.x86_64
Deleting erlang-inets-R16B-03.18.el7.x86_64
Deleting erlang-stdlib-R16B-03.18.el7.x86_64
Deleting erlang-mnesia-R16B-03.18.el7.x86_64
Deleting erlang-xmerl-R16B-03.18.el7.x86_64
Deleting erlang-os_mon-R16B-03.18.el7.x86_64
Deleting erlang-syntax_tools-R16B-03.18.el7.x86_64
Deleting erlang-runtime_tools-R16B-03.18.el7.x86_64
Deleting erlang-public_key-R16B-03.18.el7.x86_64
Deleting erlang-compiler-R16B-03.18.el7.x86_64
Trying to return system to ideal state ..
```

```
⌦  sudo ./panicfixer fix
At the moment : Your system has no issues with installed packages..
```
