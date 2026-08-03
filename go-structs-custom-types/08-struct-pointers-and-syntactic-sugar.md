# Bài 08: Truyền Con Trỏ Struct & Cú Pháp Rút Gọn Trong Go

> [!NOTE]
> *Chúng ta sẽ học cách tối ưu bộ nhớ bằng cách truyền con trỏ Struct vào hàm, đồng thời tìm hiểu cú pháp rút gọn đặc biệt mà Go hỗ trợ để làm việc với các trường của con trỏ Struct.*

---

## ⚙️ 1. Tại Sao Nên Truyền Con Trỏ Struct? (Pass by Pointer)

Theo cơ chế mặc định của Go, khi bạn truyền một Struct vào hàm, Go sẽ thực hiện **sao chép giá trị (Pass-by-value)**:

* Go tạo ra một bản sao hoàn toàn mới của Struct đó trong một ô nhớ khác.
* Đối với các Struct nhỏ, điều này không gây vấn đề gì lớn.
* Tuy nhiên, đối với các Struct lớn chứa nhiều dữ liệu phức tạp (như danh sách, bản đồ, hàng chục thuộc tính), việc sao chép liên tục qua các hàm sẽ gây hao phí tài nguyên bộ nhớ và làm chậm chương trình.

**Giải pháp:** Truyền địa chỉ của Struct gốc (con trỏ) bằng toán tử lấy địa chỉ `&`:

```go
// Truyền địa chỉ ô nhớ của appUser thay vì bản sao
outputUserDetails(&appUser)
```

---

## 👈 2. Định Nghĩa Hàm Nhận Con Trỏ Struct

Trong chữ ký hàm, chúng ta sử dụng dấu hoa thị `*` trước tên kiểu Struct để khai báo rằng hàm này nhận vào một con trỏ:

```go
func outputUserDetails(u *User) {
	// u bây giờ chứa địa chỉ ô nhớ của thực thể User
}
```

---

## 🌟 3. Cú Pháp Rút Gọn Đặc Biệt (Syntactic Sugar)

Bình thường trong Go, để truy cập vào giá trị nằm bên trong một con trỏ, ta bắt buộc phải thực hiện giải tham chiếu (**Dereference**) bằng toán tử `*`.

Vì dấu chấm truy cập trường (`.`) có độ ưu tiên cao hơn toán tử dấu hoa thị (`*`), nếu viết theo đúng lý thuyết kỹ thuật thì ta phải bọc dấu ngoặc đơn xung quanh biến giải tham chiếu:

```go
// Cách viết chuẩn lý thuyết kỹ thuật
fmt.Println("First Name:", (*u).firstName)
```

Tuy nhiên, việc viết ngoặc đơn liên tục làm code trở nên rối mắt. Do đó, Go đã bổ sung một cơ chế hỗ trợ đặc biệt (**Syntactic Sugar**):

```go
// Cách viết rút gọn được Go cho phép
fmt.Println("First Name:", u.firstName)
```

Khi bạn sử dụng `u.firstName` với `u` là một con trỏ Struct, **Go sẽ tự động giải tham chiếu con trỏ đó dưới nền** để truy cập thuộc tính `firstName`. Đây là một ngoại lệ cực kỳ hữu ích và hầu như toàn bộ lập trình viên Go đều sử dụng cú pháp rút gọn này khi làm việc với con trỏ Struct.
