package Q0072_困难_编辑距离

import "testing"

func TestQuestion72_1(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	answer := 3
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion72_2(t *testing.T) {
	word1 := "intention"
	word2 := "execution"
	answer := 5
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion72_3(t *testing.T) {
	word1 := "apple"
	word2 := "pple"
	answer := 1
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion72_4(t *testing.T) {
	word1 := ""
	word2 := ""
	answer := 0
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion72_5(t *testing.T) {
	word1 := ""
	word2 := "a"
	answer := 1
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion72_6(t *testing.T) {
	word1 := "a"
	word2 := ""
	answer := 1
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion72_7(t *testing.T) {
	word1 := "zoologicoarchaeologist"
	word2 := "zoogeologist"
	answer := 10
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion72_8(t *testing.T) {
	word1 := "oolog"
	word2 := "oog"
	answer := 2
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion72_9(t *testing.T) {
	word1 := "zoologicoarchaeologist"
	word2 := "zoopsychologist"
	answer := 10
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
func TestQuestion72_10(t *testing.T) {
	word1 := "logicoarch"
	word2 := "psych"
	answer := 8
	res := minDistance(word1, word2)
	if res != answer {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
