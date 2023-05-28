package list

import (
	"testing"
)

func TestSList(t *testing.T) {
	l := new(SList)

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Show()
}
