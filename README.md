this is GoLang Training Path\
In Maximilian's course, parts 108 to 114\
https://www.udemy.com/course/go-the-complete-guide/
### basic
```
var anyData [4]any // [<nil> <nil> <nil> <nil>]

var number [4]float64 // [0 0 0 0]
number = [4]float64{2, 4, 5, 6}
number[3] = 100.99
fmt.Println(number, number[2])
```
### slice
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:5] // tooltip VS Code ==> var slicedPrices []float64
slicedPrices[2] = 0         // change index 4(2+2) value in original array and index 2 value in new array
prices[2] = 88              // change index 2 value in original array and index 0(2-2) value in new array
```
### capacity
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:7] // Array: [22.22 33.33 44.44 55.55 66.66]; Length: 5; Capacity: 6
nestedSlicePrice := slicedPrices[2:6] // Array: [44.44 55.55 66.66 77.77]; Length: 4; Capacity: 4
```
```
prices := [8]float64{00.00, 11.11, 22.22, 33.33, 44.44, 55.55, 66.66, 77.77}
slicedPrices := prices[2:] // Array: [22.22 33.33 44.44 55.55 66.66 77.77]; Length: 6; Capacity: 6
nestedSlicePrice := slicedPrices[:2] // Array: [22.22 33.33]; Length: 2; Capacity: 6
nestedSlicePrice = nestedSlicePrice[1:4] // Array: [33.33 44.44 55.55]; Length: 3; Capacity: 5
nestedSlicePrice = append(nestedSlicePrice, 99.99) // Array: [33.33 44.44 55.55 99.99]; Length: 4; Capacity: 5
```
### dynamic array
```
prices := []float64{0, 11.11} // Array: [0 11.11]; Length: 2; Capacity: 2
prices = append(prices, 22.22) // Array: [0 11.11 22.22]; Length: 3; Capacity: 4
prices = append(prices, 33.33) // Array: [0 11.11 22.22 33.33]; Length: 4; Capacity: 4
prices = append(prices, 44.44) // Array: [0 11.11 22.22 33.33 44.44]; Length: 5; Capacity: 8
prices = append(prices, 55.55) // Array: [0 11.11 22.22 33.33 44.44 55.55]; Length: 6; Capacity: 8
prices = prices[:8] // Array: [0 11.11 22.22 33.33 44.44 55.55 0 0]; Length: 8; Capacity: 8
```
 
