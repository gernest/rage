package rage

import (
	"reflect"
	"testing"
)

func TestParseArgsLine(t *testing.T) {
	sample := []struct {
		line   string
		expect []string
	}{
		{"new card", []string{"new", "card"}},
	}
	for _, v := range sample {
		cmds, err := parseArgsLine(v.line)
		if err != nil {
			t.Error(err)
			continue
		}
		if !reflect.DeepEqual(cmds, v.expect) {
			t.Errorf("expected %v got %v", v.expect, cmds)
		}
	}
}
