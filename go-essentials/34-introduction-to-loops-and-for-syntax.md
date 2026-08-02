# Bài 33: Giới Thiệu Vòng Lặp & Cú Pháp for trong Go

> [!NOTE]
> *Để tránh việc ứng dụng bị tắt ngay sau khi thực hiện xong một giao dịch đơn lẻ, chúng ta cần giữ cho chương trình tiếp tục chạy bằng cách sử dụng vòng lặp. Trong Go, `for` là từ khóa vòng lặp duy nhất nhưng vô cùng đa năng.*

---

## 🔁 1. Vòng Lặp `for` - Từ Khóa Vòng Lặp Duy Nhất Trong Go

Khác với các ngôn ngữ lập trình như Java, C++, hay JavaScript vốn có nhiều loại vòng lặp (như `while`, `do-while`, `for`), Go tối giản hóa mọi thứ bằng cách **chỉ cung cấp duy nhất từ khóa `for`**. Tuy nhiên, cú pháp `for` trong Go cực kỳ linh hoạt và có thể đáp ứng mọi nhu cầu lặp.

---

## 🛠️ 2. Cú Pháp Vòng Lặp `for` Tiêu Chuẩn

Cú pháp vòng lặp đếm số lần chạy truyền thống của Go gồm 3 phần được phân tách bởi các dấu chấm phẩy `;`:

```go
for i := 0; i < 2; i++ {
    // Đoạn code được lặp lại
}
```

### Phân tích cấu trúc:
1. **Khởi tạo biến chạy (`i := 0`)**: Khai báo một biến cục bộ tạm thời (thường tên là `i`), được chạy một lần duy nhất trước khi vòng lặp bắt đầu.
2. **Điều kiện lặp (`i < 2`)**: Biểu thức điều kiện kiểm tra trước mỗi lần lặp. Vòng lặp sẽ tiếp tục chạy miễn là điều kiện này trả về `true`.
3. **Bước nhảy/Cập nhật (`i++`)**: Đoạn code được tự động thực thi sau mỗi lần lặp xong.
   * Cú pháp `i++` là viết tắt của việc tăng giá trị của `i` lên 1 đơn vị (`i = i + 1`).
   * Cú pháp `i--` là viết tắt của việc giảm giá trị của `i` đi 1 đơn vị.

---

## 💻 3. Áp Dụng Vòng Lặp Vào Bank Application

Chúng ta bọc toàn bộ mã nguồn hiển thị menu, nhận đầu vào và xử lý rẽ nhánh vào thân của vòng lặp `for` thiết lập chạy đúng 2 lần:

```go
package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	// Vòng lặp chạy đúng 2 lần (i = 0 và i = 1)
	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to Go Bank.")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
		}
	}
}
```

### 🚨 Ràng buộc và Hạn chế:
* **Tự ngắt khi có lỗi:** Trong các nhánh nạp/rút tiền, nếu người dùng nhập số tiền không hợp lệ, câu lệnh `return` sẽ lập tức thoát **toàn bộ hàm `main`**, nghĩa là vòng lặp cũng sẽ kết thúc sớm dù chưa chạy đủ 2 lần.
* **Giới hạn số lần lặp:** Việc thiết lập số lần chạy cứng như `2` hay `200` lần không thực tế cho một ứng dụng ngân hàng thực tế. Chúng ta cần thiết lập một **vòng lặp vô hạn** chạy liên tục và chỉ dừng lại khi người dùng chọn thoát (Exit - lựa chọn 4). Chúng ta sẽ giải quyết bài toán này ở bài học tiếp theo.
