# Bài 29: Cấu Trúc Rẽ Nhánh else if & Toán Tử Gán Nhanh (+=, -=)

> [!NOTE]
> *Để xử lý nhiều kịch bản lựa chọn khác nhau (ví dụ: Nạp tiền, Rút tiền, Thoát), chúng ta mở rộng câu lệnh `if` bằng cách sử dụng `else if`. Đồng thời, Go cung cấp các toán tử gán nhanh giúp rút gọn công thức cập nhật giá trị.*

---

## 🔀 1. Phân Biệt Nhiều Lệnh `if` Độc Lập vs Chuỗi `else if`

Khi xử lý nhiều điều kiện, việc lựa chọn cấu trúc viết code ảnh hưởng trực tiếp đến hiệu năng và logic chạy của chương trình.

### A. Sử dụng nhiều câu lệnh `if` độc lập:
```go
if choice == 1 { ... }
if choice == 2 { ... }
```
* **Cơ chế:** Go sẽ luôn đánh giá (evaluate) **tất cả** các câu lệnh `if` tuần tự từ trên xuống dưới. Ngay cả khi `choice == 1` là đúng, Go vẫn tiếp tục kiểm tra xem `choice == 2` có đúng hay không. Điều này gây lãng phí tài nguyên không cần thiết khi các điều kiện mang tính loại trừ lẫn nhau (chọn 1 thì chắc chắn không chọn 2).

### B. Sử dụng chuỗi liên kết `else if`:
```go
if choice == 1 { 
    ... 
} else if choice == 2 { 
    ... 
}
```
* **Cơ chế:** Nếu điều kiện đầu tiên (`choice == 1`) thỏa mãn, Go sẽ thực thi khối lệnh tương ứng rồi **bỏ qua toàn bộ** các khối điều kiện `else if` phía sau mà không cần kiểm tra nữa. Đây là cách viết tối ưu nhất cho các lựa chọn mang tính loại trừ lẫn nhau (mutually exclusive).

---

## 📍 2. Phạm Vi Biến Trong Nhánh Điều Kiện (Block Scope)

Khi khai báo một biến số bên trong một nhánh rẽ của câu lệnh `if` hoặc `else if`:
```go
else if choice == 2 {
    var depositAmount float64
    ...
}
```
* Biến `depositAmount` chỉ tồn tại và sử dụng được **bên trong cặp ngoặc nhọn `{}`** của nhánh `choice == 2`.
* Ở ngoài nhánh này (ví dụ ở hàm `main` hoặc nhánh `choice == 1`), biến này hoàn toàn không tồn tại. Điều này giúp tối ưu dung lượng bộ nhớ.

---

## 🛠️ 3. Toán Tử Gán Nhanh (`+=`, `-=`)

Khi bạn muốn cập nhật giá trị của một biến hiện có bằng cách cộng/trừ thêm một lượng khác:

* **Cách viết đầy đủ:**
  ```go
  accountBalance = accountBalance + depositAmount
  ```
* **Cách viết rút gọn (Shorthand):**
  ```go
  accountBalance += depositAmount
  ```

Tương tự, khi rút tiền ta sử dụng toán tử `-=`:
```go
accountBalance -= withdrawAmount // Tương đương: accountBalance = accountBalance - withdrawAmount
```

---

## 💻 4. Áp Dụng Thực Tế Vào Bank Application

Chúng ta bổ sung nhánh xử lý lựa chọn `2` (Nạp tiền - Deposit) vào file `bank.go`:

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
		
		accountBalance += depositAmount // Cập nhật số dư tài khoản
		fmt.Println("Balance updated! New amount:", accountBalance)
	}
}
```
