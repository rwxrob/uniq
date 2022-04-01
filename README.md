# 🌳 Go Unique Identifiers

[![GoDoc](https://godoc.org/github.com/rwxrob/uniq?status.svg)](https://godoc.org/github.com/rwxrob/uniq)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Package `uniq` is a utility package and Bonzai command branch that
provides common random unique identifiers in UUID, Base32, and n*2
random hexadecimal characters.

    6c671957-2f39-4ce5-9f0e-e8d5ec53bfde (16 bytes, 36 chars, hex-)
    H6M0STKP0MTSU0493GERQDCSJ5BMF3VO     (20 bytes, 32 chars, base32)
    5b ...                               (n bytes, n*2 chars, hex)

When a simple random identifier is all that is needed `Base32()` provides a better alternative to `UUID()`. It takes less space (32 characters), is safe for use with all file systems including case insensitive ones, and provides additional randomness increased from 2^128 (UUID) to 2^160 (Base32).

This package includes the following convenience commands as well for use when integrating with shell scripts:

* `uuid`
* `uid32`
* `isosec`
* `epoch [SECONDS]`
* `randhex [COUNT]`

## Install

This command can be installed as a standalone program or composed into 
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/uniq/uniq@latest
```

Composed

```go
package z

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/uniq"
)

var Cmd = &bonzai.Cmd{
	Name:     `z`,
	Commands: []*bonzai.Cmd{help.Cmd, uniq.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C uniq uniq
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
