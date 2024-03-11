package RemoveAllAdjacentDuplicatesInStringII

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard",
			args: args{
				s: "deeedbbcccbdaa",
				k: 3,
			},
			want: "aa",
		}, {
			name: "standard",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "abcd",
		}, {
			name: "standard",
			args: args{
				s: "pbbcggttciiippooaais",
				k: 2,
			},
			want: "ps",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
