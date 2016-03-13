package rage

import (
	"sort"
	"testing"
)

func TestCards(t *testing.T) {
	samle := []struct {
		id int
	}{
		{2},
		{8},
		{1},
		{4},
	}
	var c Cards
	for _, v := range samle {
		c = append(c, &Card{ID: v.id})
	}
	sort.Sort(c)
	expect := []int{1, 2, 4, 8}
	for k, v := range c {
		if v.ID != expect[k] {
			t.Errorf("expected %d got %d", expect[k], v.ID)
		}
	}
}
