package algs

/*
GenerateParenthesis solves the following problem:

	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	For example, given n = 3, a solution set is:

		[
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"
		]
*/
func GenerateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	combs := []string{"("}
	for i := 1; i < n*2; i++ {
		for l := len(combs); l > 0; l-- {
			// combs[l-1]中的左括号与右括号数量
			var (
				lp, rp int
			)
			for _, p := range combs[l-1] {
				switch p {
				case '(':
					lp++
				case ')':
					rp++
				default:
					panic("invalid character: " + string(p))
				}
			}

			// 确定可能的组合
			if lp-rp > 0 {
				if lp < n {
					combs = append(combs, combs[l-1]+")")
					combs[l-1] += "("
				} else {
					combs[l-1] += ")"
				}
			} else {
				combs[l-1] += "("
			}
		}
	}

	return combs
}
