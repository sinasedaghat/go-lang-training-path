// https://www.codewars.com/kata/5592e3bd57b64d00f3000047/go
package algorithmSeries

// import "fmt"

// first try
func FindNb(m int) int {
	nc := 0 // number of cubes

	for {
		nc++
		m -= (nc * nc * nc)
		if m <= 0 {
			break
		}
	}

	if m == 0 {
		return nc
	} else {
		return -1
	}
}
