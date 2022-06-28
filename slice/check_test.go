package slice

import (
	"fmt"
	"testing"
)

func Test_CheckSlcHasStr(t *testing.T) {
	a := []string{"a","b","c"}
	c := "c"
	fmt.Printf("result:%v\n",CheckSlcHasStr(c,a))
}
