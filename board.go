package rage

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

const (

	//boardInfoFile is the name of the file that contains board information,. This
	//file is located inside the board directory.
	boardInfoFile = "board.info"

	// board status. It is bad practice if we delete the whole board from the
	// disk, just in case the user later changed his ming( forget about version
	// control)
	//
	// These are markers that will be labeled on boards, so deleted boards will
	// still be in the system but only labelled deleted.
	statusOpen    = "open"
	statusClosed  = "closed"
	statusDeleted = "deleted"
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
		name := sanitizeName(cmd.Args)
		if name != "" {
			bPath := filepath.Join(ctx.Workspace, name)
			_, err := os.Stat(bPath)
			if os.IsExist(err) {
				ctx.Printf("The board %s already exist\n", name)
				return nil
			}
			err = os.MkdirAll(bPath, os.ModePerm)
			if err != nil {
				ctx.Println(err.Error())
				return nil
			}
			b := &Board{}
			b.Name = strings.Join(cmd.Args, " ")
			b.CreatedAt = time.Now()
			b.Status = statusOpen
			f, err := os.Create(filepath.Join(bPath, boardInfoFile))
			if err != nil {
				ctx.Println(err.Error())
				return nil
			}
			err = toml.NewEncoder(f).Encode(b)
			_ = f.Close()
			if err != nil {
				ctx.Println(err.Error())
				return nil
			}
			ctx.Println(" successful created " + name)
			return nil
		}
	}
	ctx.Println(" you need to specify the name of the board")
	return nil
}

func updateeBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

func deleteBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

//Board is the module(or container) in which cards are stored. This is a way to
//organize your ideas. Just think of the board as a blackboard in which you will
//write down what you think you should do.
//
// Inside rage, boards are directories, Name should be short and meaningful
type Board struct {
	Name      string    `toml:"name"`
	Status    string    `toml:"status"`
	CreatedAt time.Time `toml:"created_at"`

	Cards Cards `toml:"-"`
}
