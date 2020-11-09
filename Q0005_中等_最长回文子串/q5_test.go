package Q0005_中等_最长回文子串

import "testing"

func TestQuestion5_1(t *testing.T) {
	input := "babad"
	answer := "bab"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_2(t *testing.T) {
	input := "cbbd"
	answer := "bb"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_3(t *testing.T) {
	input := "a"
	answer := "a"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_4(t *testing.T) {
	input := "ac"
	answer := "a"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_5(t *testing.T) {
	input := "abb"
	answer := "bb"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_6(t *testing.T) {
	input := "aacabdkacaa"
	answer := "aca"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_7(t *testing.T) {
	input := "babaddtattarrattatddetartrateedredividerb"
	answer := "ddtattarrattatdd"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_8(t *testing.T) {
	input := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	answer := "ranynar"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_9(t *testing.T) {
	input := "ranynartionsomthee"
	answer := "ranynar"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_10(t *testing.T) {
	input := "zttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknks"
	answer := "sknks"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
func TestQuestion5_11(t *testing.T) {
	input := "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"
	answer := "sknks"
	res := longestPalindrome(input)
	if res != answer {
		t.Fatalf("answer=%v, res=%v", answer, res)
	}
}
