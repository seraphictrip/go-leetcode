package random

/*
In a string composed of 'L', 'R', and 'X' characters, like "RXXLRXRXL", a move consists of either replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR". Given the starting string start and the ending string end, return True if and only if there exists a sequence of moves to transform start to end.



Example 1:

Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
Output: true
Explanation: We can transform start to end following these steps:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX
Example 2:

Input: start = "X", end = "L"
Output: false


Constraints:

1 <= start.length <= 104
start.length == end.length
Both start and end will only consist of characters in 'L', 'R', and 'X'.
*/

func CanTransform(start string, end string) bool {
	n := len(start)
	// sliding window
	i := 0
	j := 1

	for i < n && j < n {
		need, can := canTransform(start[i:j+1], end[i:j+1])
		if !can {
			return false
		}
		if need {
			start = swap(start, i, j)
		}

		i += 2
		j += 2
	}
	return start[i:] == end[i:]
}

func swap(str string, i, j int) string {
	bytes := []byte(str)
	bytes[i], bytes[j] = bytes[j], bytes[i]
	return string(bytes)
}

func canTransform(s, e string) (need bool, ok bool) {
	if s == e {
		return false, true
	}
	switch s {
	case "LX":
		return true, e == "XL"
	case "XL":
		return true, e == "LX"
	case "RX":
		return true, e == "XR"
	case "XR":
		return true, e == "RX"
	default:
		return false, false
	}
}
