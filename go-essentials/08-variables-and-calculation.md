# Bài 7: Khai Báo Biến & Thực Hiện Phép Tính Trong Go

> [!NOTE]
> *Trong bài này, chúng ta sẽ bắt đầu viết code cho ứng dụng Tính toán đầu tư. Bạn sẽ học cách khai báo biến để lưu trữ dữ liệu, cách thực hiện các phép toán cơ bản và tìm hiểu một lỗi rất phổ biến về kiểu dữ liệu.*

---

## 📦 1. Biến (Variables) là gì?

Biến là những "hộp chứa dữ liệu" giúp bạn lưu trữ các giá trị trong bộ nhớ của máy tính. Sau khi lưu, bạn có thể gọi tên biến để tái sử dụng giá trị đó ở nhiều nơi trong chương trình.

### ✍️ Khai báo biến bằng từ khóa `var`
Trong Go, cú pháp khai báo biến cơ bản như sau:
```go
var investmentAmount = 1000
```

### 💡 Quy tắc đặt tên biến (Naming Convention):
Cộng đồng Go sử dụng quy tắc viết lạc đà (**camelCase**):
* Chữ cái đầu tiên của từ đầu tiên viết thường.
* Chữ cái đầu tiên của các từ tiếp theo viết hoa.
* *Ví dụ:* `investmentAmount`, `expectedReturnRate`, `futureValue`.
* Nên chọn tên rõ nghĩa, tránh đặt tên quá ngắn hoặc vô nghĩa như `x`, `y`, `a`, `b`.

---

## 🚫 2. Quy Tắc "Không Biến Thừa" Của Go

Khi bạn vừa viết `var investmentAmount = 1000` mà chưa dùng nó ở đâu bên dưới, trình biên dịch sẽ báo lỗi ngay lập tức:
```text
investmentAmount declared and not used
```
* **Lý do:** Go rất nghiêm ngặt trong việc quản lý bộ nhớ. Việc khai báo một biến mà không dùng sẽ làm lãng phí tài nguyên. Go bắt buộc bạn phải xóa biến thừa hoặc sử dụng nó thì chương trình mới biên dịch thành công.

---

## 🧮 3. Thiết Lập Các Biến Đầu Vào & Công Thức Tính

Chúng ta khai báo 3 biến đầu vào chứa các giá trị viết cứng (hardcoded) tạm thời:
```go
var investmentAmount = 1000       // Số tiền đầu tư gốc
var expectedReturnRate = 5.5      // Lợi suất hàng năm (5.5%)
var years = 10                    // Số năm đầu tư (10 năm)
```

### 📈 Công thức tính Giá trị tương lai (Future Value)
Công thức toán học là:
$$\text{Future Value} = \text{Investment Amount} \times \left(1 + \frac{\text{Expected Return Rate}}{100}\right)^{\text{Years}}$$

Trong code Go, chúng ta sử dụng:
* Toán tử nhân: Dấu sao `*`.
* Toán tử chia: Dấu gạch chéo `/`.
* Cặp ngoặc đơn `()` để nhóm thứ tự ưu tiên của phép toán (nhân chia trước, cộng trừ sau).

```go
var futureValue = investmentAmount * (1 + expectedReturnRate / 100) // Cần tính lũy thừa tiếp theo
```

---

## 🚨 4. Lỗi Không Đồng Nhất Kiểu Dữ Liệu (Mismatched Types)

Khi thực hiện phép tính hoặc nhân chia các biến này với nhau, trình biên dịch sẽ đưa ra thông báo lỗi:
```text
invalid operation: mismatched types int and float64
```

### Tại sao lại có lỗi này?
* Biến `investmentAmount` có giá trị `1000` ➔ Go tự động hiểu là kiểu số nguyên (**`int`**).
* Biến `expectedReturnRate` có giá trị `5.5` ➔ Go tự động hiểu là kiểu số thực (**`float64`**).
* Phép tính `(1 + expectedReturnRate / 100)` trả về một số thực (`float64`).
* Trong Go, bạn **không thể tự động thực hiện phép toán giữa các kiểu dữ liệu khác nhau** (ví dụ không thể nhân một số `int` với một số `float64` trực tiếp). Go yêu cầu sự rõ ràng tuyệt đối để tránh mất mát dữ liệu do làm tròn.

Trong bài học tiếp theo, chúng ta sẽ tìm hiểu chi tiết về các kiểu dữ liệu và cách chuyển đổi kiểu (type conversion) để giải quyết lỗi này!
