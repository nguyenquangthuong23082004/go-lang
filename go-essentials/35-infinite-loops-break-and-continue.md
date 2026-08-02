# Bài 34: Vòng Lặp Vô Hạn (Infinite Loops), Từ Khóa break & continue

> [!NOTE]
> *Để tạo ra một ứng dụng tương tác liên tục cho đến khi người dùng chủ động yêu cầu dừng, Go cung cấp cú pháp vòng lặp vô hạn `for { ... }`. Để kiểm soát luồng bên trong vòng lặp này, hai từ khóa `break` và `continue` là những công cụ đắc lực.*

---

## 🔁 1. Vòng Lặp Vô Hạn (Infinite Loops) Trong Go

Trong Go, để khai báo một vòng lặp chạy mãi mãi, bạn chỉ cần viết từ khóa `for` mà không đi kèm bất kỳ điều kiện hay bước nhảy nào:

```go
for {
    // Đoạn code này sẽ chạy lặp đi lặp lại vô tận
}
```

### Đưa câu chào mừng ra ngoài vòng lặp:
Trong ứng dụng Bank, lời chào `"Welcome to Go Bank."` chỉ nên xuất hiện duy nhất 1 lần lúc mở ứng dụng. Vì vậy, ta đưa dòng in này lên phía trước, bên ngoài khối `for { ... }`.

---

## 🚪 2. Từ Khóa `break` - Thoát Khỏi Vòng Lặp

Để dừng một vòng lặp vô hạn một cách chủ động, ta sử dụng từ khóa **`break`**:

* **Khác biệt với `return`**:
  * `return` sẽ lập tức thoát khỏi **toàn bộ hàm** hiện tại (hàm `main`), đồng nghĩa với việc toàn bộ dòng code phía sau vòng lặp sẽ bị bỏ qua và không thể thực thi (Unreachable Code).
  * `break` chỉ thoát ra khỏi **khối vòng lặp `for`** hiện tại. Chương trình sẽ tiếp tục chạy các dòng code nằm ngay bên dưới vòng lặp đó.

```go
} else {
	fmt.Println("Goodbye!")
	break // Thoát ra khỏi vòng lặp for
}
// Code ở đây sẽ tiếp tục chạy sau khi break
fmt.Println("Thanks for choosing our bank.")
```

---

## 🔄 3. Từ Khóa `continue` - Bỏ Qua Vòng Lặp Hiện Tại

Khi người dùng nhập dữ liệu sai (ví dụ nạp số tiền âm), việc tắt ứng dụng (dùng `return`) hay thoát ứng dụng (dùng `break`) là quá nghiêm khắc. Thay vào đó, ta muốn bỏ qua giao dịch hiện tại và cho người dùng chọn lại menu giao dịch từ đầu.

Để làm việc này, ta dùng từ khóa **`continue`**:

* Khi gặp `continue`, Go sẽ **bỏ qua tất cả các dòng code còn lại** của lượt lặp hiện tại (không cộng tiền vào tài khoản, không in thông báo thành công).
* Nó lập tức nhảy ngược lại về **đầu vòng lặp** để bắt đầu một lượt lặp mới (in lại menu giao dịch).

```go
if depositAmount <= 0 {
	fmt.Println("Invalid amount. Must be greater than 0.")
	continue // Nhảy lên dòng đầu tiên của vòng lặp for, bỏ qua code cộng tiền phía dưới
}
accountBalance += depositAmount // Dòng này bị bỏ qua nếu gặp continue
```

---

## 💻 4. Mã Nguồn Hoàn Thiện Sau Refactoring

Dưới đây là mã nguồn của `/home/thuong/Desktop/go-learning/learning/04-bank/bank.go`:

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

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Tiếp tục lượt lặp mới, cho người dùng chọn lại
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // Tiếp tục lượt lặp mới, cho người dùng chọn lại
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue // Tiếp tục lượt lặp mới, cho người dùng chọn lại
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break // Thoát hẳn vòng lặp để xuống dòng chào cảm ơn bên dưới
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}
```
