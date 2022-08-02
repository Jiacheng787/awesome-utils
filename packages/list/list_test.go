package list

import (
	"reflect"
	"testing"
)

func TestList_Filter(t *testing.T) {
	type args struct {
		fn func(int, int) bool
	}
	tests := []struct {
		name string
		l    List
		args args
		want List
	}{
		{
			name: "test_filter",
			l:    New([]int{2, 4, 6, 8, 9}),
			args: args{func(item, index int) bool {
				return item%2 == 0
			}},
			want: New([]int{2, 4, 6, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Filter(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_ForEach(t *testing.T) {
	type args struct {
		fn func(int, int)
	}
	tests := []struct {
		name string
		l    List
		args args
	}{
		{
			name: "test_forEach",
			l:    New([]int{2, 4, 6, 8}),
			args: args{func(item int, index int) {
				t.Logf("current value: %d, index: %d", item, index)
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.ForEach(tt.args.fn)
		})
	}
}

func TestList_Map(t *testing.T) {
	type args struct {
		fn func(int, int) int
	}
	tests := []struct {
		name string
		l    List
		args args
		want List
	}{
		{
			name: "test_map",
			l:    New([]int{2, 4, 6, 8}),
			args: args{func(item int, index int) int {
				return item + 1
			}},
			want: New([]int{3, 5, 7, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Map(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
