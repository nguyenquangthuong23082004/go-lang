# Bài 32: Câu Lệnh if Lồng Nhau (Nested If) & Thoát Sớm Với Từ Khóa return

> [!NOTE]
> *Để đảm bảo tính toàn vẹn của dữ liệu trong các giao dịch tài chính (ví dụ: không cho phép nạp tiền âm hoặc rút nhiều hơn số dư), chúng ta sử dụng các câu lệnh `if` lồng nhau (Nested If) kết hợp với từ khóa `return` để dừng chương trình ngay khi có lỗi.*

---

## 🛠️ 1. Câu Lệnh `if` Lồng Nhau (Nested `if`)

Khái niệm **Nested `if`** đơn giản là việc đặt một câu lệnh `if` kiểm tra điều kiện con nằm bên trong thân `{}` của một câu lệnh điều kiện cha khác:

```go
} else if choice == 2 {
    // Nhánh cha: Xử lý nạp tiền
    
    if depositAmount <= 0 {
        // Nhánh con: Kiểm tra số tiền hợp lệ
    }
}
```

### Các toán tử so sánh thông dụng trong Go:
* So sánh bằng: `==`
* So sánh khác: `!=`
* Nhỏ hơn / Lớn hơn: `<`, `>`
* Nhỏ hơn hoặc bằng / Lớn hơn hoặc bằng: `<=`, `>=`

---

## 🛑 2. Dừng Thực Thi Hàm Với Từ Khóa `return` Trống (Naked Return)

Trong các bài học trước, bạn đã biết `return` dùng để trả về giá trị từ một hàm. Tuy nhiên, `return` còn có một vai trò quan trọng khác: **Dừng thực thi hàm hiện tại ngay lập tức**.

* Hàm `main()` là hàm không có giá trị trả về (void function).
* Khi điều kiện nhập lỗi xảy ra (ví dụ: tiền nạp âm), chúng ta muốn in ra thông báo lỗi và dừng chương trình, không thực hiện tiếp việc cộng tiền hay in thông báo thành công ở phía dưới.
* Để làm được việc này, ta viết từ khóa `return` trơ trọi (không kèm giá trị) – hay còn gọi là **Naked Return**:

```go
if depositAmount <= 0 {
    fmt.Println("Invalid amount. Must be greater than 0.")
    return // Dừng hàm main() ngay lập tức, bỏ qua tất cả dòng code phía dưới
}
```

*Đây là một mô thức lập trình cực kỳ phổ biến trong Go, được gọi là **Guard Clauses** (Mệnh đề bảo vệ). Nó giúp tránh việc thụt lề quá nhiều và giúp luồng xử lý lỗi sạch sẽ hơn.*

---

## 💻 3. Áp Dụng Thực Tế Vào Bank Application

Mã nguồn `bank.go` sau khi bổ sung các bước ràng buộc dữ liệu:

```go
package main

import "fmt"

func main() {
	var accountBalance = 1000.0

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

		// Kiểm tra nạp tiền âm hoặc bằng 0
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

		// Kiểm tra rút tiền âm hoặc bằng 0
		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		// Kiểm tra rút vượt quá số dư hiện có
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
```
