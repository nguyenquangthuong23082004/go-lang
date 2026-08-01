# Bài 12: Các Cách Khai Báo Biến Khác Nhau (Alternative Declaration Styles) trong Go

> [!NOTE]
> *Để tối giản hóa mã nguồn và tăng hiệu suất viết code, Go cung cấp một số cú pháp viết tắt độc đáo như toán tử `:=` hoặc khai báo nhiều biến trên cùng một dòng. Hãy cùng tìm hiểu cách áp dụng chúng.*

---

## ⚡ 1. Toán Tử Khai Báo Nhanh `:=` (Short Variable Declaration)

Đây là cú pháp khai báo và khởi tạo biến phổ biến nhất trong Go khi lập trình thực tế. Cú pháp này loại bỏ từ khóa `var` và thay dấu bằng `=` bằng dấu hai chấm bằng `:=`:

```go
futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
```

### 🚨 Các quy tắc bắt buộc của `:=`
1. **Chỉ dùng trong hàm:** Toán tử `:=` chỉ được phép sử dụng **bên trong thân hàm** (ví dụ trong `func main()`). Bạn không thể dùng nó để khai báo biến toàn cục (package-level variables) ở ngoài hàm.
2. **Không khai báo kiểu:** Bạn không được ghi kèm kiểu dữ liệu (như `float64`) khi sử dụng `:=`.
3. **Mẹo khai báo float64 nhanh:** Nếu dùng `:=` để gán một số nguyên nhưng muốn Go hiểu là số thực, hãy thêm đuôi `.0` đằng sau số đó:
   ```go
   investmentAmount := 1000.0 // Go suy luận kiểu float64 thay vì int
   ```

---

## 📂 2. Khai Báo Nhiều Biến Trên Cùng Một Dòng

Go hỗ trợ gom nhóm khai báo nhiều biến để tiết kiệm diện tích viết code:

### 💡 Trường hợp 1: Sử dụng suy luận kiểu tự động (dùng `:=`)
Bạn có thể khai báo các biến có kiểu dữ liệu khác nhau trên cùng một dòng:
```go
name, age, height := "Thuong", 22, 1.75
```
* Go tự động gán `"Thuong"` cho `name` (kiểu `string`), `22` cho `age` (kiểu `int`), và `1.75` cho `height` (kiểu `float64`).

### ✍️ Trường hợp 2: Khai báo kiểu tường minh (dùng `var`)
Nếu bạn chỉ định rõ kiểu dữ liệu, tất cả các biến trên dòng đó **bắt buộc phải có cùng kiểu**:
```go
var investmentAmount, years float64 = 1000, 10
```
* **Lưu ý:** Kiểu dữ liệu `float64` được ghi ở cuối cùng và có hiệu lực cho tất cả các biến trước nó. Bạn không được khai báo kiểu riêng lẻ kiểu như `var a int, b string = 1, "test"`.

---

## ⚖️ 3. So Sánh & Lời Khuyên Về Readability (Tính Dễ Đọc)

Bạn có thể gộp tất cả các biến của chương trình tính toán đầu tư vào một dòng duy nhất:
```go
investmentAmount, expectedReturnRate, years := 5000.0, 5.5, 3.0
```

### Lời khuyên thực tế:
* **Không nên lạm dụng:** Gộp quá nhiều biến vào một dòng làm code rất khó đọc, người đọc phải tự đếm thứ tự biến và giá trị để ghép cặp.
* **Khuyên dùng:** Nên tách các biến ra nhiều dòng riêng biệt để giữ code luôn tường minh, dễ chỉnh sửa. Việc chọn cú pháp dài (`var`) hay cú pháp ngắn (`:=`) phụ thuộc vào việc bạn có cần ép kiểu tường minh ngay từ đầu hay không.
