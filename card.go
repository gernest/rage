package rage

import "time"

func card() *Commander {
	return &Commander{
		Name: "card",
	}
}

type Card struct {
	ID        int       `toml: 'id"`
	Message   string    `toml:"message"`
	Status    string    `toml:"status"`
	CreatedAt time.Time `toml:"created_at"`
	ClosedAt  time.Time `toml:"closed_at"`
}

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
