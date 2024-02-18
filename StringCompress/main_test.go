package StringCompress

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 6, args: args{chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}},
		{name: "test", want: 4, args: args{chars: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}}},
		{name: "test", want: 1, args: args{chars: []byte{'a'}}},
		{name: "test", want: 2, args: args{chars: []byte{'a', 'b'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
