# Bài 42: Phân Tích Lời Giải Bài Tập Profit Calculator

> [!NOTE]
> *Qua bài tập này, chúng ta học được cách thiết kế hàm có kiểm tra lỗi đầu vào và hai hướng tiếp cận để xử lý lỗi trong hàm `main()`, cùng cách kết hợp định dạng chuỗi nhiều dòng bằng `fmt.Sprintf` trước khi ghi file.*

---

## 🚦 1. Hai Cách Xử Lý Lỗi Đầu Vào Trong Hàm `main`

Khi gọi hàm `getUserInput` nhiều lần, chúng ta có hai cách để bắt lỗi:

### Cách 1: Kiểm tra lỗi ngay lập tức sau mỗi lần gọi (Cách Khuyên Dùng)
Mỗi lần nhận dữ liệu, ta kiểm tra lỗi ngay lập tức. Nếu có lỗi, in thông báo và dừng chương trình.
```go
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
```
* **Ưu điểm:** Đây là phong cách viết code chuẩn Go (idiomatic Go). Giúp bạn dễ dàng tùy biến thông điệp lỗi riêng cho từng tham số (Ví dụ: "Doanh thu không hợp lệ", "Chi phí không hợp lệ").

### Cách 2: Khai báo nhiều biến lỗi và kiểm tra gộp (Cách Rút Gọn)
Khai báo các biến lỗi riêng biệt (`err1`, `err2`, `err3`) rồi gộp chung điều kiện kiểm tra bằng toán tử OR (`||`):
```go
revenue, err1 := getUserInput("Nhập doanh thu: ")
expenses, err2 := getUserInput("Nhập chi phí: ")
taxRate, err3 := getUserInput("Nhập thuế suất: ")

if err1 != nil || err2 != nil || err3 != nil {
	fmt.Println("Có lỗi nhập liệu xảy ra!")
	return
}
```
* **Hạn chế:** Cách này khiến chúng ta không thể chỉ ra chính xác tham số nào bị lỗi để thông báo chi tiết cho người dùng. Hơn nữa, chương trình sẽ vẫn tiếp tục hỏi người dùng nhập tham số 2 và 3 dù tham số 1 đã nhập sai từ trước. Do đó, **Cách 1 vẫn luôn là lựa chọn tốt nhất**.

---

## 📝 2. Định Dạng Chuỗi Nhiều Dòng Với `fmt.Sprintf`

Để ghi nhiều kết quả tính toán có cấu trúc rõ ràng vào tệp tin `results.txt`, thay vì ghi đè nhiều lần, chúng ta định dạng tất cả thành một chuỗi duy nhất:

```go
func storeResults(ebt, profit, ratio float64) {
	// Sử dụng ký tự xuống dòng \n để phân tách các dòng trong file văn bản
	resultsText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	
	// Ép kiểu sang []byte và ghi file
	os.WriteFile("results.txt", []byte(resultsText), 0644)
}
```

### Giải thích các ký hiệu định dạng:
* `%.1f`: Định dạng số thực có 1 chữ số sau dấu phẩy thập phân.
* `%.3f`: Định dạng số thực có 3 chữ số sau dấu phẩy thập phân.
* `\n`: Ký tự xuống dòng (Newline).

---

## 💻 3. Kết Quả Lưu Trữ Trong Tệp `results.txt`

Sau khi chạy thành công ứng dụng với các dữ liệu hợp lệ, tệp `results.txt` được tạo ra sẽ có nội dung như sau:
```text
EBT: 12000.0
Profit: 10200.0
Ratio: 1.176
```
