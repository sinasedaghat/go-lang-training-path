// https://www.codewars.com/kata/5a662a02e626c54e87000123/go
package numbersSeries

func ExtraPerfect(n int) []int {
	var result = []int{}
	for number := range n + 1 {
		// for number := 0; number <= n; number++ {
		if number%2 == 1 {
			result = append(result, number)
		}
	}

	return result
}
