# Bài 17: Custom Types (Kiểu Tùy Chỉnh) & Bí Danh Kiểu (Type Aliases)

> [!NOTE]
> *Chúng ta sẽ tìm hiểu cách sử dụng từ khóa `type` để tạo ra các kiểu dữ liệu tùy chỉnh (Custom Types) từ các kiểu dữ liệu có sẵn của Go, từ đó cho phép đính kèm các phương thức (methods) riêng lên chúng.*

---

## 🏷️ 1. Khái Niệm Custom Types (Kiểu Tùy Chỉnh)

Từ khóa `type` trong Go không chỉ được sử dụng để tạo Struct. Chúng ta còn có thể sử dụng nó để định nghĩa ra một kiểu dữ liệu tùy chỉnh mới dựa trên các kiểu dữ liệu dựng sẵn của Go (như `string`, `int`, `float64`):

```go
type customString string // customString là kiểu tùy chỉnh dựa trên string
```

---

## ⚡ 2. Tại Sao Phải Tạo Custom Type Cho Kiểu Dữ Liệu Có Sẵn?

Trong Go, có một quy tắc vàng về tầm vực: **Bạn không thể khai báo một phương thức (Method) cho một kiểu dữ liệu không thuộc nội bộ package hiện tại (non-local type).**

Ví dụ, đoạn code sau sẽ bị báo lỗi biên dịch ngay lập tức:
```go
// LỖI BIÊN DỊCH: cannot define new methods on non-local type string
func (s string) log() {
	fmt.Println(s)
}
```

**Giải pháp:** Bằng cách tạo ra một Custom Type dựa trên kiểu dữ liệu gốc, Go sẽ tạo ra một phiên bản kiểu dữ liệu mới thuộc về package của bạn. Nhờ vậy, bạn có thể tự do đính kèm thêm phương thức vào kiểu dữ liệu tùy chỉnh đó:

```go
type customString string

// Hoàn toàn hợp lệ
func (text customString) log() {
	fmt.Println(text)
}
```

---

## 📍 3. Cách Sử Dụng Trong Thực Tế

Khi khởi tạo một biến sử dụng Custom Type, bạn cần phải khai báo rõ ràng kiểu dữ liệu của biến. Nếu bạn sử dụng cơ chế tự động suy luận kiểu dữ liệu (`:=`), Go sẽ gán kiểu `string` mặc định của hệ thống:

```go
// Khai báo rõ ràng kiểu dữ liệu tùy chỉnh
var name customString = "Max"

// Gọi phương thức đã định nghĩa trên customString
name.log() // In ra màn hình: Max
```

Mẫu thiết kế này sẽ trở nên cực kỳ mạnh mẽ khi chúng ta học về các cấu trúc dữ liệu phức tạp hơn như Slices, Maps hoặc Function (Hàm dưới dạng tham số) ở các phần tiếp theo của khóa học.
