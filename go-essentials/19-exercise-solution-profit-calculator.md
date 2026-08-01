# Bài 18: Lời Giải Bài Tập: Bộ Tính Toán Lợi Nhuận (Profit Calculator)

> [!NOTE]
> *Dưới đây là lời giải chi tiết và phân tích các quyết định kỹ thuật (như lựa chọn kiểu dữ liệu và tối ưu công thức toán học) khi xây dựng Bộ tính toán lợi nhuận.*

---

## 🛠️ 1. Khởi Tạo Dự Án & Cấu Trúc Mã Nguồn

Để có thể đóng gói thành một file thực thi nhị phân, chúng ta khởi tạo thư mục dự án mới dưới dạng một Go Module:
```bash
go mod init example.com/basics-practice
```

Trong file `profit_calculator.go`, chúng ta thiết lập package `main` và hàm `main()` làm điểm bắt đầu chạy chương trình:
```go
package main

import "fmt"

func main() {
    // Logic viết tại đây...
}
```

---

## 📐 2. Lựa Chọn Kiểu Dữ Liệu & Tránh Lỗi Mismatch Type

Ban đầu, ta có thể nghĩ đến việc dùng kiểu `int` cho Doanh thu và Chi phí, và `float64` cho thuế suất:
```go
var revenue int
var expenses int
var taxRate float64
```

### 🚨 Vấn đề phát sinh:
Khi tính lợi nhuận sau thuế (Profit):
$$\text{Profit} = \text{EBT} \times \left(1 - \frac{\text{Tax Rate}}{100}\right)$$
Vì `EBT` (Revenue - Expenses) có kiểu `int`, còn `taxRate` có kiểu `float64`, Go sẽ không cho phép nhân hai giá trị này với nhau. Chúng ta sẽ phải ép kiểu thủ công: `float64(ebt)`. Ngoài ra, người dùng cũng không thể nhập doanh thu dạng số lẻ (như `1500.50`).

### 💡 Giải pháp tối ưu:
Khai báo cả 3 biến đầu vào ở kiểu số thực **`float64`** ngay từ đầu. Điều này tăng tính linh hoạt khi nhập số liệu lẻ, đồng thời làm công thức tính toán sạch sẽ hơn:
```go
var revenue float64
var expenses float64
var taxRate float64
```

---

## 💻 3. Mã Nguồn Giải Bài Tập Hoàn Chỉnh

Dưới đây là mã nguồn đầy đủ của ứng dụng:

```go
package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// 1. Nhận thông tin đầu vào
	fmt.Print("Nhập doanh thu (Revenue): ")
	fmt.Scan(&revenue)

	fmt.Print("Nhập chi phí (Expenses): ")
	fmt.Scan(&expenses)

	fmt.Print("Nhập thuế suất (%): ")
	fmt.Scan(&taxRate)

	// 2. Tính toán kết quả
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	// 3. Xuất kết quả
	fmt.Println("Lợi nhuận trước thuế (EBT):", ebt)
	fmt.Println("Lợi nhuận sau thuế (Profit):", profit)
	fmt.Println("Tỷ lệ EBT / Profit (Ratio):", ratio)
}
```

---

## 🔬 4. Kiểm Tra Thực Tế (Testing)

Chạy ứng dụng trong thư mục bằng lệnh:
```bash
go run .
```

### Dữ liệu kiểm tra:
* Doanh thu: `10000`
* Chi phí: `7500`
* Thuế suất: `30`

### Đầu ra hiển thị:
```text
Lợi nhuận trước thuế (EBT): 2500
Lợi nhuận sau thuế (Profit): 1750
Tỷ lệ EBT / Profit (Ratio): 1.4285714285714286
```
*Kết quả EBT = 2500 ($10000 - 7500$), Profit = 1750 ($2500 \times 0.7$), và Ratio = 1.42857 ($2500 / 1750$). Chương trình hoạt động hoàn hảo!*
