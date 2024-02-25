package RemoveAllAdjacentDuplicatesInString

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abbaca", want: "ca", args: args{s: "abbaca"},
		}, {
			name: "azxxzy", want: "ay", args: args{s: "azxxzy"},
		}, {
			name: "azxxz", want: "a", args: args{s: "azxxz"},
		}, {
			name: "zxxzy", want: "y", args: args{s: "zxxzy"},
		}, {
			name: "zxaxz", want: "zxaxz", args: args{s: "zxaxz"},
		}, {
			name: "a", want: "a", args: args{s: "a"},
		}, {
			name: "zxxzzy", want: "zy", args: args{s: "zxxzzy"},
		}, {
			name: "babbcbccc", want: "bacbc", args: args{s: "babbcbccc"},
		}, {
			name: "bbc", want: "c", args: args{s: "bbc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
