# Bài 28: Câu Lệnh Điều Kiện if & Kiểu Dữ Liệu Boolean (bool)

> [!NOTE]
> *Để chương trình thực thi các đoạn code khác nhau trong từng tình huống cụ thể (Ví dụ: in số dư tài khoản chỉ khi người dùng chọn số 1), chúng ta sử dụng cấu trúc rẽ nhánh `if` kết hợp với biểu thức so sánh kiểu Boolean.*

---

## 📐 1. Cú Pháp Câu Lệnh `if` Trong Go

Câu lệnh `if` kiểm tra một điều kiện (biểu thức so sánh) và thực thi khối mã bên trong nếu điều kiện đó trả về giá trị đúng (`true`).

```go
if <điều_kiện> {
    // Code chạy khi điều kiện đúng
}
```

### 🚨 Lưu ý quan trọng về toán tử:
* **Toán tử gán `=`**: Dùng để gán giá trị cho một biến (ví dụ: `accountBalance = 1000.0`).
* **Toán tử so sánh bằng `==`**: Dùng để so sánh xem hai giá trị có bằng nhau hay không (ví dụ: `choice == 1`). Bạn **bắt buộc** phải dùng `==` trong câu lệnh `if`.

---

## 🔀 2. Kết Hợp Nhiều Điều Kiện

Go hỗ trợ các toán tử logic để kiểm tra đồng thời nhiều điều kiện:

* **Toán tử AND (`&&`)**: Cả hai điều kiện đều phải đúng.
  * *Ví dụ:* `choice == 1 && accountBalance > 0`
* **Toán tử OR (`||`)**: Chỉ cần một trong hai điều kiện đúng.
  * *Ví dụ:* `choice == 1 || choice == 2`

---

## 📍 3. Kiểu Dữ Liệu Boolean (`bool`)

Kết quả của một biểu thức so sánh (như `choice == 1`) luôn trả về một giá trị thuộc kiểu **`bool`**. Kiểu dữ liệu Boolean trong Go vô cùng đơn giản, chỉ chứa một trong hai giá trị: **`true`** (đúng) hoặc **`false`** (sai).

Bạn có thể tách biểu thức so sánh ra một biến riêng để làm code dễ đọc hơn:
```go
wantsCheckBalance := choice == 1 // wantsCheckBalance tự động có kiểu bool
if wantsCheckBalance {
    fmt.Println("Your balance is", accountBalance)
}
```
*Thông thường đối với các điều kiện ngắn, ta nên viết trực tiếp `if choice == 1` thay vì tạo biến trung gian.*

---

## 💻 4. Áp Dụng Thực Tế Vào Bank Application

Chúng ta khai báo thêm biến số dư tài khoản `accountBalance` ở kiểu số thực (`float64`) để sẵn sàng cho các nghiệp vụ nạp/rút tiền lẻ (ví dụ: `$5.5`):

```go
package main

import "fmt"

func main() {
	var accountBalance = 1000.0 // Sử dụng số lẻ .0 để Go tự suy luận kiểu float64

	fmt.Println("Welcome to Go Bank.")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// Kiểm tra nếu người dùng chọn 1 (Check balance)
	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	}
}
```
