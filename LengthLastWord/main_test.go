package LengthLastWord

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 5, args: args{s: "Hello World"}},
		{name: "test", want: 4, args: args{s: "   fly me   to   the moon  "}},
		{name: "test", want: 6, args: args{s: "luffy is still joyboy"}},
		{name: "test", want: 0, args: args{s: "   "}},
		{name: "test", want: 1, args: args{s: "s   "}},
		{name: "test", want: 1, args: args{s: "    s"}},
		{name: "test", want: 4, args: args{s: "sasd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
