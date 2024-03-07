package DietPlanPerformance

import "testing"

func Test_dietPlanPerformance(t *testing.T) {
	type args struct {
		calories []int
		k        int
		lower    int
		upper    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard", want: 0, args: args{k: 1, upper: 3, lower: 3, calories: []int{1, 2, 3, 4, 5}},
		}, {
			name: "standard", want: 1, args: args{k: 2, upper: 1, lower: 0, calories: []int{3, 2}},
		}, {
			name: "standard", want: 0, args: args{k: 2, upper: 5, lower: 1, calories: []int{6, 5, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dietPlanPerformance(tt.args.calories, tt.args.k, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("dietPlanPerformance() = %v, want %v", got, tt.want)
			}
		})
	}
}
