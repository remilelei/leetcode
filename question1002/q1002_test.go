package question1002

import (
	"reflect"
	"sort"
	"testing"
)

func TestQ1002_1(t *testing.T) {
	inputData := []string{"bella", "label", "roller"}
	answerData := []string{"e", "l", "l"}
	res := commonChars(inputData)
	sort.Strings(res)
	sort.Strings(answerData)
	if !reflect.DeepEqual(res, answerData) {
		t.Fatalf("res error! answer=%v, res=%v", answerData, res)
	}
}

func TestQ1002_2(t *testing.T) {
	inputData := []string{"cool", "lock", "cook"}
	answerData := []string{"c", "o"}
	res := commonChars(inputData)
	sort.Strings(res)
	sort.Strings(answerData)
	if !reflect.DeepEqual(res, answerData) {
		t.Fatalf("res error! answer=%v, res=%v", answerData, res)
	}
}
