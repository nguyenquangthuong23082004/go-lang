# Bài 14: Tách Struct Sang Package Riêng & Quy Tắc Xuất Khẩu (Exporting)

> [!NOTE]
> *Chúng ta sẽ học cách tổ chức mã nguồn bằng cách tách Struct, các phương thức và hàm khởi tạo sang một Package con riêng biệt, đồng thời tìm hiểu quy tắc viết hoa để xuất khẩu các trường của Struct.*

---

## 📂 1. Cấu Trúc Thư Mục Khi Tách Package

Để giữ cho dự án sạch sẽ và dễ tái sử dụng, chúng ta tách toàn bộ logic liên quan đến người dùng sang một package con tên là `user`:

```
/learning/06-structs/
├── go.mod
├── structs.go  (Gói main - Điểm chạy chính)
└── user/
    └── user.go (Gói user - Định nghĩa cấu trúc User)
```

---

## 🔑 2. Quy Tắc Xuất Khẩu (Exporting Rules) Cho Struct & Các Trường

Quy tắc chữ cái viết hoa đầu tiên áp dụng cho toàn bộ các thành phần trong Go khi làm việc liên package:

### A. Xuất khẩu Tên Struct
Để package `main` có thể khai báo biến kiểu `user.User`, tên của struct bắt buộc phải viết hoa chữ cái đầu tiên:
```go
type User struct { ... } // Viết hoa chữ U
```

### B. Xuất khẩu Các Trường Bên Trong Struct (Fields)
Một điểm cực kỳ quan trọng và dễ nhầm lẫn: **Việc xuất khẩu Struct không đồng nghĩa với việc xuất khẩu các trường bên trong nó.**
* Nếu bạn khai báo trường viết thường: `firstName string` -> Trường này là **Private (Unexported)**, gói ngoài không thể đọc hay ghi giá trị vào trường này.
* Nếu bạn muốn gói ngoài truy cập được: Bắt buộc viết hoa chữ cái đầu tiên: `FirstName string`.

```go
type User struct {
	FirstName string    // Public (Exported)
	LastName  string    // Public (Exported)
	Birthdate string    // Public (Exported)
	CreatedAt time.Time // Public (Exported)
}
```

### C. Xuất khẩu Constructor & Methods
Tương tự, hàm khởi tạo và các phương thức cũng phải viết hoa chữ cái đầu để bên ngoài có thể gọi:
* `newUser` -> Đổi thành `NewUser`
* `outputUserDetails` -> Đổi thành `OutputUserDetails`
* `clearUserName` -> Đổi thành `ClearUserName`

---

## 📦 3. Cách Sử Dụng Tại `structs.go`

Tại gói `main`, chúng ta import gói `user` thông qua đường dẫn module:

```go
package main

import (
	"fmt"
	"example.com/structs/user" // Import package con
)

func main() {
	// ... Nhận dữ liệu input
	
	// Gọi NewUser từ package user
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails() // Gọi method đã xuất khẩu
}
```
