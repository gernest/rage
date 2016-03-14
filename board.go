package rage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var boadrDescription = `
board is a namespace in which you can manage cards
calling this command alone will display this message.

Use the subcommands to do maipulations of the boards
run the following command on the rage prompt
help board
		`

func board() *Commander {
	return &Commander{
		Name:        "board",
		Description: boadrDescription,
		Subcommands: []*Commander{
			{
				Name:        "create",
				Description: "creates a new board",
				Exec:        createBoard,
			},
			{
				Name:        "update",
				Description: "update  board",
				Exec:        createBoard,
			},
			{
				Name:        "delete",
				Description: "delete  board",
				Exec:        deleteBoard,
			},
		},
		Exec: func(ctx *Context, cmd *CommaandArgs) error {
			ctx.Help("board")
			return nil
		},
	}

}

//createBoard creates a new board with the name specified. The creation date is
//set to time.Now()
//
// It is illegal to create a board that already exists. A board is simply a
// direcory identified by name(board name) in which cards are stored
//
// Directory for boards are relative to workspace, like
// $RAGE_WORKSPACE/{board_name}. Boards can never be nested as it is kinda
// owkward to do that.
func createBoard(ctx *Context, cmd *CommaandArgs) error {
	if len(cmd.Args) > 0 {
		name := sainizeName(cmd.Args)
		if name != "" {
			bPath := filepath.Join(ctx.Workspace, name)
			_, err := os.Stat(bPath)
			if os.IsExist(err) {
				fmt.Fprintf(ctx, "The board %s already exist\n", name)
				return nil
			}
			err = os.MkdirAll(bPath, os.ModePerm)
			if err != nil {
				fmt.Fprintln(ctx, err.Error())
				return nil
			}
			return nil
		}
	}
	fmt.Fprintln(ctx, " you need to specify the name of the board")
	return nil
}

func updateeBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

func deleteBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

type Board struct {
	Name      string    `toml:"name"`
	CreatedAt time.Time `toml:"created_at"`

	Cards Cards `toml:"-"`
}
