# Bài 09: Đính Kèm Phương Thức (Methods) & Tham Số Nhận (Receiver Arguments)

> [!NOTE]
> *Chúng ta sẽ học cách biến một hàm độc lập thành một **Method** gắn liền với một Struct bằng cách sử dụng **Receiver Argument (Tham số nhận)**.*

---

## ⚙️ 1. Khái Niệm Method (Phương Thức) Trong Go

Trong Go, một **Method** thực chất là một hàm được đính kèm vào một kiểu dữ liệu cụ thể (thường là một Struct). 

* Nó giúp đóng gói dữ liệu và các hành vi xử lý dữ liệu đó lại cùng một chỗ, giúp code có tổ chức và hướng đối tượng (OOP-like) hơn.
* Ví dụ, thay vì có một hàm độc lập `outputUserDetails(u User)` nhận vào đối tượng User, ta có thể đính kèm hàm đó trực tiếp vào struct `User` để bất kỳ đối tượng User nào cũng có thể tự xuất thông tin của chính mình.

---

## 👈 2. Khai Báo Receiver Argument (Tham Số Nhận)

Để biến một hàm bình thường thành một method của Struct, chúng ta thêm một cặp dấu ngoặc đơn chứa **Receiver Argument (Tham số nhận)** ở giữa từ khóa `func` và tên hàm:

```go
// u là tên biến đại diện cho đối tượng nhận method
// User là kiểu dữ liệu Struct nhận method này
func (u User) outputUserDetails() {
	// Truy cập các trường thông qua biến receiver u
	fmt.Println("First Name:", u.firstName)
}
```

* **Ưu điểm:** Hàm bây giờ không cần phải định nghĩa bất kỳ tham số đầu vào nào trong phần `()` sau tên hàm nữa. Toàn bộ thông tin của Struct đã được cung cấp sẵn thông qua biến receiver `u`.

---

## 📍 3. Cách Gọi Method (Dot Notation)

Khi một hàm đã biến thành method của Struct, ta **không thể** gọi nó như một hàm độc lập thông thường được nữa:

```go
// LỖI: Không thể gọi trực tiếp
outputUserDetails(appUser) 
```

Thay vào đó, ta sử dụng cú pháp dấu chấm (`.`) gọi trực tiếp trên instance của Struct đó:

```go
// Hợp lệ: Gọi method trên đối tượng appUser
appUser.outputUserDetails()
```

Khi dòng code trên chạy, Go sẽ tự động ngầm hiểu và truyền thực thể `appUser` vào làm giá trị cho biến receiver `u` trong định nghĩa của method.
