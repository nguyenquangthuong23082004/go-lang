# Bài 10: Giới Thiệu Về Maps (Bản Đồ Cặp Key-Value)

> [!NOTE]
> *Tìm hiểu cấu trúc dữ liệu **Map** trong Go: Lý do cần dùng Map thay vì Array/Slice khi cần lưu dữ liệu dạng cặp Khóa - Giá trị (Key-Value), và cú pháp khai báo, khởi tạo Map.*

---

## 💡 1. Tại Sao Cần Dùng Map?

Khi lưu danh sách dữ liệu (ví dụ danh sách URL trang web) bằng Array hoặc Slice:
```go
websites := []string{"https://google.com", "https://aws.com"}
```
- **Hạn chế 1 (Khó nhận biết):** Nhìn vào `"https://aws.com"`, người xem chưa chắc biết đây là trang web của công ty nào nếu không được giải thích trước.
- **Hạn chế 2 (Khó truy xuất):** Để lấy URL của Google, bạn phải nhớ vị trí chỉ mục số (`0`). Nếu danh sách có hàng trăm phần tử, việc nhớ chỉ mục số để lấy đúng dữ liệu là rất khó khăn.

👉 **Giải pháp:** Sử dụng **Map** để gán nhãn (Key) cho từng giá trị (Value).

---

## 🗺️ 2. Map Trong Go Là Gì?

**Map** là cấu trúc dữ liệu lưu trữ dưới dạng **Cặp Key - Value** (Khóa - Giá trị). Giống như từ điển: bạn tìm từ khóa (Key) để xem nghĩa của từ đó (Value).

---

## 📝 3. Cú Pháp Khai Báo Và Khởi Tạo Map

```go
map[Kiểu_Dữ_Liệu_Key]Kiểu_Dữ_Liệu_Value{ ... }
```

### Ví dụ cụ thể:

```go
package main

import "fmt"

func main() {
    // Khai báo Map có Key là string và Value cũng là string
    websites := map[string]string{
        "Google":              "https://google.com",
        "Amazon Web Services": "https://aws.com",
    }

    // In toàn bộ Map ra màn hình
    fmt.Println(websites)
    // Output: map[Amazon Web Services:https://aws.com Google:https://google.com]
}
```

### Phân tích chi tiết:
1. `map[string]`: Định nghĩa kiểu dữ liệu cho **Key** là `string` (tên công ty: `"Google"`).
2. `string`: Định nghĩa kiểu dữ liệu cho **Value** tương ứng là `string` (đường dẫn: `"https://google.com"`).
3. Đặt trong cặp ngoặc `{}` là các cặp `Key: Value` phân tách nhau bởi dấu phẩy `,`.

> [!TIP]
> Key và Value không bắt đầu bắt buộc phải cùng kiểu dữ liệu. Bạn có thể dùng `map[string]int`, `map[int]string`, hoặc `map[string]MyStruct` tùy thuộc vào nhu cầu.
