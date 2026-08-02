# Bài 39: Xử Lý Lỗi (Error Handling) trong Go

> [!IMPORTANT]
> *Go không sử dụng cơ chế bắt ngoại lệ (Exceptions) kiểu `try-catch` như nhiều ngôn ngữ khác. Thay vào đó, Go coi Lỗi (Errors) như một giá trị thông thường được trả về từ các hàm và yêu cầu lập trình viên kiểm tra lỗi một cách tường minh.*

---

## 🚫 1. Triết Lý Xử Lý Lỗi Của Go

Ở các ngôn ngữ khác, các hành vi sai sót (như không tìm thấy tệp tin) thường trực tiếp sinh ra ngoại lệ và làm sập ứng dụng (crash). 

Trong Go:
* Các hàm có thể xảy ra lỗi sẽ trả về thêm một giá trị cuối cùng có kiểu dữ liệu là **`error`** (một interface tích hợp sẵn).
* Giá trị đặc biệt **`nil`** đại diện cho việc **không có lỗi xảy ra**.
* Nếu giá trị lỗi trả về khác `nil` (`err != nil`), nghĩa là đã có lỗi xảy ra và chúng ta phải xử lý nó.

Cấu trúc kiểm tra lỗi chuẩn trong Go:
```go
value, err := someFunction()
if err != nil {
    // Xử lý khi có lỗi xảy ra
}
```

---

## 🛠️ 2. Tạo Lỗi Tùy Chỉnh Bằng Package `errors`

Để tự tạo ra một đối tượng lỗi mới chứa thông điệp mong muốn, chúng ta sử dụng package **`errors`** và hàm **`errors.New`**:

```go
import "errors"

err := errors.New("Thông điệp báo lỗi của bạn ở đây")
```

---

## 🔄 3. Tích Hợp Xử Lý Lỗi Vào Hàm Đọc File Số Dư

Chúng ta thay đổi hàm `getBalanceFromFile` để nó trả về hai giá trị: `(float64, error)` thay vì chỉ trả về một số thực như trước. Nếu xảy ra lỗi đọc file hoặc chuyển đổi số thực, hàm sẽ trả về giá trị mặc định là `1000` kèm thông báo lỗi tương ứng.

```go
func getBalanceFromFile() (float64, error) {
	// 1. Đọc tệp tin
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		// Trả về số dư mặc định 1000 và lỗi tự định nghĩa
		return 1000, errors.New("failed to find balance file")
	}

	// 2. Chuyển đổi dữ liệu
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		// Trả về số dư mặc định 1000 và lỗi phân tích cú pháp
		return 1000, errors.New("failed to parse stored balance value")
	}

	// Trả về số dư thực tế đọc được và nil (không có lỗi)
	return balance, nil
}
```

---

## 💻 4. Đón Nhận Và Hiển Thị Lỗi Trong Hàm `main`

Tại hàm `main()`, chúng ta hứng lỗi và kiểm tra xem có xảy ra lỗi hay không. Nếu có lỗi (Ví dụ: Lần đầu tiên chạy ứng dụng, file `balance.txt` chưa được tạo), ta in cảnh báo ra màn hình nhưng ứng dụng **vẫn tiếp tục chạy bình thường** với số dư mặc định là `1000`:

```go
func main() {
	var accountBalance, err = getBalanceFromFile()

	// Nếu xảy ra lỗi đọc hoặc phân tích file số dư
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err) // In thông điệp lỗi tự động
		fmt.Println("-----------------")
	}

	fmt.Println("Welcome to Go Bank.")
    // ...
}
```
*Phương pháp này giúp ứng dụng hoạt động vô cùng mạnh mẽ, tránh được các trường hợp crash đột ngột và mang lại trải nghiệm người dùng liền mạch.*
