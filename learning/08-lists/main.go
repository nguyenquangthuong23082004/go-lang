package main

import "fmt"

// Định nghĩa Custom Type (Type Alias) cho map[string]float64
type floatMap map[string]float64


// Thêm custom method cho type alias floatMap
func (m floatMap) output() {
	fmt.Println("Custom Output floatMap:", m)
}

func main() {
	// -------------------------------------------------------------
	// 1. Khởi tạo Slice với hàm make() để tối ưu hóa bộ nhớ
	// -------------------------------------------------------------
	// make([]Kiểu, len, cap)
	// Tạo slice có len=2 (2 ô rỗng ban đầu) và cấp phát sẵn cap=5 bộ nhớ
	userNames := make([]string, 2, 5)

	// Gán giá trị trực tiếp cho 2 ô rỗng đã được khởi tạo theo index
	userNames[0] = "Max"
	userNames[1] = "Manuel"

	// -------------------------------------------------------------
	// 2. Dùng append() khi cap vẫn còn đủ chỗ
	// -------------------------------------------------------------
	// Vì cap = 5, đã dùng 2 ô (len=2), nên khi append tiếp Go KHÔNG cần
	// cấp phát lại bộ nhớ mới mà dùng trực tiếp sức chứa còn dư!
	userNames = append(userNames, "Julie")
	userNames = append(userNames, "Chris")

	fmt.Println("Usernames:", userNames)
	fmt.Printf("len: %d, cap: %d\n", len(userNames), cap(userNames))

	// -------------------------------------------------------------
	// 3. Khởi tạo Map với hàm make() và Custom Type (floatMap)
	// -------------------------------------------------------------
	// Dùng custom type floatMap ngắn gọn thay vì gõ map[string]float64
	courseRatings := make(floatMap, 3)

	// Thêm các cặp Key-Value
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5

	// Gọi method tùy chỉnh trực tiếp trên custom type
	courseRatings.output()

	// -------------------------------------------------------------
	// 4. Vòng lặp `for range` duyệt qua Slice
	// -------------------------------------------------------------
	fmt.Println("\n--- Vòng lặp for range qua Slice ---")
	for index, value := range userNames {
		fmt.Printf("Index: %d | Value: %s\n", index, value)
	}

	// -------------------------------------------------------------
	// 5. Vòng lặp `for range` duyệt qua Map
	// -------------------------------------------------------------
	fmt.Println("\n--- Vòng lặp for range qua Map ---")
	for key, value := range courseRatings {
		fmt.Printf("Key: %s | Value: %.1f\n", key, value)
	}
}




