package main

import "fmt"

// Định nghĩa Custom Type đại diện cho kiểu Hàm (Function Type Alias)
// Nhận vào 1 tham số int và trả về 1 kết quả int
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	// Truyền hàm `double` vào làm tham số cho `transformNumbers`
	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled Numbers:", doubled) // Output: [2 4 6 8]

	// Truyền hàm `triple` vào làm tham số cho `transformNumbers`
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled Numbers:", tripled) // Output: [3 6 9 12]
}

// Hàm biến đổi chung (Higher-Order Function)
// Nhận vào một slice numbers và một hàm biến đổi `transform`
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		// Gọi hàm `transform` được truyền vào như một hàm bình thường
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Các hàm Helper tuân thủ kiểu `transformFn` (func(int) int)
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
