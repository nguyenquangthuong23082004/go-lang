# Bài 06: Giá Trị Mặc Định Của Con Trỏ (Zero Value - nil)

> [!IMPORTANT]
> *Trong Go, khi một biến được khai báo mà không gán giá trị khởi tạo, nó sẽ nhận một giá trị mặc định được gọi là **Zero Value**. Đối với con trỏ, giá trị mặc định này là `nil`.*

---

## 📊 1. Bảng So Sánh Giá Trị Mặc Định (Zero Values)

| Kiểu dữ liệu | Giá trị mặc định (Zero Value) |
| :--- | :--- |
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` (Chuỗi rỗng) |
| `bool` | `false` |
| **Pointer (Con trỏ)** | **`nil`** |

---

## 🔍 2. Ý Nghĩa Của `nil` Trong Con Trỏ

`nil` là một giá trị đặc biệt được xây dựng sẵn trong Go. Nó đại diện cho sự **vắng mặt của địa chỉ ô nhớ**. 
Khi một con trỏ có giá trị là `nil`, nghĩa là nó đang không trỏ tới bất kỳ ô nhớ nào trong RAM.

Ví dụ:
```go
var agePointer *int // Chỉ khai báo, chưa khởi tạo
fmt.Println(agePointer) // In ra kết quả: <nil>
```

---

## ⚠️ 3. Nguy Hiểm: Giải Tham Chiếu Con Trỏ `nil` (Nil Pointer Dereference)

Nếu bạn cố tình giải tham chiếu (`*`) một con trỏ đang chứa giá trị `nil`, chương trình của bạn sẽ bị sụp đổ (crash) lập tức với lỗi `panic`:

```go
var agePointer *int
fmt.Println(*agePointer) // 🚫 LỖI CRASH: panic: runtime error: invalid memory address or nil pointer dereference
```

### Cách xử lý an toàn (Best Practice):
Luôn kiểm tra con trỏ khác `nil` trước khi thực hiện giải tham chiếu nếu con trỏ đó có khả năng chưa được khởi tạo:

```go
var agePointer *int

// Kiểm tra an toàn trước khi dùng
if agePointer != nil {
    fmt.Println(*agePointer)
} else {
    fmt.Println("Con trỏ chưa được khởi tạo!")
}
```
