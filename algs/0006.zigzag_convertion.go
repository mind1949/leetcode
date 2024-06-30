package algs

/*
Convert solves the following problem:

	The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

	P   A   H   N
	A P L S I I G
	Y   I   R

	And then read line by line: "PAHNAPLSIIGYIR"

	Write the code that will take a string and make this conversion given a number of rows:

	string convert(string s, int numRows);

	Example 1:

		Input: s = "PAYPALISHIRING", numRows = 3
		Output: "PAHNAPLSIIGYIR"

	Example 2:

		Input: s = "PAYPALISHIRING", numRows = 4
		Output: "PINALSIGYAHRPI"
		Explanation:

		P     I    N
		A   L S  I G
		Y A   H R
		P     I
*/
func Convert(s string, numRows int) string {
	length := len([]rune(s))
	if length <= numRows || numRows == 1 {
		return s
	}

	row := 0
	step := -1
	list := make([][]rune, numRows)
	for _, v := range s {
		list[row] = append(list[row], v)
		if row == numRows-1 || row == 0 {
			step = step * (-1)
		}
		row += step
	}

	result := make([]rune, 0, length)
	for _, v := range list {
		result = append(result, v...)
	}
	return string(result)
}
