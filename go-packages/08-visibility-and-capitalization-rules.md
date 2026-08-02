# Bài 08: Quy Tắc Tầm Vực Và Viết Hoa Đầu Hàm (Visibility & Capitalization Rules)

> [!IMPORTANT]
> *Go không sử dụng các từ khóa như `public` hay `private` để cấu hình quyền truy cập. Thay vào đó, Go có một thiết kế độc đáo: **Quyền truy cập của bất kỳ định danh nào (hàm, biến, hằng, struct) được quyết định hoàn toàn bởi chữ cái đầu tiên của nó là viết hoa hay viết thường**.*

---

## ⚖️ 1. So Sánh Với Các Ngôn Ngữ Khác

| Ngôn ngữ | Quyền Public | Quyền Private |
| :--- | :--- | :--- |
| **Java / C# / C++** | Sử dụng từ khóa `public` | Sử dụng từ khóa `private` |
| **TypeScript** | Sử dụng từ khóa `export` | Không dùng export hoặc dùng `private` |
| **Go (Golang)** | **Viết hoa** chữ cái đầu | **Viết thường** chữ cái đầu |

---

## 🛠️ 2. Quy Tắc Viết Hoa Trong Go (Exported vs Unexported)

Quy tắc này áp dụng cho mọi loại định danh trong Go, bao gồm: Hàm (Functions), Biến (Variables), Structs, Trường của Struct (Struct Fields), Interfaces và Hằng số (Constants).

### A. Định danh Xuất khẩu (Exported - tương đương Public)
* **Quy tắc:** Bắt đầu bằng một chữ cái viết hoa (A-Z).
* **Khả năng:** Có thể được truy cập tự do từ bất kỳ package nào ngoài gói hiện tại.
* **Ví dụ trong gói `fileops`**:
  ```go
  func WriteFloatToFile(...) // Bắt đầu bằng 'W' -> Gói main gọi được
  ```

### B. Định danh Không xuất khẩu (Unexported - tương đương Private)
* **Quy tắc:** Bắt đầu bằng một chữ cái viết thường (a-z) hoặc dấu gạch dưới (`_`).
* **Khả năng:** Chỉ có thể truy cập được bởi các file nằm trong **cùng một package** đó. Gói bên ngoài hoàn toàn không nhìn thấy.
* **Ví dụ trong gói `fileops`**:
  ```go
  func writeFloatToFile(...) // Bắt đầu bằng 'w' -> Gói main gọi sẽ báo lỗi biên dịch
  ```

---

## 💡 3. Ví Dụ Về Trường Của Struct (Struct Fields)

Quy tắc viết hoa đầu chữ này đặc biệt quan trọng khi làm việc với `Struct` và chuyển đổi sang JSON (JSON Serialization):

```go
type User struct {
    Name string // Viết hoa -> Các gói khác truy cập được, có thể chuyển sang JSON
    age  int    // Viết thường -> Chỉ dùng nội bộ, thư viện JSON không thể đọc được
}
```

---

## 🧠 4. Tại Sao Người Thiết Kế Go Lại Làm Như Vậy?

1. **Đọc code nhanh hơn**: Khi đọc mã nguồn Go, bạn chỉ cần nhìn vào tên hàm/biến là biết ngay nó là Public hay Private mà không cần phải cuộn lên đầu hàm hay xem định nghĩa để tìm các từ khóa `public`/`private`.
2. **Tối giản cú pháp**: Giảm thiểu số lượng từ khóa giúp trình biên dịch của Go (Go Compiler) chạy cực kỳ nhanh và giữ cho cú pháp Go luôn tinh gọn, dễ tiếp cận.
