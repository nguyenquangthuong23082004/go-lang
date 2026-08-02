# Bài 26: Lời Giải Bài Tập: Refactoring Profit Calculator Bằng Hàm

> [!IMPORTANT]
> *Trong bài tập thực hành này, chúng ta sẽ tối ưu hóa Bộ tính toán Lợi nhuận (Profit Calculator) bằng cách chuyển đổi logic nhận dữ liệu nhập và logic tính toán tài chính thành các hàm riêng biệt.*

---

## 📋 1. Nhiệm Vụ Cần Đạt Được

1. **Xây dựng hàm `getUserInput`**:
   * Nhận vào lời nhắc hiển thị (ví dụ: `"Nhập doanh thu: "`).
   * Quét và lưu trữ giá trị nhập từ người dùng.
   * Trả về giá trị đã quét (kiểu `float64`).
2. **Xây dựng hàm `calculateFinancials`**:
   * Nhận vào 3 thông số: `revenue`, `expenses`, `taxRate`.
   * Tính toán và trả về đồng thời 3 kết quả: `ebt` (lợi nhuận trước thuế), `profit` (lợi nhuận sau thuế), và `ratio` (tỷ lệ EBT / Profit).
3. **Cập nhật hàm `main`**:
   * Gọi hàm `getUserInput` để lấy dữ liệu cho cả 3 thông số.
   * Gọi hàm `calculateFinancials` để xử lý số liệu.
   * In các kết quả ra màn hình với định dạng chữ số thập phân đẹp mắt bằng `fmt.Printf`.

---

## 💻 2. Mã Nguồn Giải Bài Tập Hoàn Chỉnh

Dưới đây là mã nguồn của file `/home/thuong/Desktop/go-learning/learning/03-profit-calculator/main.go`:

```go
package main

import "fmt"

func main() {
	// 1. Nhận dữ liệu nhập từ người dùng thông qua hàm getUserInput
	revenue := getUserInput("Nhập doanh thu (Revenue): ")
	expenses := getUserInput("Nhập chi phí (Expenses): ")
	taxRate := getUserInput("Nhập thuế suất (%): ")

	// 2. Tính toán các chỉ số tài chính thông qua hàm calculateFinancials
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// 3. Hiển thị kết quả đã được định dạng làm tròn số thập phân
	fmt.Printf("Lợi nhuận trước thuế (EBT): %.1f\n", ebt)
	fmt.Printf("Lợi nhuận sau thuế (Profit): %.1f\n", profit)
	fmt.Printf("Tỷ số EBT / Profit (Ratio): %.3f\n", ratio)
}

// Hàm hỗ trợ nhận đầu vào từ terminal
func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

// Hàm tính toán và trả về đồng thời 3 giá trị float64
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
```

---

## 🔍 3. Giải Thích Chi Tiết Logic Hàm

### A. Cơ chế hoạt động của hàm `getUserInput`
* Hàm nhận đối số `infoText string` để in lời nhắc ra màn hình bằng `fmt.Print`.
* Khai báo biến cục bộ `var userInput float64`. Nhờ cơ chế này, mỗi lần hàm được gọi, một biến trống mới sẽ được tạo trong bộ nhớ.
* Sử dụng con trỏ `&userInput` để nạp dữ liệu nhập từ bàn phím.
* Trả về giá trị của `userInput`. Khi gán ở hàm `main`, chúng ta chỉ cần dùng cú pháp gán nhanh `:=` mà không cần khai báo biến trước.

### B. Cơ chế trả về nhiều giá trị của hàm `calculateFinancials`
* Do cả 3 tham số đầu vào có cùng kiểu `float64`, chúng ta viết rút gọn: `revenue, expenses, taxRate float64`.
* Chữ ký hàm định nghĩa 3 kiểu trả về: `(float64, float64, float64)`.
* Ở hàm `main()`, chúng ta sử dụng phép phân tách (destructuring) để gán đồng thời 3 kết quả vào các biến:
  ```go
  ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
  ```

### C. Bonus: Định dạng đầu ra đẹp mắt bằng `Printf`
* Các giá trị tiền tệ như EBT và Profit được làm tròn tới 1 chữ số thập phân (`%.1f`).
* Tỷ số Ratio được hiển thị chi tiết hơn với 3 chữ số thập phân (`%.3f`).
* Chèn `\n` ở cuối mỗi chuỗi định dạng để ngắt dòng sạch sẽ.
