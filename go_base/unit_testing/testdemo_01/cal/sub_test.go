package cal

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := getSub(3, 2)
	if res != 1 {
		t.Fatalf("getSub(3,2)执行错误,期望值=%v,实际值为=%v\n", 1, res)
	}
	t.Logf("getSub(3,2)执行正确!!!!!!")
}
