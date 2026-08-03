# Bài 15: Đóng Gói (Encapsulation) Với Trường Private & Pattern user.New

> [!NOTE]
> *Chúng ta sẽ tìm hiểu cách bảo vệ dữ liệu (Đóng gói) bằng cách biến các trường của Struct thành Private (không xuất khẩu), bắt buộc khởi tạo thông qua Constructor và tìm hiểu quy ước đặt tên ngắn gọn là `New`.*

---

## 🔒 1. Ẩn Thuộc Tính (Private Fields) Để Bảo Vệ Dữ Liệu

Trong nhiều trường hợp thực tế, chúng ta không muốn các package bên ngoài tự do đọc hoặc sửa đổi các thuộc tính bên trong Struct (vì có thể làm dữ liệu bị sai lệch, không nhất quán).

Để làm điều này, chúng ta áp dụng tính chất **Đóng gói (Encapsulation)** bằng cách viết thường chữ cái đầu tiên của các trường:

```go
type User struct {
	firstName string    // Private (Unexported)
	lastName  string    // Private (Unexported)
	birthdate string    // Private (Unexported)
	createdAt time.Time // Private (Unexported)
}
```

### Hệ quả:
Các package bên ngoài (như `package main`) sẽ không thể đọc hay ghi đè thuộc tính này nữa:
```go
// LỖI BIÊN DỊCH: firstName là thuộc tính unexported
appUser.firstName = "Max" 
```

---

## 🏗️ 2. Bắt Buộc Sử Dụng Constructor Cho Việc Khởi Tạo

Khi các trường của Struct là private, bên ngoài cũng **không thể** tự khởi tạo bằng Struct Literal:

```go
// LỖI BIÊN DỊCH: Không thể truyền dữ liệu cho các trường private
appUser := user.User{
	firstName: "John", 
}
```

Lúc này, cách duy nhất để tạo ra một đối tượng `User` hợp lệ là gọi thông qua Constructor Function đã được xuất khẩu công khai. Điều này đảm bảo mọi thực thể `User` được sinh ra đều bắt buộc phải đi qua các bước Validation của Constructor.

---

## 🌟 3. Quy Ước Đặt Tên Constructor Ngắn Gọn: `New`

Khi Struct nằm trong một package con độc lập, thay vì đặt tên hàm khởi tạo là `NewUser`:

```go
// Cách cũ: Bị lặp từ "user" hai lần
appUser, err := user.NewUser(firstName, lastName, birthdate)
```

Quy ước chuẩn và phổ biến trong cộng đồng Go là đổi tên hàm khởi tạo chỉ đơn thuần thành **`New`**:

```go
// Cách mới: Ngắn gọn, súc tích và cực kỳ tự nhiên
appUser, err := user.New(firstName, lastName, birthdate)
```

### Ví dụ thực tế từ Thư viện chuẩn Go:
* Package `errors` cung cấp hàm `errors.New()` để tạo lỗi mới (không phải `errors.NewError()`).
* Package `bytes` cung cấp hàm `bytes.NewBuffer()` (được rút gọn hợp lý khi khởi tạo Buffer).
* Quy ước đặt tên `New` giúp cú pháp của Go trở nên ngắn gọn và mang đậm phong cách thiết kế của ngôn ngữ (Idiomatic Go).
