package m

import (
	"fmt"
	"testing"
)

var p = Policy{
	SA{"13800010002", "u1", "client"},
	AO{"D100010001", "00:11:22:33:44:55"},
	SP{[]string{"read", "write"}},
	AE{1575468182, 1576468182, "*.*.*.*"},
}

func Test_ToBytes(t *testing.T) {
	fmt.Printf("%s\n", p.ToBytes())
}

func Test_GetID(t *testing.T) {
	println(p.GetID())
}
