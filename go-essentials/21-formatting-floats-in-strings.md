# Bài 20: Tinh Chỉnh Định Dạng Số Thực (Formatting Floats) trong Go

> [!NOTE]
> *Hàm `fmt.Printf` cung cấp khả năng kiểm soát hiển thị cực kỳ mạnh mẽ. Bài học này hướng dẫn bạn cách làm tròn chữ số thập phân của số thực, cách in ký tự đặc biệt `%` và cách kiểm tra kiểu dữ liệu của biến.*

---

## 🎨 1. Các Ký Tự Định Dạng Phổ Biến

Ngoài ký tự mặc định `%v`, `fmt.Printf` còn hỗ trợ các bộ định dạng chuyên sâu:

* **`%T` (Type Specifier):** In ra tên kiểu dữ liệu của giá trị truyền vào (Ví dụ: `int`, `float64`, `string`) thay vì in giá trị thực tế của nó. Rất hữu ích khi cần gỡ lỗi (debug).
* **`%%` (Literal Percent):** In ra một ký tự `%` duy nhất. Do dấu `%` đứng một mình luôn được trình biên dịch hiểu là bắt đầu của một placeholder, nên bạn phải dùng hai dấu liên tiếp `%%` nếu muốn hiển thị ký tự phần trăm thực tế lên màn hình.

---

## 🎛️ 2. Làm Tròn Số Thập Phân Với `%f`

Khi hiển thị số thực (`float64`), Go mặc định in ra rất nhiều chữ số thập phân dài ngoằng (ví dụ: `38288.446875`), điều này không cần thiết đối với người dùng cuối. 

Để khắc phục, chúng ta sử dụng bộ định dạng số thực **`%f`** kết hợp với chỉ số làm tròn theo cú pháp: `%.<số_chữ_số>f`

| Cú pháp | Ý nghĩa | Ví dụ giá trị gốc: `38288.446875` |
| :--- | :--- | :--- |
| **`%f`** | Hiển thị đầy đủ số thập phân mặc định | `38288.446875` |
| **`%.0f`** | Làm tròn về số nguyên gần nhất (0 chữ số thập phân) | `38288` |
| **`%.1f`** | Giữ lại 1 chữ số thập phân | `38288.4` |
| **`%.2f`** | Giữ lại 2 chữ số thập phân (tự động làm tròn toán học) | `38288.45` |

---

## 💻 3. Áp Dụng Thực Tế Vào Chương Trình

Chúng ta cập nhật hàm hiển thị trong file `main.go` để làm tròn cả hai kết quả tính toán đầu tư về **2 chữ số thập phân**:

```go
fmt.Printf("Giá trị tương lai danh nghĩa: %.2f\n", giaTriTuongLai)
fmt.Printf("Giá trị thực tế sau lạm phát: %.2f\n", giaTriThucTe)
```

### Kết quả hiển thị trên terminal sau khi cập nhật:
```text
Giá trị tương lai danh nghĩa: 38288.45
Giá trị thực tế sau lạm phát: 32237.84
```

Nhờ có `Printf` kết hợp `%.2f`, kết quả hiển thị đã trở nên gọn gàng, trực quan và dễ hiểu hơn rất nhiều đối với người dùng!
