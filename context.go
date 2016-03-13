package main

import (
	"fmt"
	"io"

	"github.com/gosuri/uitable"
)

type Context struct {
	Commands []*Commander
	Stdout   io.Writer
}

func (ctx *Context) Help(cmd string, sub ...string) {
	args := struct {
		command, subcommand string
	}{
		cmd, "",
	}
	if len(sub) > 1 {
		args.subcommand = sub[0]
	}
	for _, c := range ctx.Commands {
		if c.Name == args.command {
			if args.subcommand != "" {
				for _, sub := range c.Subcommands {
					if sub.Name == args.subcommand {
						ctx.Printhelp(sub)
						return
					}
					println("unrecoginsed subcommand " + args.subcommand + " for " + args.command)
					return
				}
			}
			ctx.Printhelp(c)
			return
		}
	}
	println("unrecognised command " + args.command)
}
func (ctx *Context) Printhelp(cmd *Commander) {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.AddRow(cmd.Name, cmd.Description)
	if len(cmd.Subcommands) > 0 {
		table.AddRow("----")
	}
	for _, sub := range cmd.Subcommands {
		table.AddRow("  "+sub.Name, sub.Description)
	}
	fmt.Println(table)
}
