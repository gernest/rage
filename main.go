package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

// Rage is the rage application
type Rage struct {
	cmds       map[string]Command
	ctx        *Context
	commanders map[string]*Commander
}

//NewRage retturns a new instance of* Rage
func NewRage() *Rage {
	return &Rage{
		cmds:       make(map[string]Command),
		commanders: make(map[string]*Commander),
		ctx:        &Context{},
	}
}

//Register registers cmd into the rage application. registering the same command
// result into reasigning the command name to the new command
func (r *Rage) Register(c *Commander) {
	r.ctx.Commands = append(r.ctx.Commands, c)
}

//Exec executes the command that matches cmd
func (r *Rage) Exec(cmd *CommaandArgs) error {
	for _, c := range r.ctx.Commands {
		if c.Name == cmd.Name {
			return c.Exec(r.ctx, cmd)
		}
	}
	fmt.Println("unrecogised command: " + cmd.Name)
	return nil
}

//Run runs the rage application, it boots up an interactive shell.
func (r *Rage) Run() {
	rl, err := readline.New("rage> ")
	if err != nil {
		panic(err)
	}
	defer func() { _ = rl.Close() }()

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

func main() {
	rage := NewRage()
	rage.Run()
}
