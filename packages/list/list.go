package list

type List []int

func New(arr []int) *List {
	ins := List(arr)
	return &ins
}

func (l *List) ForEach(fn func(int, int)) {
	for index, item := range *l {
		fn(item, index)
	}
}

func (l *List) Map(fn func(int, int) int) *List {
	res := make([]int, 0)
	for index, item := range *l {
		res = append(res, fn(item, index))
	}
	return New(res)
}

func (l *List) Filter(fn func(int, int) bool) *List {
	res := make([]int, 0)
	for index, item := range *l {
		if fn(item, index) {
			res = append(res, item)
		}
	}
	return New(res)
}
