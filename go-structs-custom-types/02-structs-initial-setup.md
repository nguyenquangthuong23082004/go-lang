# Bài 02: Khởi Tạo Dự Án & Mã Nguồn Ban Đầu (Chưa Dùng Structs)

> [!NOTE]
> *Để chuẩn bị tìm hiểu về Structs, chúng ta sẽ bắt đầu bằng cách dựng cấu trúc dự án mới và viết mã nguồn nhập/xuất dữ liệu người dùng bằng các biến đơn lẻ thông thường.*

---

## 📂 1. Cấu Trúc Thư Mục Dự Án Mới

Tạo một thư mục dự án độc lập mới tại `/learning/06-structs/` với hai tệp tin:
1. **`go.mod`**: Khai báo tên module của dự án:
   ```go
   module example.com/structs
   
   go 1.22
   ```
2. **`structs.go`**: Triển khai mã nguồn cơ bản sử dụng package `main` thu thập thông tin người dùng.

---

## 💻 2. Mã Nguồn Ban Đầu (`structs.go`)

Mã nguồn truyền thống sử dụng các biến đơn lẻ để lưu trữ thông tin:
```go
package main

import "fmt"

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// In ra dữ liệu đã thu thập
	fmt.Println("First Name:", userFirstName)
	fmt.Println("Last Name:", userLastName)
	fmt.Println("Birthdate:", userBirthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
```

---

## ⚠️ 3. Vấn Đề Gặp Phải Khi Không Sử Dụng Structs

Khi chương trình phát triển lớn hơn, việc quản lý các biến rời rạc như `userFirstName`, `userLastName`, `userBirthdate` sẽ mang lại nhiều bất cập:

1. **Khó Truyền Nhận Dữ Liệu:** Nếu bạn muốn viết một hàm để lưu thông tin người dùng vào file hoặc gửi qua mạng, bạn sẽ phải truyền riêng rẽ từng tham số:
   ```go
   func saveUser(firstName string, lastName string, birthdate string) { ... }
   ```
   Nếu cần thêm 5 trường thông tin khác, chữ ký hàm (function signature) sẽ trở nên quá dài và khó kiểm soát.
2. **Không Có Tính Đóng Gói (Lack of Encapsulation):** Các biến không có bất kỳ ràng buộc logic nào thể hiện rằng chúng thuộc cùng một thực thể (Entity) duy nhất đại diện cho một `User`.
3. **Khó Khăn Trong Bảo Trì:** Việc mở rộng thuộc tính mới buộc ta phải cập nhật thủ công tất cả các hàm và biến ở mọi nơi liên quan.
