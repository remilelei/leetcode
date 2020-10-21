package question0925

import "testing"

func TestQuestion925_1(t *testing.T) {
	name := "alex"
	typed := "aaleex"
	answer := true
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_2(t *testing.T) {
	name := "leelee"
	typed := "lleeelee"
	answer := true
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_3(t *testing.T) {
	name := "saeed"
	typed := "ssaaedd"
	answer := false
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_4(t *testing.T) {
	name := "laiden"
	typed := "laiden"
	answer := true
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_5(t *testing.T) {
	name := "saeed"
	typed := "ssaaeedd"
	answer := true
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_6(t *testing.T) {
	name := "saeed"
	typed := "ssaaeedf"
	answer := false
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}

func TestQuestion925_7(t *testing.T) {
	name := "pyplrz"
	typed := "ppyypllr"
	answer := false
	res := isLongPressedName(name, typed)
	if res != answer {
		t.Fatalf("wrong result, res=%v but answer=%v", res, answer)
	}
}
