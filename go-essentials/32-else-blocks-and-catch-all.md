# Bài 31: Khối Catch-All else & Hoàn Thiện Các Nhánh Lựa Chọn

> [!NOTE]
> *Để xử lý trường hợp cuối cùng trong menu (Thoát chương trình) hoặc khi người dùng nhập dữ liệu không hợp lệ, chúng ta sử dụng khối `else` đóng vai trò là nhánh hứng tất cả các trường hợp còn lại (catch-all).*

---

## 🌎 1. Cơ Chế Hoạt Động Của Khối `else`

Trong chuỗi câu lệnh kiểm tra điều kiện:
* Bạn có thể viết nhiều nhánh `else if` để kiểm tra các điều kiện cụ thể.
* Nhưng bạn chỉ được phép định nghĩa **duy nhất một** khối `else` ở cuối cùng của chuỗi lệnh đó.
* Khối `else` này không đi kèm bất kỳ biểu thức điều kiện nào. Nó sẽ tự động được kích hoạt nếu **tất cả** các nhánh `if` và `else if` phía trước đều trả về kết quả sai (`false`).

```go
if choice == 1 {
    // Chạy khi choice = 1
} else if choice == 2 {
    // Chạy khi choice = 2
} else {
    // Chạy khi choice KHÔNG phải 1 hoặc 2
}
```

---

## 👻 2. Xử Lý Khi Nhập Dữ Liệu Lỗi (Invalid Input)

Trong ví dụ của chúng ta:
* Biến `choice` có kiểu là `int`.
* Khi người dùng nhập một ký tự không phải là số (ví dụ: chữ cái `"abc"`), hàm `fmt.Scan` sẽ không thể chuyển đổi ký tự đó thành số nguyên được.
* Lúc này, Go sẽ đặt giá trị của biến `choice` về giá trị mặc định (Zero Value) của kiểu `int`, tức là số `0`.
* Vì `0` không thỏa mãn các điều kiện `choice == 1`, `choice == 2`, hay `choice == 3`, chương trình sẽ tự động chuyển tiếp tới nhánh `else` cuối cùng và in ra lời chào tạm biệt `"Goodbye!"`.

---

## 💻 3. Áp Dụng Khối `else` Cho Lựa Chọn Thoát (Exit)

Chúng ta hoàn thiện chuỗi rẽ nhánh trong file `bank.go` bằng cách thêm khối `else` để xử lý việc thoát chương trình:

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
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your withdrawal: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		// Nhánh catch-all xử lý khi người dùng chọn 4 hoặc nhập bất cứ thứ gì khác
		fmt.Println("Goodbye!")
	}
}
```
*Hiện tại, chương trình vẫn sẽ kết thúc sau khi thực hiện xong một giao dịch. Trong các bài học tiếp theo, chúng ta sẽ áp dụng vòng lặp để giữ cho chương trình luôn chạy.*
