package linked_list

import "github.com/duckhue01/ps/data/list"

func makeList(es []int) *list.S {
	l := new(list.S)
	for _, v := range es {
		l.Append(v)
	}
	return l
}
