airpaste
---

**This is a Go port of [mafintosh/airpaste][mfa].**

A 1-1 network pipe that auto discovers other peers using mdns.

#### Installation

```
go get github.com/rikonor/airpaste/airpaste
```

#### Usage

```
# On sending end
echo hello | airpaste

# On receiving end
airpaste
```

```
# Pipe files

# On sending end
airpaste < file.txt

# On receiving end
airpaste > file.txt
```

#### License
MIT

[mfa]: <https://github.com/mafintosh/airpaste>
