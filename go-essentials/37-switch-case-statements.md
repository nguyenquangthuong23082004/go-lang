# Bài 36: Cấu Trúc Switch-Case trong Go

> [!NOTE]
> *Khi cần so sánh giá trị của một biến duy nhất với nhiều trường hợp cụ thể khác nhau (như biến `choice`), việc sử dụng cấu trúc `switch-case` sẽ giúp mã nguồn rõ ràng, tường minh và dễ đọc hơn nhiều so với chuỗi `if-else if` dài dòng.*

---

## 🔀 1. Cú Pháp Switch-Case Cơ Bản

Cấu trúc `switch` đánh giá một giá trị và so khớp nó với các nhánh `case` tương ứng:

```go
switch <biến_cần_so_khớp> {
case <giá_trị_1>:
    // Chạy khi biến = giá trị 1
case <giá_trị_2>:
    // Chạy khi biến = giá trị 2
default:
    // Nhánh catch-all chạy khi không khớp bất kỳ case nào
}
```

---

## ⚙️ 2. Các Đặc Điểm Độc Đáo Của Switch Trong Go

Mặc dù cú pháp trông giống nhiều ngôn ngữ khác, `switch` trong Go có hai điểm khác biệt rất lớn mà lập trình viên cần lưu ý:

### A. Không tự động chạy tràn (No Fallthrough)
* Ở các ngôn ngữ khác (như JavaScript, C++, Java), nếu bạn quên viết câu lệnh `break` ở cuối mỗi case, chương trình sẽ chạy tràn xuống thực thi tiếp code của case phía dưới.
* **Trong Go:** Go mặc định tự ngắt sau khi chạy xong case khớp. Bạn **không cần** ghi câu lệnh `break` ở cuối mỗi case.

### B. Hành vi của từ khóa `break` thay đổi
* Vì Go tự động ngắt case, từ khóa `break` ít khi được sử dụng trong `switch`.
* Tuy nhiên, nếu bạn viết từ khóa `break` bên trong một `case`, nó sẽ chỉ **thoát ra khỏi cấu trúc `switch`**, chứ **không thoát ra khỏi vòng lặp `for`** đang bao quanh bên ngoài.
* **Hệ quả đối với Bank Application:** Chúng ta không thể sử dụng `break` để thoát vòng lặp vô hạn `for` từ bên trong `switch`. Thay vào đó, ta sử dụng từ khóa `return` để kết thúc toàn bộ hàm `main()`, qua đó dừng chương trình.

---

## 💻 3. Áp Dụng Thực Tế Vào Bank Application

Mã nguồn file `bank.go` sau khi chuyển đổi sang sử dụng `switch-case`:

```go
package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank.")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // continue này tác động lên vòng lặp for bên ngoài
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			// Xử lý khi chọn 4 hoặc bất cứ số/lỗi nhập nào khác
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return // Thoát hẳn hàm main để dừng chương trình
		}
	}
}
```
