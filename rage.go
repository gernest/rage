//Package rage is just like trello but on the terminal.
package rage

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/chzyer/readline"
	"github.com/gosuri/uitable"
)

// CommaandArgs is the parsed commands from the rage shell.
type CommaandArgs struct {

	// Name is the first command parsed on the rage she;;
	Name string

	//Sub command is the second argument passed on the rage shell
	Subcommand string

	// Args is  everything else after the command and subcommand. This is a slice
	// of words not complete strings, so you can join them to actually get the
	// complete text
	Args []string
}

//LoadCommandArgs returns *CommandArgs from the src, src is a slice of words
//which are passed to the rage console
//
// The assumption is the first element [0] is the command, and the second if any
// [1] is the subcommand, the reminder of the elements are kept as is and stored
// in the returnder *CommandArgs.Args field.
func LoadCommandArgs(src []string) *CommaandArgs {
	cmd := &CommaandArgs{}
	switch len(src) {
	case 0:
		return nil
	case 1:
		cmd.Name = src[0]
	case 2:
		cmd.Name, cmd.Subcommand = src[0], src[1]
	default:
		cmd.Name, cmd.Subcommand = src[0], src[1]
		cmd.Args = append(cmd.Args, src[2:]...)
	}
	return cmd
}

//Commander is the instace of the commandline processor for coomands passed on
//the rage console
type Commander struct {
	Name        string
	Description string
	Short       string // short description of the command
	Subcommands []*Commander
	Exec        Command
}

//Command is an interface for a function that exeutes a command.
//
// It is good practice to always return nil even when there was an error in
// exeution , since if returned error is not nil then the rage console retimates
//
// Return error when you want the shell to be exited, for instance you can
// return an error when implementing the exit command( see function exit) as an
// example
type Command func(*Context, *CommaandArgs) error

func parseArgsLine(line string) ([]string, error) {
	s := bufio.NewScanner(strings.NewReader(line))
	s.Split(bufio.ScanWords)
	var cmds []string
	for s.Scan() {
		cmds = append(cmds, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return cmds, nil
}

/*

	SOME UTILITY COMMANDS FOR RAGE

*/
func help() *Commander {
	return &Commander{
		Name:        "help",
		Description: "shows the help message for a command",
		Exec: func(ctx *Context, cmd *CommaandArgs) error {
			if cmd.Subcommand != "" {
				ctx.Help(cmd.Subcommand)
				return nil
			}
			for _, c := range ctx.Commands {
				printShort(ctx, c)
			}
			return nil
		},
	}
}

// print hep message but with only short description
func printShort(ctx *Context, cmd *Commander) {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.AddRow(cmd.Name, cmd.Short)
	if len(cmd.Subcommands) > 0 {
		table.AddRow("---")
	}
	for _, sub := range cmd.Subcommands {
		table.AddRow("  "+sub.Name, sub.Short)
	}
	fmt.Println(table)
}

func clear() *Commander {
	return &Commander{
		Name:        "clear",
		Description: "clears the screen",
		Exec:        clearScreen,
	}
}
func clearScreen(ctx *Context, args *CommaandArgs) error {
	switch runtime.GOOS {
	case "darwin", "linux":
		fmt.Print("\033[H\033[2J")
	}
	return nil
}
func exit() *Commander {
	return &Commander{
		Name:        "exit",
		Description: "exits the rage shell",
		Exec: func(ctx *Context, cmd *CommaandArgs) error {
			return errors.New("exiting rage shell")
		},
	}
}

// Rage is the rage application
type Rage struct {
	cmds       map[string]Command
	ctx        *Context
	commanders map[string]*Commander
}

//NewRage returns a new instance of *Rage
// TThe working space( directory where rage stores files) is properly set. First
// his checks if the environment variable RAGE_WORKSPACE is set, if so then it
// will be used.
//
//In case the RAGE_WORKSPACE environment is not set, $HOME/.rage is used
//instead, note that this direcory will be created if it doesnot exist yet.
func NewRage() *Rage {
	r := &Rage{
		cmds:       make(map[string]Command),
		commanders: make(map[string]*Commander),
		ctx:        &Context{Stdout: os.Stdout},
	}
	workspace := os.Getenv("RAGE_WORKSPACE")
	if workspace == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		workspace = filepath.Join(usr.HomeDir, ".rage")
		err = os.MkdirAll(workspace, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	r.ctx.Workspace = workspace
	return r
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
			if cmd.Subcommand != "" {
				for _, sub := range c.Subcommands {
					if sub.Name == cmd.Subcommand {
						return sub.Exec(r.ctx, cmd)
					}

				}
				r.ctx.Println("command not found")
				return nil
			}
			return c.Exec(r.ctx, cmd)
		}
	}
	r.ctx.Println("unrecogised command: " + cmd.Name)
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

//sanitizeName join src to a single string suitable to be used as a filename.
//this removes some of characters that are deemed no worthy to appear as file
//names.
//
// The information removed is irrecoverable, so in some cases it wont be
// possible to recover the original string slice from the returned string
func sanitizeName(src []string) string {
	if len(src) > 0 {
		var rst []string
		for _, v := range src {
			switch v {
			case ",":
				continue
			default:
				rst = append(rst, v)
			}
		}
		return strings.Join(rst, "-")
	}
	return ""
}
