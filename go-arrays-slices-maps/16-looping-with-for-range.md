# Bài 16: Vòng Lặp `for range` Duyệt Qua Array, Slice Và Map

> [!NOTE]
> *Tìm hiểu cú pháp từ khóa **`range`** kết hợp với vòng lặp **`for`** để lặp qua từng phần tử của Array, Slice hoặc Map trong Go.*

---

## 🔄 1. Cú Pháp Vòng Lặp `for range` Cho Array / Slice

Khi duyệt qua một Array hoặc Slice, `range` sẽ trả về 2 giá trị ở mỗi vòng lặp:
1. **`index` (Chỉ mục):** Vị trí của phần tử trong mảng (`0, 1, 2...`).
2. **`value` (Giá trị):** Giá trị của phần tử tại vị trí đó.

```go
package main

import "fmt"

func main() {
    userNames := []string{"Max", "Manuel", "Julie"}

    for index, value := range userNames {
        fmt.Printf("Index: %d | Value: %s\n", index, value)
    }
}
```

---

## 🗺️ 2. Cú Pháp Vòng Lặp `for range` Cho Map

Khi duyệt qua một Map, `range` trả về:
1. **`key` (Khóa):** Tên khóa của phần tử.
2. **`value` (Giá trị):** Giá trị lưu trữ tương ứng với khóa đó.

```go
courseRatings := map[string]float64{
    "go":      4.7,
    "react":   4.8,
    "angular": 4.5,
}

for key, value := range courseRatings {
    fmt.Printf("Key: %s | Value: %.1f\n", key, value)
}
```

> [!IMPORTANT]
> **Lưu ý về thứ tự duyệt Map:** Thức tự duyệt qua các phần tử của Map trong Go là **ngẫu nhiên (un-ordered)**. Không nên dựa vào thứ tự in ra của Map.

---

## 🙈 3. Bỏ Qua `index` Hoặc `value` Với Khai Báo Biến Blank (`_`)

Nếu bạn chỉ quan tâm đến giá trị (`value`) mà không cần `index` / `key` (hoặc ngược lại), bạn **phải** dùng dấu gạch dưới `_` để bỏ qua, vì Go không cho phép biến được khai báo mà không sử dụng:

### Bỏ qua `index`:
```go
for _, value := range userNames {
    fmt.Println("User:", value)
}
```

### Bỏ qua `value` (Chỉ lấy `key` / `index`):
```go
for key := range courseRatings {
    fmt.Println("Course:", key)
}
```

---

## 📝 Tóm Tắt

- `for index, value := range slice`: Duyệt Array/Slice.
- `for key, value := range myMap`: Duyệt Map.
- Dùng `_` để bỏ qua các biến không cần dùng.
