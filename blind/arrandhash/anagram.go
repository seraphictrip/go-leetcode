package arrandhash

import "slices"

func isAnagramSort(s string, t string) bool {
	sprime := []byte(s)
	tprime := []byte(t)
	slices.Sort(sprime)
	slices.Sort(tprime)
	return string(sprime) == string(tprime)

}

func isAnagramHashMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make(map[byte]int, len(s))
	for i := range s {
		seen[s[i]]++
	}

	for i := range t {
		if seen[t[i]] == 0 {
			return false
		}
		seen[t[i]]--
	}
	return true
}

// This seems to run significantly faster
// I will probably continue to code using map, but this is an interesting thing to bring
// up if asked
func isAnagramFixedSizeArray(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := [26]int{}
	for i := range s {
		seen[s[i]-'a']++
	}

	for i := range t {
		if seen[t[i]-'a'] == 0 {
			return false
		}
		seen[t[i]-'a']--
	}
	return true
}
