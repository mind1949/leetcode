package algs

/*
letterCombinations sloves the following problem:
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
	{
		2: []{"a", "b", "c"},
		3: []{"d", "e", "f"},
		4: []{"g", "h", "i"},
		5: []{"j", "k", "l"},
		6: []{"m", "n", "o"},
		7: []{"p", "q", "r", "s"},
		8: []{t", "u", "v"},
		9: []{w", "x", "y", "x"},
	}
Example:
	Input: "23"
	Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combs := []string{""}
	for i := 0; i < len(digits); i++ {
		combsLen := len(combs)
		for j := 0; j < combsLen; j++ {
			suffix := combs[j]
			for k, letter := range m[digits[i]-'2'] {
				v := suffix + string(letter)
				if k == 0 {
					combs[j] = v
				} else {
					combs = append(combs, v)
				}
			}
		}
	}

	return combs
}
