# Bài 16: Struct Embedding (Nhúng Struct) & Kế Thừa Trong Go

> [!NOTE]
> *Chúng ta sẽ khám phá tính năng **Struct Embedding (Nhúng Struct)** trong Go. Đây là cách Go kế thừa thuộc tính và phương thức từ một Struct khác mà không cần sử dụng cơ chế class truyền thống.*

---

## 🧬 1. Khái Niệm Struct Embedding (Nhúng Struct)

Trong các ngôn ngữ lập trình hướng đối tượng (OOP) truyền thống như Java hay C#, bạn sử dụng từ khóa `extends` hoặc dấu hai chấm `:` để kế thừa thuộc tính và phương thức từ lớp cha.

Go không có lớp (class) và không hỗ trợ kế thừa truyền thống. Thay vào đó, Go sử dụng một cơ chế mạnh mẽ và linh hoạt hơn gọi là **Composition (Kết hợp)** thông qua **Struct Embedding (Nhúng Struct)**:
* Chúng ta định nghĩa một Struct mới xây dựng dựa trên một Struct đã tồn tại.
* Struct mới sẽ tự động thừa hưởng toàn bộ các trường (fields) và các phương thức (methods) của Struct được nhúng vào.

---

## 🏗️ 2. Nhúng Ẩn Danh (Anonymous Embedding)

Khi khai báo trường được nhúng trong Struct mới, chúng ta không cần đặt tên cho trường đó (chỉ khai báo tên kiểu dữ liệu):

```go
type Admin struct {
	email    string
	password string
	User     // Nhúng ẩn danh struct User
}
```

### Lợi ích cực lớn của nhúng ẩn danh:
Khi bạn sử dụng nhúng ẩn danh, toàn bộ các trường và phương thức của Struct được nhúng (`User`) sẽ được **quảng bá (promoted)** trực tiếp lên Struct bên ngoài (`Admin`).
* Bạn có thể gọi trực tiếp các phương thức của `User` ngay trên thực thể `Admin` mà không cần thông qua đối tượng trung gian.

```go
// Thay vì viết dài dòng:
admin.User.OutputUserDetails()

// Bạn có thể viết trực tiếp:
admin.OutputUserDetails() // Rất ngắn gọn và tự nhiên
```

---

## 🛠️ 3. Khởi Tạo Struct Nhúng

Mặc dù lúc sử dụng ta có thể gọi trực tiếp (nhờ cơ chế quảng bá), nhưng khi khởi tạo Struct bằng Struct Literal, bạn **vẫn bắt buộc** phải chỉ rõ tên kiểu dữ liệu nhúng làm tên trường:

```go
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{ // Bắt buộc phải viết rõ trường User ở đây
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
```

---

## 💻 4. Mã Nguồn Thực Tế Minh Họa

### Trong file `user.go`:
[user/user.go](file:///home/uwal/Desktop/go-lang/learning/06-structs/user/user.go) chứa khai báo Struct `Admin` nhúng `User` và constructor `NewAdmin`.

### Trong file `structs.go`:
Chúng ta khởi tạo `Admin` và gọi các method của `User` trực tiếp trên đối tượng `admin`:

```go
func main() {
	// ... Khởi tạo appUser thông thường

	// Tạo đối tượng admin mới
	admin := user.NewAdmin("admin@example.com", "supersecret")

	fmt.Println("\n--- ADMIN DETAILS ---")
	// Gọi các phương thức của User trực tiếp trên admin
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
}
```
