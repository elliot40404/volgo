package cli

import (
	"errors"
	"flag"
	"fmt"
)

const (
	Version = "0.0.1"
	Author  = "elliot40404<avishek40404@gmail.com>"
	Name    = "volgo"
	Desc    = "volgo is a cross platform cli app that helps to control volume"
	Example = "volgo <cmd> <level>"
)

var (
	ErrHelp    = errors.New("help")
	ErrVersion = errors.New("version")
)

func help() {
	fmt.Printf(`%s %s

%s

Commands:
	status  - Get current volume and mute status
	set     - Set volume level
	mute    - Mute audio
	unmute  - Unmute audio
	inc     - Increase volume level
	dec     - Decrease volume level

Options:
	-h, --help            Show this help message
	-v, --version         Show version information

Usage:
	volgo <cmd> <level>
	volgo <mute|unmute|get>
	volgo <set|inc|dec> <level>

Examples:
	volgo get
	volgo set 50
	volgo inc 10
	volgo mute
`, Name, Version, Desc)
}

type ParsedArgs struct {
	Cmd string
	Lvl string
}

func ParseArgs() (ParsedArgs, error) {
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("v", false, "Show version information")
	versionLongFlag := flag.Bool("version", false, "Show version information")

	flag.Usage = help

	flag.Parse()

	p := ParsedArgs{}

	if *helpFlag {
		help()
		return p, ErrHelp
	}

	if *versionFlag || *versionLongFlag {
		fmt.Printf("%s %s\n", Name, Version)
		return p, ErrVersion
	}

	switch flag.NArg() {
	case 0:
		return p, nil
	case 1:
		p.Cmd = flag.Arg(0)
	case 2:
		p.Cmd = flag.Arg(0)
		p.Lvl = flag.Arg(1)
	default:
		return ParsedArgs{}, errors.New("too many arguments")
	}

	return p, nil
}
