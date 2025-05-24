this is GoLang Training Path\
In Maximilian's course, parts 108 to 114 and 117\
https://www.udemy.com/course/go-the-complete-guide/
### Declaration and Initialization
- 1
```
var anyData [4]any // [<nil> <nil> <nil> <nil>]
numbers := [3]int{1, 2} // Array: [1 2 0]; Length: 3; Capacity: 3
myNumbers := [...]int{10, 20, 30, 40} // Array: [10 20 30 40]; Length: 4; Capacity: 4
```
- 2
```
var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
} // Array: [[1 2 3] [4 5 6]]; Length: 2; Capacity: 2
fmt.Println(matrix[1][2]) // 6
```
- 3
```
var number [4]float64 // [0 0 0 0]
number = [4]float64{2, 4, 5, 6}
number[3] = 100.99
fmt.Println(number, number[2])
```

### Value Semantics
- 1
```
a := [3]int{1, 2, 3}
b := a // copy of a
b[0] = 99

// a => Array: [1 2 3]; Length: 3; Capacity: 3
// b => Array: [99 2 3]; Length: 3; Capacity: 3
```
- 2
```
a := [3]int{1, 2, 3}
b := &a
b[0] = 99

// a => Array: [1 2 3]; Length: 3; Capacity: 3
// b => Array: [99 2 3]; Length: 3; Capacity: 3
```

### Array as Function Parameters
- 1
```
func modify(arr [3]int) {
	arr[0] = 999
}

a := [3]int{1, 2, 3}
modify(a)
// a => Array: [1 2 3]; Length: 3; Capacity: 3
```
- 2
```
func modify(arr *[3]int) {
	arr[0] = 999
}

a := [3]int{1, 2, 3}

modify(&a)
// a => Array: [999 2 3]; Length: 3; Capacity: 3
```
- 3
```
func modify(s []int) {
    s[0] = 100
}

s := []int{0, 1, 2}

modify(s)
// s => Array: [100 1 2]; Length: 3; Capacity: 3
```

### Slice (Dynamic array)
- 1
```
var s []int // Array: []; Length: 0; Capacity: 0
```
- 2
```
prices := []float64{0, 11.11}  // Array: [0 11.11]; Length: 2; Capacity: 2
prices = append(prices, 22.22) // Array: [0 11.11 22.22]; Length: 3; Capacity: 4
prices = append(prices, 33.33) // Array: [0 11.11 22.22 33.33]; Length: 4; Capacity: 4
prices = append(prices, 44.44) // Array: [0 11.11 22.22 33.33 44.44]; Length: 5; Capacity: 8
prices = append(prices, 55.55) // Array: [0 11.11 22.22 33.33 44.44 55.55]; Length: 6; Capacity: 8
prices = prices[:8]            // Array: [0 11.11 22.22 33.33 44.44 55.55 0 0]; Length: 8; Capacity: 8
```
- 3
```
s := []int{0, 1, 2}       // Array: [0 1 2]; Length: 3; Capacity: 3
s = append(s, 3, 4, 5, 6) // Array: [0 1 2 3 4 5 6]; Length: 7; Capacity: 8
VS
s := []int{0, 1, 2} // Array: [0 1 2]; Length: 3; Capacity: 3
s = append(s, 3)    // Array: [0 1 2 3]; Length: 4; Capacity: 6
s = append(s, 4)    // Array: [0 1 2 3 4]; Length: 5; Capacity: 6
s = append(s, 5)    // Array: [0 1 2 3 4 5]; Length: 6; Capacity: 6
s = append(s, 6)    // Array: [0 1 2 3 4 5 6]; Length: 7; Capacity: 12
```

