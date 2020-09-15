package MyRegularExpressionMatching

import (
	"reflect"
	"testing"
)

func TestRegularExpressionMatching(t *testing.T) {

	type test struct {
		s    string
		p    string
		want bool
	}
	tests := map[string]test{
		"case_1": {"aa", "a", false},
		"case_2": {"aa", "a*", true},
		"case_3": {"ab", ".*", true},
		"case_4": {"aab", "c*a*b", true},
		"case_5": {"mississippi", "mis*is*p*.", false},
	}

	for name, ts := range tests {
		got := isMatch(ts.s, ts.p)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, ts.want, got)
		}
	}
}
