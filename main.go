package main

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/gosuri/uitable"
)

type Rage struct {
	cmds       map[string]Command
	ctx        *Context
	commanders map[string]*Commander
}

func NewRage() *Rage {
	return &Rage{
		cmds:       make(map[string]Command),
		commanders: make(map[string]*Commander),
		ctx:        &Context{},
	}
}

func (r *Rage) Register(c *Commander) {
	r.ctx.Commands = append(r.ctx.Commands, c)
}

func (r *Rage) Exec(cmd *CommaandArgs) error {
	for _, c := range r.ctx.Commands {
		if c.Name == cmd.Name {
			return c.Exec(r.ctx, cmd)
		}
	}
	fmt.Println("unrecogised command: " + cmd.Name)
	return nil
}

func (r *Rage) Run() {
	rl, err := readline.New("rage> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	/*

		REGISTER RAGE COMMANDS

	*/
	r.Register(board())
	r.Register(team())
	r.Register(card())
	r.Register(assign())

	/*

		REGISTER UTILITY COMMANDS

	*/
	r.Register(clear())
	r.Register(help())
	r.Register(exit())

	msg := ""
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}
		words, err := parseArgsLine(line)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		cmd := LoadCommandArgs(words)
		err = r.Exec(cmd)
		if err != nil {
			msg = err.Error()
			break
		}
	}
	fmt.Println(msg)
}

func (r *Rage) Help(ctx *Context, cmd *CommaandArgs) error {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	table.AddRow("command name", "description")
	for k, _ := range r.cmds {
		table.AddRow(k, k) // blank
	}
	fmt.Println(table)
	return nil
}

func main() {
	rage := NewRage()
	rage.Run()
}
