package Q0406_中等_根据身高重建队列

import "testing"

func compareArray(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v[0] != b[i][0] || v[1] != b[i][1] {
			return false
		}
	}
	return true
}

func TestQuestion406_1(t *testing.T) {
	inputPeople := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	answer := [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
	res := reconstructQueue(inputPeople)
	if !compareArray(answer, res) {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}

func TestQuestion406_2(t *testing.T) {
	inputPeople := [][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}}
	answer := [][]int{{3, 0}, {6, 0}, {7, 0}, {5, 2}, {3, 4}, {5, 3}, {6, 2}, {2, 7}, {9, 0}, {1, 9}}
	// res = [[3 0] [6 0] [7 0] [5 2] [5 3] [3 4] [9 0] [2 7] [6 2] [1 9]]
	res := reconstructQueue(inputPeople)
	if !compareArray(answer, res) {
		t.Fatalf("answer=%v but res=%v", answer, res)
	}
}
