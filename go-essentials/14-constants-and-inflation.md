# Bài 13: Sử Dụng Hằng Số (Constants) & Bài Toán Lạm Phát Trong Go

> [!NOTE]
> *Trong bài này, chúng ta sẽ nâng cấp Bộ tính toán đầu tư bằng cách tính thêm tỷ lệ lạm phát hàng năm. Bạn sẽ được làm quen với khái niệm hằng số (`const`) và cách bảo vệ dữ liệu không bị thay đổi ngoài ý muốn.*

---

## 📉 1. Tác Động Của Lạm Phát (Inflation) Đến Giá Trị Tiền Tệ

Trong tài chính, số tiền bạn nhận được trong tương lai (Nominal Future Value - Giá trị danh nghĩa) sẽ bị giảm sức mua do lạm phát. Để tính được giá trị thực tế sau khi trừ đi lạm phát (Real Future Value - Giá trị thực tế), chúng ta sử dụng công thức:

$$\text{Future Real Value} = \frac{\text{Future Value}}{\left(1 + \frac{\text{Inflation Rate}}{100}\right)^{\text{Years}}}$$

Nếu tỷ lệ lạm phát hàng năm cao hơn tỷ suất lợi nhuận đầu tư, giá trị thực tế của tiền bạn thu về sẽ nhỏ hơn số tiền gốc ban đầu.

---

## 🔒 2. Khái Niệm Hằng Số (`const`) Trong Go

Trong khi các thông số đầu tư (tiền gốc, số năm, lãi suất) sẽ do người dùng nhập vào, chúng ta muốn cố định tỷ lệ lạm phát là một thông số hệ thống và không được phép thay đổi. Để làm được điều này, Go cung cấp từ khóa **`const`**:

```go
const inflationRate = 2.5 // Tỷ lệ lạm phát 2.5% cố định
```

### 🛡️ Đặc điểm của Hằng số (`const`):
* **Không thể sửa đổi:** Giá trị của hằng số được quyết định tại thời điểm biên dịch và **không thể gán lại** trong suốt quá trình chương trình chạy.
* **Lỗi biên dịch:** Nếu bạn cố tình gán lại giá trị cho hằng số (ví dụ: `inflationRate = 3.0`), trình biên dịch Go sẽ báo lỗi ngay lập tức:
  ```text
  cannot assign to inflationRate (neither addressable nor map index expression)
  ```
* **Lưu ý kiểu số:** Nếu giá trị hằng số là số nguyên (ví dụ: `2`), hãy ép kiểu tường minh hoặc viết thêm phần thập phân để tránh lỗi mismatch type khi tính toán với số thực: `const inflationRate float64 = 2`.

---

## 💻 3. Tích Hợp Lạm Phát Vào Mã Nguồn Go

Hãy cập nhật mã nguồn để thực hiện phép tính điều chỉnh lạm phát:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, expectedReturnRate, years := 5000.0, 5.5, 3.0

	// 1. Tính giá trị tương lai danh nghĩa
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// 2. Tính giá trị tương lai thực tế (trừ lạm phát)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// 3. Hiển thị kết quả
	fmt.Println("Giá trị tương lai danh nghĩa:", futureValue)
	fmt.Println("Giá trị thực tế sau lạm phát:", futureRealValue)
}
```

---

## 🔬 4. Chạy Thử Nghiệm & So Sánh Kết Quả

Chạy chương trình trong terminal bằng lệnh:
```bash
go run .
```

### Kết quả hiển thị:
```text
Giá trị tương lai danh nghĩa: 5871.206874999999
Giá trị thực tế sau lạm phát: 5451.977462002361
```

* **Trường hợp lạm phát cao:** Nếu bạn thử đổi hằng số `inflationRate` thành `6.5` (lớn hơn tỷ suất lợi nhuận `5.5`), bạn sẽ thấy giá trị thực tế sau lạm phát nhận về là khoảng `4860`, nhỏ hơn số tiền gốc `5000` ban đầu. Điều này hoàn toàn chính xác với nguyên lý tài chính!
