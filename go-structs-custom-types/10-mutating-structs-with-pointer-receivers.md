# Bài 10: Thay Đổi Dữ Liệu Struct & Pointer Receivers

> [!NOTE]
> *Chúng ta sẽ tìm hiểu lý do tại sao Value Receivers không thể thay đổi dữ liệu của Struct gốc, từ đó học cách sử dụng **Pointer Receivers** để chỉnh sửa trực tiếp giá trị của Struct.*

---

## ❌ 1. Vấn Đề Với Value Receiver (Tham Số Nhận Dạng Giá Trị)

Tham số nhận dạng giá trị (**Value Receiver**) hoạt động tương tự như việc truyền tham trị vào một hàm bình thường trong Go:

* Khi gọi method, Go tự động tạo một **bản sao (copy)** của Struct trong bộ nhớ để truyền vào method.
* Nếu chúng ta thay đổi thuộc tính của Struct bên trong method, chúng ta chỉ đang thay đổi trên bản sao đó.

Ví dụ:
```go
func (u User) clearUserName() {
	u.firstName = "" // Lỗi logic: chỉ xóa tên trên bản sao u
	u.lastName = ""
}
```

Kết quả là đối tượng gốc `appUser` bên ngoài hàm `main` vẫn giữ nguyên giá trị cũ, không hề bị thay đổi.

---

## 👉 2. Giải Pháp: Sử Dụng Pointer Receiver (Tham Số Nhận Dạng Con Trỏ)

Để cho phép một method thay đổi trực tiếp các trường dữ liệu của Struct gốc, chúng ta phải khai báo receiver dưới dạng một con trỏ (**Pointer Receiver**) bằng cách thêm dấu hoa thị `*`:

```go
func (u *User) clearUserName() {
	u.firstName = "" // Thay đổi trực tiếp trên Struct gốc
	u.lastName = ""
}
```

Lúc này, thay vì sao chép toàn bộ Struct, Go truyền địa chỉ ô nhớ của Struct gốc vào method. Mọi thay đổi bên trong method sẽ tác động trực tiếp lên Struct gốc.

---

## 🌟 3. Cú Pháp Gọi Method Tiện Lợi Trong Go

Thông thường, đối với các hàm nhận con trỏ, bạn phải truyền địa chỉ bằng toán tử `&` (ví dụ: `outputUserDetails(&appUser)`).

Tuy nhiên, đối với Method trong Go, bạn **không cần** phải viết `(&appUser).clearUserName()`. Go cung cấp cơ chế tự động chuyển đổi:

```go
appUser.clearUserName() 
```

Nếu một method yêu cầu một **Pointer Receiver**, nhưng bạn lại gọi nó từ một biến dạng giá trị thông thường (như `appUser`), Go sẽ tự động lấy địa chỉ của biến đó (`&appUser`) để truyền vào method.

---

## ⚖️ 4. So Sánh: Pointer Receiver vs Value Receiver

| Tiêu chí | Value Receiver `(u User)` | Pointer Receiver `(u *User)` |
| :--- | :--- | :--- |
| **Mục đích chính** | Chỉ đọc dữ liệu (Read-only) | Thay đổi dữ liệu (Mutate) hoặc đọc dữ liệu |
| **Cơ chế bộ nhớ** | Tạo bản sao mới (Tốn thêm bộ nhớ nếu Struct lớn) | Truyền địa chỉ ô nhớ (Tiết kiệm bộ nhớ) |
| **Khuyên dùng** | Dành cho các Struct rất nhỏ, bất biến | Hầu hết mọi trường hợp thực tế (đặc biệt khi Struct lớn) |
