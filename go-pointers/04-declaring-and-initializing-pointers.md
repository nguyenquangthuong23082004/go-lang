# Bài 04: Cách Khai Báo Và Khởi Tạo Con Trỏ Trong Go

> [!IMPORTANT]
> *Để tạo con trỏ trong Go, chúng ta sử dụng toán tử địa chỉ `&` trước tên biến. Kiểu dữ liệu của con trỏ sẽ được biểu diễn bằng ký tự dấu sao `*` đặt trước kiểu dữ liệu gốc.*

---

## 📍 1. Toán Tử Lấy Địa Chỉ Ô Nhớ `&` (Address-of Operator)

Để lấy địa chỉ ô nhớ của một biến bất kỳ, Go cung cấp toán tử dấu và `&` (ampersand):

```go
age := 32
agePointer := &age // Lấy địa chỉ ô nhớ của biến age
```

Lúc này, `agePointer` không lưu giá trị `32`, mà lưu trữ một địa chỉ bộ nhớ (ví dụ: `0xc000012088`).

---

## 🛠️ 2. Kiểu Dữ Liệu Con Trỏ `*T` (Pointer Type)

Mỗi kiểu dữ liệu trong Go đều có một kiểu con trỏ tương ứng. Chúng ta định nghĩa kiểu con trỏ bằng cách thêm ký tự dấu sao `*` vào trước kiểu dữ liệu gốc:

* `*int`: Con trỏ trỏ tới ô nhớ chứa dữ liệu kiểu `int`.
* `*string`: Con trỏ trỏ tới ô nhớ chứa dữ liệu kiểu `string`.
* `*float64`: Con trỏ trỏ tới ô nhớ chứa dữ liệu kiểu `float64`.
* `*bool`: Con trỏ trỏ tới ô nhớ chứa dữ liệu kiểu `bool`.

---

## 💻 3. Hai Cách Khai Báo Con Trỏ Trong Thực Tế

### Cách 1: Khai báo và khởi tạo nhanh (Khuyên dùng)
Sử dụng toán tử `:=`, Go sẽ tự động suy luận kiểu dữ liệu của biến con trỏ là `*int`:
```go
age := 32
agePointer := &age
```

### Cách 2: Khai báo tường minh kiểu dữ liệu
Định nghĩa biến con trỏ trước bằng từ khóa `var` và kiểu dữ liệu `*T`, sau đó gán giá trị sau:
```go
var agePointer *int // Khai báo một con trỏ kiểu *int
agePointer = &age   // Gán địa chỉ của biến age cho con trỏ
```
*(Cả hai cách đều cho ra kết quả hoàn toàn giống nhau trong bộ nhớ RAM).*
