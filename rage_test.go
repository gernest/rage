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

func TestSanitizeName(t *testing.T) {
	sample := []struct {
		src    []string
		expect string
	}{
		{[]string{"hello", "world"}, "hello-world"},
		{[]string{"hello", ",", "world"}, "hello-world"},
	}

	for _, v := range sample {
		s := sanitizeName(v.src)
		if s != v.expect {
			t.Errorf("expected %s got %s", v.expect, s)
		}

	}
}
