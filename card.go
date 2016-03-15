package rage

import "time"

func card() *Commander {
	return &Commander{
		Name: "card",
	}
}

//Card is a short description of what you would like to do. The message for the
//card should be answeres to the question  what do I want to do?
//
// Example of messages are
//	* Master golang
//	* Nail an interview
//	* Save the world
type Card struct {
	ID        int       `toml: 'id"`
	Message   string    `toml:"message"`
	Status    string    `toml:"status"`
	CreatedAt time.Time `toml:"created_at"`
	ClosedAt  time.Time `toml:"closed_at"`
}

//Cards implemets sort.Sort interface for sorting a slice of cards by ID for
//sorting a slice of cards by ID.
type Cards []*Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i].ID < c[j].ID
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
