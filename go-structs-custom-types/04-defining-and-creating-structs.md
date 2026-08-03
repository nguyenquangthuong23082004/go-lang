# Bài 04: Định Nghĩa & Khởi Tạo Struct Trong Go

> [!NOTE]
> *Chúng ta sẽ học cách định nghĩa một Struct (kiểu dữ liệu cấu trúc tùy chỉnh) và các cách khởi tạo (instantiate) một đối tượng cụ thể từ Struct đó trong Go.*

---

## 🛠️ 1. Cách Định Nghĩa Struct (Defining a Struct)

Cú pháp khai báo một Struct sử dụng hai từ khóa `type` và `struct`:

```go
type User struct {
	firstName string
	lastName  string
	birthdate string
}
```

Trong đó:
* **`User`**: Tên của kiểu dữ liệu cấu trúc (thường viết hoa chữ cái đầu).
* **`firstName`, `lastName`, `birthdate`**: Các trường (fields) của struct kèm theo kiểu dữ liệu tương ứng.

---

## 🏗️ 2. Các Cách Khởi Tạo Struct (Instantiating a Struct)

Go cung cấp một vài cách để tạo ra một instance từ một Struct:

### Cách 1: Sử dụng Key-Value (Khuyên dùng)
Gán trực tiếp giá trị kèm tên trường. Cách này rõ ràng nhất và không bị ảnh hưởng nếu sau này ta thay đổi thứ tự khai báo các trường trong struct.
```go
var appUser User = User{
	firstName: userFirstName,
	lastName:  userLastName,
	birthdate: userBirthdate,
}
```

### Cách 2: Khởi tạo theo vị trí (Positional Struct)
Bỏ qua tên trường, truyền trực tiếp giá trị theo đúng thứ tự đã khai báo trong struct.
```go
appUser := User{userFirstName, userLastName, userBirthdate}
```
> [!WARNING]
> *Hạn chế dùng cách này cho các struct lớn vì nếu thay đổi thứ tự các trường hoặc thêm trường mới, code sẽ bị lỗi biên dịch hoặc gán sai giá trị.*

### Cách 3: Sử dụng giá trị mặc định (Zero Value Struct)
Khi tạo struct không truyền giá trị, các trường sẽ nhận giá trị mặc định của kiểu dữ liệu đó (`""` cho string, `0` cho số, `false` cho bool).
```go
var appUser User // Tất cả các trường là ""
// Gán giá trị sau đó bằng dấu chấm (.)
appUser.firstName = userFirstName
```

---

## 💻 3. Áp Dụng Struct Vào Dự Án (`structs.go`)

Dưới đây là mã nguồn dự án sau khi refactor sử dụng Struct `User`:

```go
package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	birthdate string
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
	}

	// Gọi hàm xuất dữ liệu người dùng bằng cách truyền Struct
	outputUserDetails(appUser)
}

func outputUserDetails(user User) {
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
	fmt.Println("Birthdate:", user.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
```
