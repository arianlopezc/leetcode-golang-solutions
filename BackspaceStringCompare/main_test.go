package BackspaceStringCompare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abc", want: true, args: args{
				s: "ab#c",
				t: "ad#c",
			},
		}, {
			name: "none", want: true, args: args{
				s: "ab##",
				t: "c#d#",
			},
		}, {
			name: "none", want: true, args: args{
				s: "ab###",
				t: "c#d##",
			},
		}, {
			name: "acb", want: false, args: args{
				s: "a#c",
				t: "t",
			},
		}, {
			name: "c", want: true, args: args{
				s: "a##c",
				t: "#a#c",
			},
		}, {
			name: "xywrrmp", want: true, args: args{
				s: "xywrrmp",
				t: "xywrrmu#p",
			},
		}, {
			name: "bxj##tw", want: true, args: args{
				s: "bxj##tw",
				t: "bxo#j##tw",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
