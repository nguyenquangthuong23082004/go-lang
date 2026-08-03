# Bài 13: Ràng Buộc Dữ Liệu (Validation) & Xử Lý Lỗi Trong Constructor

> [!NOTE]
> *Chúng ta sẽ học cách thêm logic kiểm tra tính hợp lệ của dữ liệu đầu vào (Validation) trực tiếp bên trong Constructor Function, đồng thời trả về lỗi (error) nếu dữ liệu không hợp lệ.*

---

## 🎯 1. Tại Sao Nên Viết Validation Trong Constructor?

Constructor Function (`newUser`) đóng vai trò là cổng kiểm soát duy nhất chịu trách nhiệm tạo ra thực thể Struct trong ứng dụng:

* **Tập trung hóa logic (Centralization):** Giúp bạn kiểm tra dữ liệu đầu vào tại một nơi duy nhất. Bạn không cần phải viết đi viết lại code kiểm tra ở tất cả những nơi khởi tạo người dùng trong chương trình.
* **Ngăn chặn dữ liệu rác:** Đảm bảo không bao giờ có một đối tượng `User` không hợp lệ (ví dụ tên bị trống) được tạo ra và sử dụng trong hệ thống.

---

## 🛠️ 2. Thay Đổi Chữ Ký Hàm Constructor Để Trả Về Lỗi

Để trả về lỗi khi kiểm tra không đạt yêu cầu, chúng ta cần thay đổi kiểu trả về của hàm `newUser` thành **2 giá trị**:
1. Con trỏ Struct `*User` (đối tượng được tạo).
2. Đối tượng lỗi `error` (thông báo lỗi nếu khởi tạo thất bại).

```go
import "errors"

func newUser(firstName, lastName, birthdate string) (*User, error) {
	// Kiểm tra nếu có bất kỳ trường nào bị để trống
	if firstName == "" || lastName == "" || birthdate == "" {
		// Trả về nil cho đối tượng và lỗi mới bằng errors.New
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	// Nếu hợp lệ, trả về đối tượng và nil cho lỗi
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
```

---

## ⚙️ 3. Xử Lý Lỗi Tại Nơi Gọi Khởi Tạo (main.go)

Tại hàm `main()`, chúng ta nhận cả hai giá trị trả về và tiến hành kiểm tra lỗi thông qua biến `err`:

```go
func main() {
	// ... Nhận input từ bàn phím
	
	// Gọi constructor và kiểm tra lỗi
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println(err) // In ra thông báo lỗi
		return           // Dừng chương trình vì không thể tiếp tục nếu thiếu user
	}

	// Dữ liệu hợp lệ, tiếp tục xử lý
	appUser.outputUserDetails()
}
```

---

## 💡 4. Phân Biệt `fmt.Scan` và `fmt.Scanln` Khi Nhập Trống

Trong thực tế, khi bạn muốn kiểm tra hành vi nhập trống từ bàn phím:
* Nếu sử dụng **`fmt.Scan(&value)`**: Nếu bạn chỉ nhấn phím `Enter` mà không nhập ký tự nào, chương trình sẽ tiếp tục đứng chờ bạn nhập tiếp (không kết thúc dòng nhập).
* Nếu sử dụng **`fmt.Scanln(&value)`**: Nhấn phím `Enter` sẽ kết thúc dòng nhập ngay lập tức và truyền giá trị rỗng `""` vào biến, giúp logic kiểm tra trống hoạt động chính xác.
