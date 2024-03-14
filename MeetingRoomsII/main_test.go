package MeetingRoomsII

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard",
			args: args{intervals: [][]int{{0, 30}, {5, 10}, {15, 20}, {35, 40}}},
			want: 2,
		}, {
			name: "standard",
			args: args{intervals: [][]int{{7, 10}, {2, 4}}},
			want: 1,
		}, {
			name: "standard",
			args: args{intervals: [][]int{{2, 7}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
