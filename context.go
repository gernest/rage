package main

import (
	"fmt"
	"io"

	"github.com/gosuri/uitable"
)

//Context is the environmant in which the commands are executed
type Context struct {
	Commands []*Commander
	Stdout   io.Writer
}

//Help prints the hep message for the command  or optionally the subcommand  of
//command cmd if provided
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

//Printhelp prints cmd in a good interface
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
