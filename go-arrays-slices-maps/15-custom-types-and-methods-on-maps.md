# Bài 15: Tạo Type Alias (Custom Type) Cho Map & Thêm Receiver Method

> [!NOTE]
> *Tìm hiểu cách sử dụng **Type Alias (Custom Type)** để rút gọn các kiểu Map/Slice dài dòng và gán thêm các **Receiver Methods** tùy chỉnh cho kiểu dữ liệu này.*

---

## 🎨 1. Tại Sao Nên Tạo Custom Type Cho Map/Slice?

Khi làm việc với các kiểu dữ liệu phức tạp hoặc dài dòng như `map[string]float64`, `map[string]string`, việc viết đi viết lại kiểu dữ liệu này ở nhiều nơi:
1. Làm code trở nên dài dòng và kém gọn gàng.
2. Không thể gắn thêm các phương thức (Methods) riêng cho các kiểu dữ liệu built-in mặc định.

👉 **Giải pháp:** Định nghĩa một **Custom Type (Type Alias)**.

---

## 🛠️ 2. Định Nghĩa Custom Type & Gán Receiver Method

```go
package main

import "fmt"

// 1. Khai báo Custom Type (tên ngắn gọn thế chỗ cho map[string]float64)
type floatMap map[string]float64

// 2. Thêm phương thức tùy chỉnh (Receiver Method) cho floatMap
func (m floatMap) output() {
    fmt.Println("Bảng điểm khóa học:", m)
}

func main() {
    // 3. Sử dụng floatMap với hàm make() cực kỳ gọn gàng
    courseRatings := make(floatMap, 3)

    courseRatings["go"] = 4.7
    courseRatings["react"] = 4.8
    courseRatings["angular"] = 4.5

    // 4. Gọi phương thức trực tiếp từ biến courseRatings
    courseRatings.output()
}
```

---

## 💎 3. Lợi Ích Của Cụm Kỹ Thuật Này

1. **Rút ngắn cú pháp (Concise Code):** Thay vì gõ `map[string]float64`, bạn chỉ cần ghi `floatMap`.
2. **Khả năng mở rộng (Extensibility):** Bạn có thể tự viết thêm các hàm tiện ích riêng gắn với kiểu dữ liệu đó (`.output()`, `.clear()`, `.printKeys()`,...).
3. **Mã nguồn sạch sẽ (Clean Code):** Giúp code dễ đọc, dễ bảo trì và thể hiện rõ ý đồ nghiệp vụ (Domain Meaning).
