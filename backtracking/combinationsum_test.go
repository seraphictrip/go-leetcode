package backtracking_test

import (
	"dsa/backtracking"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var combinationSumTests = []struct {
	candidates []int
	target     int
	expected   [][]int
}{
	// {[]int{2}, 1, nil},
	// 					[1,2]
	//				[1]				[2]
	//			[1,1] [1,2]
	// [1,1,1]	[1,1,2]
	// {[]int{1, 2}, 3, [][]int{{1, 1, 1}, {1, 2}}},

	//							[]
	//			2					3			6			7
	//	[2,2]
	// {[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
}

func TestCombinationSums(t *testing.T) {
	for i, e := range combinationSumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := backtracking.CombinationSum(e.candidates, e.target)
			AssertCombinations(t, e.candidates, e.target, actual, e.expected)
		})
	}
}

func AssertCombinations(t *testing.T, candidates []int, target int, actual [][]int, expected [][]int) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("CombinationSum(%v, %v) = %v, want %v", candidates, target, actual, expected)
	}
}

func intToRoman(num int) string {
	result := strings.Builder{}
	if num >= 1000 {
		// calc thousands place
		ms := num / 1000
		for i := 0; i < ms; i++ {
			result.WriteString("M")
		}
		num -= ms * 1000
	}

	if num >= 100 {
		// calc hundreds
		hundreds := num / 100
		if hundreds == 9 {
			result.WriteString("CM")
		} else if hundreds == 4 {
			result.WriteString("CD")
		} else if hundreds >= 5 {
			result.WriteString("D")
			for i := 0; i < hundreds-5; i++ {
				result.WriteString("C")
			}
		} else {
			for i := 0; i < hundreds; i++ {
				result.WriteString("C")
			}
		}
		num -= hundreds * 100

	}

	if num >= 10 {
		tens := num / 10
		if tens == 9 {
			result.WriteString("XC")
		} else if tens == 4 {
			result.WriteString("XL")
		} else if tens >= 5 {
			result.WriteString("L")
			for i := 0; i < tens-5; i++ {
				result.WriteString("X")
			}
		} else {
			for i := 0; i < tens; i++ {
				result.WriteString("X")
			}
		}
		num -= tens * 10
	}

	if num == 9 {
		result.WriteString("IX")
	} else if num == 4 {
		result.WriteString("IV")
	} else if num >= 5 {
		result.WriteString("V")
		for i := 0; i < num-5; i++ {
			result.WriteString("I")
		}
	} else {
		for i := 0; i < num; i++ {
			result.WriteString("I")
		}
	}

	return result.String()
}

type RomanNumeral struct {
	numeric int
	roman   string
}

func intToRoman3(num int) string {
	lookup := []RomanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman strings.Builder
	i := 0
	for num > 0 {
		if num >= lookup[i].numeric {
			roman.WriteString(lookup[i].roman)
			num -= lookup[i].numeric
		} else {
			i++
		}
	}
	return roman.String()
}

func intToRoman2(num int) string {
	roman := strings.Builder{}
	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	for num > 0 {
		// num=1994 i=0, roman = ""
		// num= 994 i=0, roman="M"
		//			i=1
		// num=94   i=1  roman="MCM"
		//			i=2
		//          i=3
		// 			i=4
		// num=4    i=5  romant=MCMXC
		// 			i=6
		//			i=7
		//			i=8..i=10
		// num=0    i=11 roman=MCMXCIV
		if num >= n[i] {
			roman.WriteString(s[i])
			num -= n[i]
		} else {
			i++
		}
	}
	return roman.String()
}

var intToRomanTests = []struct {
	num      int
	expected string
}{
	{3749, "MMMDCCXLIX"},
	{58, "LVIII"},
	{1994, "MCMXCIV"},
}

func TestIntToRoman(t *testing.T) {
	for i, e := range intToRomanTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := intToRoman3(e.num)
			if actual != e.expected {
				t.Fatalf("intToRoman(%v) = %v, want %v", e.num, actual, e.expected)
			}
		})
	}
}

func lengthOfLastWord(s string) int {
	regex := regexp.MustCompile(`\w+`)
	words := regex.FindAllString(s, -1)
	return len(words[len(words)-1])
}
