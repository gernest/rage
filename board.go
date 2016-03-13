package rage

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
func createBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

func updateeBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}

func deleteBoard(ctx *Context, cmd *CommaandArgs) error {
	return nil
}
