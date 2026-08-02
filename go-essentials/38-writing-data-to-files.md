# Bài 37: Ghi Dữ Liệu Ra File (Writing Data to Files) trong Go

> [!NOTE]
> *Để lưu trữ dữ liệu bền vững (persistence) và không bị mất số dư tài khoản mỗi khi ứng dụng Bank tắt, chúng ta cần ghi dữ liệu này ra một file văn bản (`balance.txt`) trong bộ nhớ của ổ đĩa.*

---

## 📦 1. Thư Viện Chuẩn `os` (Operating System)

Để tương tác với hệ điều hành và hệ thống tệp tin (file system), Go cung cấp package tích hợp sẵn mang tên **`os`**.

Chúng ta sử dụng hàm **`os.WriteFile`** để ghi dữ liệu ra tệp:
```go
os.WriteFile(filename string, data []byte, perm os.FileMode) error
```

### Các tham số đầu vào:
1. **`filename` (string)**: Tên file hoặc đường dẫn đến file muốn ghi (Ví dụ: `"balance.txt"`). Nếu file chưa tồn tại, Go sẽ tự động tạo mới.
2. **`data` (`[]byte`)**: Dữ liệu ghi dưới dạng **mảng các byte (byte slice)**. Hàm này không trực tiếp nhận vào chuỗi văn bản (`string`) hay số (`float64`), do đó chúng ta phải thực hiện ép kiểu.
3. **`perm` (os.FileMode)**: Quyền truy cập tệp (File Permissions).

---

## ⚙️ 2. Ép Kiểu Dữ Liệu Sang Byte Slice (`[]byte`)

Để ghi một giá trị số `float64` (như số dư) vào file, chúng ta thực hiện 2 bước chuyển đổi:

### Bước 1: Chuyển float64 thành chuỗi (string)
Sử dụng hàm `fmt.Sprint` từ package `fmt` để chuyển đổi số thực thành chuỗi ký tự văn bản:
```go
balanceText := fmt.Sprint(balance) // Ví dụ: 1000.0 (float64) -> "1000" (string)
```

### Bước 2: Chuyển string thành byte slice (`[]byte`)
Sử dụng cú pháp ép kiểu `[]byte(...)`:
```go
byteData := []byte(balanceText)
```

---

## 🔑 3. Quyền Truy Cập File `0644` (File Permissions)

Khi ghi file, chúng ta truyền tham số quyền là số Octal (hệ bát phân) **`0644`**:
* Đây là cách mã hóa quyền truy cập tệp chuẩn của hệ điều hành Linux/Unix.
* **Ý nghĩa:**
  * **Chủ sở hữu (Owner)**: Có quyền Đọc (Read - 4) và Ghi (Write - 2) $\rightarrow$ $4 + 2 = 6$.
  * **Nhóm sở hữu (Group)**: Chỉ có quyền Đọc $\rightarrow$ $4$.
  * **Người dùng khác (Others)**: Chỉ có quyền Đọc $\rightarrow$ $4$.
* Nhờ quyền `0644`, tệp tin sẽ được bảo vệ an toàn, tránh việc người dùng lạ chỉnh sửa trái phép.

---

## 💻 4. Hàm Ghi Số Dư Ra File Hoàn Chỉnh

Dưới đây là hàm `writeBalanceToFile` được định nghĩa ở cuối file `bank.go`:

```go
func writeBalanceToFile(balance float64) {
	// 1. Chuyển đổi float64 sang string
	balanceText := fmt.Sprint(balance)
	
	// 2. Chuyển string sang []byte và ghi ra file balance.txt với quyền 0644
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
```

Chúng ta gọi hàm này trong `bank.go` ngay sau khi số dư được cập nhật thành công:
* **Nhánh nạp tiền (case 2)**:
  ```go
  accountBalance += depositAmount
  fmt.Println("Balance updated! New amount:", accountBalance)
  writeBalanceToFile(accountBalance)
  ```
* **Nhánh rút tiền (case 3)**:
  ```go
  accountBalance -= withdrawalAmount
  fmt.Println("Balance updated! New amount:", accountBalance)
  writeBalanceToFile(accountBalance)
  ```
*Khi chạy chương trình, thực hiện giao dịch rồi thoát, bạn sẽ thấy file `balance.txt` xuất hiện trong thư mục dự án chứa số dư mới nhất.*
