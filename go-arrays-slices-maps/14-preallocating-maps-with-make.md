# Bài 14: Khởi Tạo Map Tối Ưu Bộ Nhớ Bằng Hàm `make()`

> [!NOTE]
> *Giống như Slice, hàm **`make()`** cũng có thể dùng để khởi tạo **Map**. Tìm hiểu cú pháp và sự khác biệt giữa `make()` dùng cho Map và Slice.*

---

## 🗺️ 1. Tại Sao Nên Dùng `make()` Cho Map?

Khi tạo một Map rỗng bằng cú pháp literal `map[string]float64{}`:
- Ban đầu Go chưa cấp phát dung lượng lưu trữ lớn.
- Mỗi lần thêm một Key mới vượt quá dung lượng hiện tại, Go phải phân bổ (reallocate) lại bộ nhớ và tổ chức lại bảng băm (hash table) ngầm bên dưới.
- Nếu biết trước số lượng phần tử dự kiến sẽ thêm vào Map, sử dụng `make()` giúp Go **đặt trước sức chứa trong bộ nhớ**, giúp chương trình chạy mượt mà và tối ưu hơn.

---

## 🛠️ 2. Cú Pháp Khai Báo `make()` Với Map

Khác với Slice (truyền 2 tham số `len` và `cap`), **Map chỉ truyền tối đa 1 tham số sức chứa (capacity)**:

```go
courseRatings := make(map[Kiểu_Key]Kiểu_Value, capacity_dự_kiến)
```

### Ví dụ thực tế:
```go
package main

import "fmt"

func main() {
    // Khởi tạo Map với sức chứa dự kiến trước là 3 phần tử
    courseRatings := make(map[string]float64, 3)

    // Thêm các phần tử mà KHÔNG làm Go phải cấp phát lại bộ nhớ
    courseRatings["go"] = 4.7
    courseRatings["react"] = 4.8
    courseRatings["angular"] = 4.5

    fmt.Println(courseRatings)
    // Output: map[angular:4.5 go:4.7 react:4.8]
}
```

---

## ⚖️ 3. So Sánh `make()` Giữa Slice Và Map

| Đặc điểm | `make()` với Slice ✂️ | `make()` với Map 🗺️ |
| :--- | :--- | :--- |
| **Số tham số phụ** | Có thể truyền **2 tham số** (`len`, `cap`). | Chỉ truyền **1 tham số** (`capacity`). |
| **Khái niệm ô rỗng** | Có (nếu `len > 0`, các ô ban đầu chứa zero value). | **Không có ô rỗng**. Map không có khái niệm chỉ mục rỗng vì Key phải được gán đích danh. |
| **Cú pháp mẫu** | `make([]string, 0, 5)` | `make(map[string]float64, 5)` |

---

## 📝 Tóm Tắt

- `make(map[K]V, size)` giúp cấp phát sẵn bộ nhớ cho Map.
- Giúp ứng dụng hoạt động tối ưu hơn khi bạn xử lý các tập dữ liệu Map có số lượng phần tử lớn hoặc biết trước số lượng.
