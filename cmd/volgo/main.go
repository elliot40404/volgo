package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/elliot40404/volgo/internal/cli"
	"github.com/elliot40404/volgo/internal/controller"
	"github.com/elliot40404/volgo/internal/renderer"
)

func main() {
	args, err := cli.ParseArgs()
	if err != nil {
		if errors.Is(err, cli.ErrHelp) || errors.Is(err, cli.ErrVersion) {
			return
		}
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
	c, err := controller.NewController(args)
	if err != nil {
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
	if args.Cmd != "" {
		output, err := c.Exec()
		if err != nil {
			slog.Error("something went wrong", "error", err)
			os.Exit(1)
		}
		fmt.Println(output)
		return
	}
	r := renderer.NewRenderer(c)
	err = r.Render()
	if err != nil {
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
}
