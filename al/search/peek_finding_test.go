package search

import "testing"

// Peek Finding: Find an element in array that satisfy a[x] >= x[x-1] && a[x] <= a[x+1]

func PeekFindingWithBinarySearch(a []int) (int, bool) {
	if len(a) < 1 {
		return 0, false
	}
	if len(a) == 1 {
		return a[0], true
	}

	if len(a) == 2 {
		if a[0] >= a[1] {
			return a[0], true
		} else {
			return a[1], true
		}
	}

	i := len(a) / 2
	for i > 0 && i < len(a)-1 {
		if a[i] >= a[i-1] && a[i] >= a[i+1] {
			return a[i], true
		}
		if a[i] < a[i-1] {
			i = i / 2
			continue
		}
		if a[i] < a[i+1] {
			i = (len(a) + i) / 2
			continue
		}

	}

	return a[i], true
}

func TestPeekFindingWithBinarySearch(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "input five elements expect 5 and true",
			args: args{
				a: []int{1, 2, 3, 4, 5},
			},
			want:  5,
			want1: true,
		},
		{
			name: "input one elements expect 1 and true",
			args: args{
				a: []int{1},
			},
			want:  1,
			want1: true,
		},

		{
			name: "input two elements expect 2 and true",
			args: args{
				a: []int{1, 2},
			},
			want:  2,
			want1: true,
		},
		{
			name: "input two elements expect 2 and true",
			args: args{
				a: []int{2, 1},
			},
			want:  2,
			want1: true,
		},
		{
			name: "input 4 elements expect 6 and true",
			args: args{
				a: []int{5, 4, 6, 3},
			},
			want:  6,
			want1: true,
		},
		{
			name: "input two elements expect 2 and true",
			args: args{
				a: []int{5, 4, 3, 6, 7},
			},
			want:  5,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PeekFindingWithBinarySearch(tt.args.a)
			if got != tt.want {
				t.Errorf("PeekFindingWithBinarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PeekFindingWithBinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
