package Router_test

import (
	"github.com/next-core-x/Router"
	"testing"
)

func TestTest(t *testing.T) {

	sum := Router.Test(1, 2)
	if sum != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", sum)
	}
	// 假设 Router 包中有一个 Add 函数，我们在这里测试它

}
