# Bài 11: Constructor Functions (Hàm Khởi Tạo) Trong Go

> [!NOTE]
> *Chúng ta sẽ tìm hiểu về Constructor Functions (Hàm khởi tạo) - một mẫu thiết kế (design pattern) rất phổ biến trong Go để khởi tạo Struct tập trung và tối ưu hiệu năng bằng cách trả về con trỏ.*

---

## 🛠️ 1. Khái Niệm Constructor Function (Hàm Khởi Tạo)

Go không có các lớp (classes) và do đó cũng không có các hàm khởi tạo (constructors) được tích hợp sẵn dưới dạng cú pháp của ngôn ngữ như Java hay C++.

Thay vào đó, cộng đồng Go sử dụng một quy ước (pattern) gọi là **Constructor Function**:
* Đây là một hàm tiện ích thông thường có nhiệm vụ khởi tạo và trả về đối tượng Struct.
* **Quy ước đặt tên:** Tên hàm thường bắt đầu bằng `new` (hoặc `New` nếu muốn xuất khẩu ra ngoài package) kết hợp với tên Struct. Ví dụ: `newUser` hoặc `NewUser`.

---

## 🏗️ 2. Cách Định Nghĩa Constructor Function

Dưới đây là cách định nghĩa hàm khởi tạo `newUser` nhận các tham số đầu vào và trả về một thực thể của Struct `User`:

```go
func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```

---

## ⚡ 3. Tại Sao Nên Trả Về Con Trỏ (Pointer) Thay Vì Giá Trị?

Trong định nghĩa hàm trên, kiểu trả về là `*User` (con trỏ trỏ tới User) và chúng ta trả về `&User{...}`:

1. **Tránh sao chép bộ nhớ (Avoid Copying):** Nếu hàm trả về một giá trị thông thường (`User`), Go sẽ tạo ra một bản sao của đối tượng đó khi trả về cho hàm gọi. Trả về con trỏ giúp tránh việc sao chép này, tối ưu hóa hiệu năng, đặc biệt là với các Struct lớn.
2. **Khả năng thay đổi dữ liệu:** Đối tượng nhận được dưới dạng con trỏ sẽ dễ dàng được truyền tiếp vào các method sửa đổi (Pointer Receivers) mà không cần lấy địa chỉ lại.

*Lưu ý: Đối với các Struct nhỏ, việc trả về giá trị thông thường vẫn được chấp nhận và không ảnh hưởng nhiều đến hiệu năng. Tuy nhiên, trả về con trỏ là một pattern rất phổ biến.*

---

## 📍 4. Cách Sử Dụng Trong Hàm main()

Thay vì khởi tạo Struct thủ công bằng cú pháp ngoặc nhọn ở nhiều nơi, chúng ta chỉ cần gọi hàm khởi tạo tập trung:

```go
func main() {
	// ... thu thập input
	var appUser *User = newUser(userFirstName, userLastName, userBirthdate)

	// Các method vẫn được gọi bình thường bằng cú pháp dấu chấm
	appUser.outputUserDetails()
	appUser.clearUserName()
}
```

* Nhờ cơ chế tự động giải tham chiếu (auto-dereferencing) của Go, biến `appUser` dù là kiểu con trỏ `*User` vẫn gọi được các method có value receiver như `outputUserDetails` mà không cần viết cú pháp phức tạp.
* Việc đóng gói logic khởi tạo trong `newUser` giúp code của chúng ta sạch sẽ, tái sử dụng tốt hơn và dễ bảo trì khi cấu trúc Struct thay đổi (ví dụ khi thêm trường dữ liệu mặc định như `createdAt`).
