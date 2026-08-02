# Bài 09: Ứng Dụng Thực Tế Của Con Trỏ Trong Thư Viện Chuẩn (fmt.Scan)

> [!NOTE]
> *Thực ra, bạn đã sử dụng con trỏ ngay từ những bài học đầu tiên mà không hề nhận ra. Ví dụ điển hình nhất chính là hàm nhận thông tin từ bàn phím `fmt.Scan()`.*

---

## 🔍 1. Phân Tích Lệnh `fmt.Scan`

Hãy nhớ lại đoạn mã lấy lựa chọn từ menu của người dùng:
```go
var choice int
fmt.Scan(&choice) // Tại sao phải có dấu & ?
```

Nếu Go không có con trỏ và chúng ta chỉ truyền biến thông thường:
```go
fmt.Scan(choice) // Giả thuyết: không dùng dấu &
```
1. Go sẽ nhân bản giá trị mặc định của `choice` (là `0`) sang một ô nhớ sao chép mới.
2. Hàm `fmt.Scan` nhận dữ liệu nhập từ bàn phím (ví dụ người dùng nhập `2`), lưu số `2` này vào **bản sao**.
3. Hàm kết thúc, bản sao bị hủy. Biến `choice` gốc ở ngoài vẫn bằng `0`. Chương trình không bao giờ nhận được dữ liệu thực.

---

## 🛠️ 2. Cách Con Trỏ Giải Quyết Vấn Đề

Bằng cách thêm toán tử địa chỉ `&` để biến thành `fmt.Scan(&choice)`:
1. Chúng ta gửi **địa chỉ ô nhớ** của biến `choice` cho hàm `fmt.Scan`.
2. Bên trong mã nguồn của thư viện `fmt.Scan`, Go sẽ **giải tham chiếu (`*`)** địa chỉ này và ghi đè giá trị người dùng vừa gõ trực tiếp lên ô nhớ đó.
3. Khi hàm kết thúc, biến `choice` gốc của bạn ở hàm `main` đã được cập nhật thành công với giá trị mới.

---

## 💡 3. Mô Hình Mô Phỏng Đơn Giản

Chúng ta có thể tự viết một hàm mô phỏng hành vi của `fmt.Scan` như sau:

```go
package main

import "fmt"

func main() {
	var choice int // Giá trị mặc định là 0
	
	// Mô phỏng nhập liệu
	mockScan(&choice, 3) 
	
	fmt.Println("Lựa chọn đã nhập:", choice) // Kết quả in ra: 3
}

// Hàm nhận con trỏ và ghi đè dữ liệu giả lập
func mockScan(inputPointer *int, valueToSet int) {
	*inputPointer = valueToSet // Ghi trực tiếp giá trị vào ô nhớ gốc
}
```
*Đây là minh chứng rõ ràng nhất cho thấy con trỏ là công cụ nền tảng giúp Go tương tác với dữ liệu động từ môi trường bên ngoài.*
