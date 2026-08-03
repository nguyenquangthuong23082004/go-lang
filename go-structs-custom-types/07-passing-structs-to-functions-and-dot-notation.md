# Bài 07: Truyền Struct Vào Hàm & Sử Dụng Cú Pháp Dấu Chấm (Dot Notation)

> [!NOTE]
> *Chúng ta sẽ học cách truyền một Struct làm tham số cho hàm và cách sử dụng cú pháp dấu chấm (`.`) để truy cập dữ liệu bên trong các trường của Struct đó.*

---

## 📤 1. Truyền Struct Làm Đối Số (Passing Struct as Argument)

Một trong những ưu điểm lớn nhất của Struct là giúp thu gọn danh sách đối số truyền vào hàm. Thay vì phải truyền từng biến đơn lẻ như trước:

```go
// Cách cũ: dài dòng và dễ nhầm thứ tự
outputUserDetails(userFirstName, userLastName, userBirthdate)
```

Chúng ta chỉ cần truyền duy nhất một thực thể chứa toàn bộ dữ liệu:

```go
// Cách mới: gọn gàng và an toàn hơn
outputUserDetails(appUser)
```

---

## 👈 2. Định Nghĩa Tham Số Với Kiểu Struct (Parameter Definition)

Để nhận vào một Struct làm tham số, trong chữ ký hàm (function signature), chúng ta chỉ định kiểu dữ liệu là tên Struct đã định nghĩa:

```go
func outputUserDetails(u User) {
	// u là tên tham số (parameter name)
	// User là kiểu dữ liệu (struct type)
}
```

*Trong ví dụ trên, tham số được đặt tên ngắn gọn là `u` để tránh bị trùng lặp hoặc nhầm lẫn với tên kiểu dữ liệu `User`.*

---

## 📍 3. Truy Cập Trường Bằng Cú Pháp Dấu Chấm (Dot Notation)

Để đọc hoặc thay đổi dữ liệu của các trường trong Struct, chúng ta sử dụng **cú pháp dấu chấm (`.`)**:

* Tương tự như cách bạn dùng dấu chấm để truy cập vào các hàm/biến được xuất khẩu từ một package (ví dụ: `fmt.Println`, `time.Now`).
* Cú pháp: `<tên_biến_struct>.<tên_trường>`

Ví dụ:
```go
fmt.Println("First Name:", u.firstName)
fmt.Println("Last Name:", u.lastName)
fmt.Println("Birthdate:", u.birthdate)
```

---

## 💻 4. Mã Nguồn Thực Tế (`structs.go`)

Đoạn code hoàn chỉnh thể hiện việc truyền Struct và sử dụng dot notation:

```go
package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// Gọi hàm xuất dữ liệu bằng cách truyền Struct
	outputUserDetails(appUser)
}

func outputUserDetails(u User) {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
```
