# Bài 19: Định Dạng Dữ Liệu Đầu Ra Bằng fmt.Printf & Cách Viết Chú Thích (Comments)

> [!NOTE]
> *Để kết quả đầu ra trực quan và chuyên nghiệp hơn, Go cung cấp hàm định dạng chuỗi nâng cao `fmt.Printf`. Đồng thời, bạn cũng sẽ học cách sử dụng chú thích (`//`) để ghi chú hoặc tạm thời vô hiệu hóa code.*

---

## 📝 1. Cách Viết Chú Thích (Comments) Trong Go

Chú thích là những đoạn văn bản được viết nhằm giải thích mã nguồn hoặc tạm thời vô hiệu hóa một đoạn code khi chạy thử. Trình biên dịch Go sẽ bỏ qua hoàn toàn các dòng chú thích này.

* **Cú pháp:** Sử dụng hai dấu gạch chéo xuôi `//` ở đầu dòng.
  ```go
  // Đây là dòng chú thích giải thích công thức tính toán
  // ebt := revenue - expenses
  ```

---

## 📺 2. Ghép Chuỗi Cơ Bản Với `fmt.Println()`

Cách đơn giản nhất để giải thích con số đầu ra là truyền thêm một nhãn văn bản (string label) phân tách bằng dấu phẩy `,` vào hàm `fmt.Println`:
```go
fmt.Println("Giá trị tương lai danh nghĩa:", futureValue)
```
* **Đặc điểm:** `Println` tự động thêm một khoảng trắng phân cách giữa các đối số và tự động xuống dòng ở cuối.

---

## 🛠️ 3. Định Dạng Chuỗi Chuyên Nghiệp Bằng `fmt.Printf()`

Hàm `fmt.Printf()` (viết tắt của Print Formatted) giúp bạn cấu trúc nội dung hiển thị thông qua các **ký tự giữ chỗ (placeholders)** nằm bên trong chuỗi văn bản.

### 📍 Ký tự giữ chỗ `%v` (Value Specifier)
Ký tự `%v` đại diện cho giá trị mặc định của một biến (hoặc hằng số) bất kể kiểu dữ liệu là gì.

```go
fmt.Printf("Giá trị tương lai danh nghĩa: %v", futureValue)
```

### 🚨 Các lưu ý quan trọng khi dùng `Printf`:

1. **Xuống dòng thủ công bằng `\n`:** Hàm `Printf` không tự động xuống dòng sau khi in. Bạn phải chèn ký tự đặc biệt `\n` (newline) tại vị trí muốn ngắt dòng:
   ```go
   fmt.Printf("Giá trị tương lai danh nghĩa: %v\n", futureValue)
   ```

2. **Sử dụng nhiều Placeholders:** Bạn có thể kết hợp nhiều `%v` trong một dòng lệnh duy nhất, và truyền các biến vào phía sau theo đúng thứ tự:
   ```go
   fmt.Printf("Danh nghĩa: %v\nThực tế (trừ lạm phát): %v\n", futureValue, futureRealValue)
   ```
   *Số lượng biến truyền sau chuỗi định dạng phải bằng số lượng ký tự giữ chỗ `%v` trong chuỗi.*
