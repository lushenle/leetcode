package zeller

import "testing"

type args struct {
	year  int
	month int
	day   int
}

func Test_zeller(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				year:  2022,
				month: 1,
				day:   13,
			},
			want: 4,
		},
		{
			name: "test02",
			args: args{
				year:  1992,
				month: 7,
				day:   21,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeller(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("Zeller() = %v, want %v", got, tt.want)
			}
		})
	}
}
