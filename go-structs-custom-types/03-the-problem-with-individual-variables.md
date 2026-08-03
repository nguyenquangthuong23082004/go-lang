# Bài 03: Vấn Đề Khi Truyền Nhiều Biến Đơn Lẻ Vào Hàm

> [!NOTE]
> *Chúng ta sẽ xem xét ví dụ cụ thể về việc truyền các biến đơn lẻ vào một hàm xuất dữ liệu người dùng (`outputUserDetails`), từ đó thấy rõ các điểm bất lợi như viết code dài dòng, dễ sai sót về thứ tự tham số trước khi chuyển đổi sang sử dụng Structs.*

---

## 💻 1. Cập Nhật Mã Nguồn (`structs.go`)

Để thực hiện tác vụ xuất dữ liệu người dùng một cách tập trung, chúng ta đã thêm hàm `outputUserDetails` nhận vào 3 biến `string`:

```go
package main

import "fmt"

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Gọi hàm xuất dữ liệu người dùng
	outputUserDetails(userFirstName, userLastName, userBirthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Birthdate:", birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
```

---

## ⚠️ 2. Phân Tích Những Bất Lợi Thực Tế

Mặc dù chương trình trên chạy hoàn toàn bình thường, việc tổ chức code như thế này mang lại nhiều nhược điểm nghiêm trọng khi ứng dụng mở rộng:

### A. Dễ Nhầm Lẫn Thứ Tự Tham Số (Error-Prone)
Vì cả 3 tham số truyền vào hàm `outputUserDetails` đều có kiểu dữ liệu là `string`, trình biên dịch Go (Go compiler) không thể phát hiện nếu bạn đảo lộn thứ tự đối số lúc gọi hàm:

```go
// Lỗi logic: Truyền nhầm lastName trước firstName
outputUserDetails(userLastName, userFirstName, userBirthdate)
```
Chương trình vẫn biên dịch và chạy bình thường, nhưng kết quả in ra sẽ bị sai lệch thông tin mà không có bất kỳ thông báo lỗi nào.

### B. Chữ Ký Hàm Dài Dòng & Khó Bảo Trì (Verbose & Hard to Maintain)
Mỗi khi ta cần bổ sung thêm thông tin người dùng mới (ví dụ: `email`, `phoneNumber`), ta phải:
1. Tạo thêm các biến đơn lẻ trong hàm `main()`.
2. Sửa đổi danh sách tham số (parameter list) của hàm `outputUserDetails`.
3. Sửa đổi danh sách đối số (argument list) tại nơi gọi hàm `outputUserDetails`.

### C. Giải Pháp Hợp Lý Nhất
Chúng ta cần một kiểu dữ liệu có thể đóng gói tất cả các thông tin riêng lẻ (`firstName`, `lastName`, `birthdate`) thành một thực thể duy nhất. Đó chính là **Struct**.
