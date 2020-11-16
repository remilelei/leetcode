package Q0007_简单_整数反转

import (
	"math"
	"testing"
)

func TestQuetion7_1(t *testing.T) {
	inputX := 123
	answer := 321
	res := reverse(inputX)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuetion7_2(t *testing.T) {
	inputX := -123
	answer := -321
	res := reverse(inputX)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuetion7_3(t *testing.T) {
	inputX := 120
	answer := 21
	res := reverse(inputX)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}

func TestQuetion7_4(t *testing.T) {
	var inputX int32
	inputX = math.MaxInt32
	t.Logf("intpuX=%v", inputX)
	answer := 0
	res := reverse(int(inputX))
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
