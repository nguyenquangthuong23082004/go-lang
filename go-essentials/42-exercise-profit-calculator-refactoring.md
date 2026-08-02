# Bài 41: Bài Tập Thực Hành: Xử Lý Lỗi & Lưu File Cho Profit Calculator

> [!NOTE]
> *Để thực hành các kiến thức đã học về cấu trúc điều khiển, xử lý lỗi và lưu file, chúng ta tiến hành nâng cấp dự án Profit Calculator (Bộ tính toán lợi nhuận).*

---

## 🎯 1. Yêu Cầu Đề Bài

1. **Ràng buộc đầu vào:**
   * Thay đổi hàm nhận dữ liệu `getUserInput` để trả về thêm lỗi (`error`).
   * Kiểm tra nếu giá trị doanh thu, chi phí hay thuế suất nhập vào nhỏ hơn hoặc bằng 0 thì sinh lỗi bằng `errors.New`.
   * Ở hàm `main()`, đón nhận lỗi này. Nếu phát hiện lỗi, in thông báo và dừng ứng dụng ngay lập tức (không chạy tính toán tiếp).
2. **Lưu trữ kết quả ra file:**
   * Sau khi tính toán ra các giá trị `ebt` (Lợi nhuận trước thuế), `profit` (Lợi nhuận sau thuế) và `ratio` (Tỷ số lợi nhuận), hãy lưu chúng vào file `results.txt` trong cùng thư mục dự án.

---

## 🛠️ 2. Hướng Dẫn Giải Quyết Chi Tiết

### A. Ràng buộc dữ liệu & Trả về lỗi
Chúng ta cập nhật chữ ký hàm `getUserInput` để trả về `(float64, error)`:
```go
func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid value. Value must be greater than 0.")
	}

	return userInput, nil
}
```

Ở hàm `main()`, kiểm tra biến lỗi `err` sau mỗi lần gọi `getUserInput`:
```go
revenue, err := getUserInput("Nhập doanh thu (Revenue): ")
if err != nil {
	fmt.Println(err)
	return // Dừng ứng dụng sớm
}
```

### B. Định dạng kết quả và ghi file
Chúng ta tạo thêm hàm `storeResults` nhận vào ba số thực và lưu trữ chúng bằng `os.WriteFile`:
```go
func storeResults(ebt, profit, ratio float64) {
	// Định dạng chuỗi văn bản gồm nhiều dòng
	resultsText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	
	// Ghi mảng byte ra kết quả kết hợp phân quyền 0644
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}
```

---

## 💻 3. Mã Nguồn Hoàn Chỉnh

Mã nguồn file `/home/thuong/Desktop/go-learning/learning/03-profit-calculator/main.go` sau khi nâng cấp:

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Nhập doanh thu (Revenue): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Nhập chi phí (Expenses): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Nhập thuế suất (%): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Lợi nhuận trước thuế (EBT): %.1f\n", ebt)
	fmt.Printf("Lợi nhuận sau thuế (Profit): %.1f\n", profit)
	fmt.Printf("Tỷ số EBT / Profit (Ratio): %.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid value. Value must be greater than 0.")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	resultsText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}
```
*Sau khi chạy thành công chương trình, kết quả tính toán sẽ được lưu trữ tự động trong file `results.txt` tại thư mục `/03-profit-calculator/`.*
