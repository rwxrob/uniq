package uniq

import (
	"strconv"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

// Cmd is a composable Bonzai command branch
var Cmd = &Z.Cmd{
	Name:      `uniq`,
	Summary:   `universal unique identifiers`,
	Version:   `v0.2.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{help.Cmd,
		IsosecCmd, IsosecTCmd, IsonanCmd, HexCmd, SecondCmd, UUIDCmd, Base32Cmd,
	},
}

// IsosecCmd is an composable Bonzai leaf command
var IsosecCmd = &Z.Cmd{
	Name:     `isosec`,
	Summary:  `sortable unique second in UTC`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command returns the UTC current time in ISO8601
		(RFC3339) without any punctuation or the T. This is frequently
		a very good unique suffix that has the added advantage of being
		chronologically sortable and more readable than the epoch. (Also see
		Second) `,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(Isosec())
		return nil
	},
}

var IsosecTCmd = &Z.Cmd{
	Name:     `isosect`,
	Summary:  `sortable unique second in UTC (with T)`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command returns the UTC current time in ISO8601
		(RFC3339) without any punctuation but includes the T. This is frequently
		a very good unique suffix that has the added advantage of being
		chronologically sortable and more readable than the epoch. (Also see
		Second) `,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(IsosecT())
		return nil
	},
}

// IsonanCmd is an composable Bonzai leaf command
var IsonanCmd = &Z.Cmd{
	Name:     `isonan`,
	Summary:  `sortable unique nanosecond in UTC`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command returns the UTC current time in ISO8601
		(RFC3339) nanosecond without any punctuation or the T. This is frequently
		a very good unique suffix that has the added advantage of being
		chronologically sortable and more readable than the epoch. This
		provides considerable more uniqueness compared to {{cmd "isosec"}}`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(Isonan())
		return nil
	},
}

// Isodate is an composable Bonzai leaf command
var IsodateCmd = &Z.Cmd{
	Name:     `isodate`,
	Summary:  `prints human-friendly time to the second UTC`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{cmd .Name}} command returns the UTC current time in ISO8601
		(RFC3339) in a human friendly format with the dashes and colons and the
		Z at the end. This identifier does have a space but is easier to ingest
		when dealing with databases that also use timestamps for uniqueness the
		Z is always present preventing mistakes with different time zones.`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(Isodate())
		return nil
	},
}

// HexCmd is an composable Bonzai leaf command
var HexCmd = &Z.Cmd{
	Name:    `hex`,
	Summary: `pseudo-random bytes as hexadecimal`,
	Usage:   `<length>`,
	MinArgs: 1,

	Description: `
		 The {{cmd .Name}} command returns a random hexadecimal string that
		 is n*2 characters in length.  Calling {{cmd .Name}} is superior to
		 {{cmd "uuid"}} because it fits into the same 36 character space but
		 sometimes content limitations and validators require only
		 hexadecimal characters.`,

	Call: func(_ *Z.Cmd, args ...string) error {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		term.Print(Hex(n))
		return nil
	},
}

// SecondCmd is an composable Bonzai leaf command
var SecondCmd = &Z.Cmd{
	Name:    `second`,
	Summary: `UNIX second since epoch`,

	Description: `
		 The {{cmd .Name}} command returns the current (time.Unix) second
		 since Jan 1, 1970 UTC. This is frequently a very good unique suffix
		 that has the added advantage of being chronologically sortable.`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(Second())
		return nil
	},
}

// UUIDCmd is an composable Bonzai leaf command
var UUIDCmd = &Z.Cmd{
	Name:    `uuid`,
	Summary: `standard UUID v4 (RFC 4122)`,

	Description: `
		The {{cmd .Name}} command returns a standard UUID v4 according to
		RFC 4122.  UUIDs have become deeply entrenched universally but are
		inferior to Base32 as the need for a greater range of randomness
		emerges (IPv6, etc.) Returns empty string if unable to read random
		data for any reason.`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(UUID())
		return nil
	},
}

// Base32Cmd is an composable Bonzai leaf command
var Base32Cmd = &Z.Cmd{
	Name:    `base32`,
	Summary: `base32 20 byte unique identifier`,

	Description: `
		 The {{cmd .Name}} returns a base32 encoded 20 byte string. This has
		 a greater range than {{cmd "uuid"}} and is safe for use with
		 filesystems.  Base32 is rendered in uppercase for clarity and
		 because it is case insensitive.  Base32 depends on 40 bit chunks.
		 20 bytes exceeds UUID randomness and is the closest. (15 would be
		 insufficient to cover the same range.) Base32 is therefore superior
		 to UUID both in range of randomness and practicality. Returns an
		 empty string if unable to read random data.`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		term.Print(Base32())
		return nil
	},
}
