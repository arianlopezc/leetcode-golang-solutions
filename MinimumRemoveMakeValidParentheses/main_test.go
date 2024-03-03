package MinimumRemoveMakeValidParentheses

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "lee(t(c)o)de)", want: "lee(t(c)o)de", args: args{s: "lee(t(c)o)de)"},
		}, {
			name: "a)b(c)d", want: "ab(c)d", args: args{s: "a)b(c)d"},
		}, {
			name: "))((", want: "", args: args{s: "))(("},
		}, {
			name: ")((a)(", want: "(a)", args: args{s: ")((a)("},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
