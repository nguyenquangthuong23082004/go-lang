# Bài 38: Đọc Dữ Liệu Từ File (Reading Data from Files) trong Go

> [!NOTE]
> *Để khôi phục lại số dư tài khoản của khách hàng mỗi khi ứng dụng khởi động lại, chúng ta sử dụng hàm `os.ReadFile` để đọc dữ liệu từ tệp tin `balance.txt` và chuyển đổi nó ngược về kiểu số thực `float64`.*

---

## 🛠️ 1. Khai Báo Hằng Số Ở Cấp Độ Package (Package-level Constant)

Để tránh việc lặp đi lặp lại tên tệp tin `"balance.txt"` ở nhiều hàm khác nhau (hàm đọc và hàm ghi), ta nên khai báo một hằng số toàn cục ở cấp độ file:

```go
const accountBalanceFile = "balance.txt"
```
* **Lợi ích:** Khi cần thay đổi tên file lưu trữ, bạn chỉ cần chỉnh sửa một nơi duy nhất tại khai báo hằng số này thay vì đi tìm kiếm và thay đổi ở nhiều nơi trong code, giúp mã nguồn dễ bảo trì hơn.

---

## 📖 2. Đọc Tệp Tin Với Hàm `os.ReadFile`

Hàm **`os.ReadFile`** trong package `os` đọc toàn bộ nội dung của tệp tin được chỉ định:
```go
data, err := os.ReadFile(filename string)
```
* Hàm trả về hai giá trị:
  1. `data` (`[]byte`): Mảng chứa các byte dữ liệu đọc được từ file.
  2. `err` (`error`): Đối tượng chứa thông tin lỗi nếu có (ví dụ: không tìm thấy file).
* **Bỏ qua lỗi tạm thời bằng ký tự gạch dưới (`_`)**:
  Nếu tạm thời chưa muốn xử lý lỗi, Go yêu cầu bạn dùng ký tự `_` (Blank Identifier) để hứng giá trị lỗi đó. Nếu không dùng `_` mà tạo biến nhưng không sử dụng, Go sẽ báo lỗi biên dịch (Compiler Error).
  ```go
  data, _ := os.ReadFile(accountBalanceFile)
  ```

---

## 🔄 3. Chuyển Đổi Byte Slice Sang Kiểu `float64`

Dữ liệu thô đọc ra từ file ở định dạng `[]byte`. Chúng ta cần thực hiện quá trình chuyển đổi ngược lại so với lúc ghi file:

### Bước 1: Chuyển `[]byte` sang chuỗi (`string`)
Sử dụng cú pháp ép kiểu của hàm `string(...)`:
```go
balanceText := string(data) // Ví dụ: []byte{'1', '2', '0', '0'} -> "1200" (string)
```

### Bước 2: Chuyển `string` sang số thực (`float64`)
Chúng ta không thể dùng ép kiểu trực tiếp `float64(balanceText)`. Thay vào đó, ta sử dụng package chuẩn **`strconv`** (String Conversions) và hàm **`ParseFloat`**:
```go
balance, _ := strconv.ParseFloat(balanceText, 64)
```
* Hàm `strconv.ParseFloat` nhận vào:
  1. Chuỗi ký tự cần chuyển đổi (`balanceText`).
  2. Độ chính xác số thực: `64` cho kiểu `float64` (mặc định) hoặc `32` cho kiểu `float32`.
* Hàm trả về hai giá trị: kết quả số thực và lỗi. Chúng ta tạm thời dùng `_` để bỏ qua lỗi.

---

## 💻 4. Hàm Đọc Số Dư Hoàn Chỉnh

Dưới đây là hàm `getBalanceFromFile` được viết trong file `bank.go`:

```go
func getBalanceFromFile() float64 {
	// 1. Đọc mảng byte thô từ file
	data, _ := os.ReadFile(accountBalanceFile)
	
	// 2. Chuyển đổi mảng byte sang string
	balanceText := string(data)
	
	// 3. Phân tích chuỗi string thành số float64 và trả về
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}
```

Chúng ta gọi hàm này để khởi tạo biến `accountBalance` ở đầu hàm `main()`:
```go
func main() {
	var accountBalance = getBalanceFromFile()
	// ...
}
```
*Lưu ý: Đoạn mã này hoạt động tốt với giả định tệp tin `balance.txt` đã tồn tại trước đó trên hệ thống và chứa dữ liệu số thực hợp lệ. Nếu file chưa tồn tại, hàm `os.ReadFile` sẽ báo lỗi và ứng dụng có thể gặp vấn đề đọc dữ liệu trống. Chúng ta sẽ giải quyết việc xử lý lỗi này ở các bài học tiếp theo.*
