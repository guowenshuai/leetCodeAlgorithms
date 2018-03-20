package _57JudgeRouteCircle

import "testing"

func Test_JudgeRouteCricle(t *testing.T) {
	if judgeCircle("LLRR") {
		t.Log("test success")
	} else {
		t.Error("can't pass")
	}
}
