// package main

// import "fmt"

// func main() {
// 	// Khai báo và khởi tạo một Map chứa cặp Key: Value (string: string)
// 	websites := map[string]string{
// 		"Google":              "https://google.com",
// 		"Amazon Web Services": "https://aws.com",
// 	}

// 	// In ra toàn bộ Map
// 	fmt.Println("Toàn bộ Map ban đầu:", websites)

// 	// 1. Đọc giá trị bằng Key
// 	fmt.Println("\nURL của AWS:", websites["Amazon Web Services"])

// 	// 2. Thêm một cặp Key-Value mới (Map luôn động)
// 	websites["LinkedIn"] = "https://linkedin.com"
// 	fmt.Println("\nMap sau khi thêm LinkedIn:", websites)

// 	// 3. Xóa một cặp Key-Value khỏi Map với hàm delete()
// 	delete(websites, "Google")
// 	fmt.Println("\nMap sau khi xóa Google:", websites)
// }

