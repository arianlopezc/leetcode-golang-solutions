package RomanToInteger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard", want: 3, args: args{s: "III"},
		}, {
			name: "standard", want: 1, args: args{s: "I"},
		}, {
			name: "standard", want: 58, args: args{s: "LVIII"},
		}, {
			name: "standard", want: 1994, args: args{s: "MCMXCIV"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
