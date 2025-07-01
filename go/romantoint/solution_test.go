package romantoint

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i),
		func(t *testing.T) {
			t.Parallel()
			res := RomanToInt(ts.input)

			if !reflect.DeepEqual(ts.output, res) {
				t.Errorf("Expected: %v; got: %v", ts.output, res)
			}
		})
	}
}
