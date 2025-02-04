package controller

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/itchyny/volume-go"

	"github.com/elliot40404/volgo/internal/cli"
)

type command string

const (
	Status command = "status"
	Set    command = "set"
	Mute   command = "mute"
	Unmute command = "unmute"
	Inc    command = "inc"
	Dec    command = "dec"
)

type Controller struct {
	Cmd command
	Lvl int
}

var (
	ErrInvalidCommand = errors.New("invalid command")
	ErrInvalidLevel   = errors.New("invalid level")
)

func isValidCommand(cmd command) bool {
	switch cmd {
	case Status, Set, Mute, Unmute, Inc, Dec, "":
		return true
	default:
		return false
	}
}

func isValidLevel(lvl string) (bool, int) {
	if lvl == "" {
		return false, 0
	}
	level, err := strconv.Atoi(lvl)
	if err != nil {
		return false, 0
	}
	if level < 0 || level > 100 {
		return false, 0
	}
	return true, level
}

func NewController(args cli.ParsedArgs) (*Controller, error) {
	if !isValidCommand(command(args.Cmd)) {
		return nil, ErrInvalidCommand
	}
	if slices.Contains([]command{Inc, Dec, Set}, command(args.Cmd)) {
		valid, lvl := isValidLevel(args.Lvl)
		if !valid {
			return nil, ErrInvalidLevel
		}
		return &Controller{
			Cmd: command(args.Cmd),
			Lvl: lvl,
		}, nil
	}
	return &Controller{
		Cmd: command(args.Cmd),
		Lvl: 0,
	}, nil
}

func (Controller) GetVolume() int {
	vol, err := volume.GetVolume()
	if err != nil {
		return -1
	}
	return vol
}

func (c *Controller) SetVolume() {
	err := volume.SetVolume(c.Lvl)
	if err != nil {
		return
	}
}

func (Controller) Mute() {
	err := volume.Mute()
	if err != nil {
		return
	}
}

func (Controller) Unmute() {
	err := volume.Unmute()
	if err != nil {
		return
	}
}

func (c *Controller) IncreaseVolume(vol int) {
	err := volume.IncreaseVolume(c.Lvl + vol)
	if err != nil {
		return
	}
}

func (c *Controller) DecreaseVolume(vol int) {
	err := volume.IncreaseVolume(-1 * (c.Lvl + vol))
	if err != nil {
		return
	}
}

func (Controller) GetMuted() bool {
	muted, err := volume.GetMuted()
	if err != nil {
		return false
	}
	return muted
}

func (c *Controller) Exec() (string, error) {
	switch c.Cmd {
	case Status:
		vol := c.GetVolume()
		return fmt.Sprintf("Current volume: %d%%\nMuted: %t", vol, c.GetMuted()), nil
	case Set:
		c.SetVolume()
		return fmt.Sprintf("Volume set to %d%%", c.Lvl), nil
	case Mute:
		c.Mute()
		return "Audio muted", nil
	case Unmute:
		c.Unmute()
		return "Audio unmuted", nil
	case Inc:
		c.IncreaseVolume(0)
		return fmt.Sprintf("Volume increased by %d%%", c.Lvl), nil
	case Dec:
		c.DecreaseVolume(0)
		return fmt.Sprintf("Volume decreased by %d%%", c.Lvl), nil
	}
	return "", nil
}