### Slice (From an array)
- 1
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:5] // tooltip VS Code ==> var slicedPrices []float64
slicedPrices[2] = 0         // change index 4(2+2) value in original array and index 2 value in new array
prices[2] = 88              // change index 2 value in original array and index 0(2-2) value in new array
```
- 2
```
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]      // a => Array: [1 2 3 4 5]; Length: 5; Capacity: 5 | s => Array: [2 3 4]; Length: 3; Capacity: 4
s = append(s, 6) // a => Array: [1 2 3 4 6]; Length: 5; Capacity: 5 | s => Array: [2 3 4 6 7 8]; Length: 6; Capacity: 8
s = append(s, 7) // a => Array: [1 2 3 4 6]; Length: 5; Capacity: 5 | a => Array: [2 3 4 6 7]; Length: 5; Capacity: 8
s = append(s, 8) // a => Array: [1 2 3 4 6]; Length: 5; Capacity: 5 | s => Array: [2 3 4 6 7 8]; Length: 6; Capacity: 8
```
- 3 
```
a := [4]int{0, 1, 2, 3} // Array: [0 1 2 3]; Length: 4; Capacity: 4
s := a[:2]              // Array: [0 1]; Length: 2; Capacity: 4
s[1] = 99               // a => Array: [0 99 2 3]; Length: 4; Capacity: 4 | s => Array: [0 99]; Length: 2; Capacity: 4
```
- 4
```
a := [4]int{0, 1, 2, 3}   // Array: [0 1 2 3]; Length: 4; Capacity: 4
s := a[:2]                // Array: [0 1]; Length: 2; Capacity: 4

s = append(s, 99) // a => Array: [0 1 99 3]; Length: 4; Capacity: 4 | s => Array: [0 1 99]; Length: 3; Capacity: 4

s = append(s, 99, 88) // a => Array: [0 1 99 88]; Length: 4; Capacity: 4 | s => Array: [0 1 99 88]; Length: 4; Capacity: 4

s = append(s, 99, 88, 77) // a => Array: [0 1 2 3]; Length: 4; Capacity: 4 | s => Array: [0 1 99 88 77]; Length: 5; Capacity: 8
```
- 5
```
src := []int{0, 1, 2, 3}       // Array: [0 1 2 3]; Length: 4; Capacity: 4
	dst := make([]int, len(src)) // Array: [0 0 0 0]; Length: 4; Capacity: 4
copy(dst, src)                 // dst => Array: [0 1 2 3]; Length: 4; Capacity: 4 | src => Array: [0 1 2 3]; Length: 4; Capacity: 4

src[0] = 99 
dst[1] = 88 
// dst => Array: [0 88 2 3]; Length: 4; Capacity: 4 | src => Array: [99 1 2 3]; Length: 4; Capacity: 4

src = append(src, 99) // dst => Array: [0 1 2 3]; Length: 4; Capacity: 4 | src => Array: [0 1 2 3 99]; Length: 5; Capacity: 8
```
- 6
```
var a []int  // Array: []; Length: 0; Capacity: 0
	b := []int{} // Array: []; Length: 0; Capacity: 0

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
```
- 7 
```
matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	} // Array: [[1 2 3] [4 5 6]]; Length: 2; Capacity: 2

	matrix[1][2] = 999 // Array: [[1 2 3] [4 5 6 99]]; Length: 2; Capacity: 2

	matrix[1] = append(matrix[1], 99) // Array: [[1 2 3] [4 5 6 99]]; Length: 2; Capacity: 2

	matrix = append(matrix, []int{99}) // Array: [[1 2 3] [4 5 6] [99]]; Length: 3; Capacity: 4
```
- 8 
```
i := 2
s := []int{0, 1, 2, 3, 4} // s => Array: [0 1 2 3 4]; Length: 5; Capacity: 5
fmt.Printf("Original ====>\n\n Array: %v; Length: %v; Capacity: %v\n\n\n", s, len(s), cap(s))

s[i] = s[len(s)-1]
s = s[:len(s)-1] // s => Array: [0 1 4 3]; Length: 4; Capacity: 5

s = append(s[:i], s[i+1:]...) // s => Array: [0 1 3 4]; Length: 4; Capacity: 5

import "slices"
s = slices.Delete(s, i, i+1) // s => Array: [0 1 3 4]; Length: 4; Capacity: 5
```
### capacity
- 1
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:7]           // Array: [22.22 33.33 44.44 55.55 66.66]; Length: 5; Capacity: 6
nestedSlicePrice := slicedPrices[2:6] // Array: [44.44 55.55 66.66 77.77]; Length: 4; Capacity: 4
```
- 2
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:]                         // Array: [22.22 33.33 44.44 55.55 66.66 77.77]; Length: 6; Capacity: 6
nestedSlicePrice := slicedPrices[:2]               // Array: [22.22 33.33]; Length: 2; Capacity: 6
nestedSlicePrice = nestedSlicePrice[1:4]           // Array: [33.33 44.44 55.55]; Length: 3; Capacity: 5
nestedSlicePrice = append(nestedSlicePrice, 99.99) // Array: [33.33 44.44 55.55 99.99]; Length: 4; Capacity: 5
```

 
