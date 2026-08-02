# Bài 40: Dừng Chương Trình Đột Ngột Với Hàm panic & Tránh Lỗi Nghiêm Trọng

> [!WARNING]
> *Trong trường hợp ứng dụng gặp phải một lỗi nghiêm trọng không thể tiếp tục thực thi (ví dụ: mất kết nối Database cấu hình, file hệ thống quan trọng bị hỏng), Go cung cấp hàm `panic()` để dừng chương trình ngay lập tức kèm theo thông tin chi tiết về lỗi.*

---

## 🚪 1. Dừng Chương Trình Sớm Bằng Câu Lệnh `return`

Cách đơn giản nhất để dừng ứng dụng khi có lỗi xảy ra mà không muốn sử dụng giá trị số dư mặc định là gọi câu lệnh `return` ngay trong khối kiểm tra lỗi của hàm `main()`:

```go
func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		return // Thoát khỏi hàm main, kết thúc ứng dụng một cách êm đẹp
	}
    
    // Code phía dưới sẽ không chạy nếu có lỗi xảy ra
}
```

---

## 🚨 2. Hàm Dừng Khẩn Cấp `panic()`

Nếu muốn kết thúc ứng dụng một cách mạnh mẽ hơn (tương tự như một cú sập nguồn - crash), bạn có thể gọi hàm **`panic()`** tích hợp sẵn:

```go
panic("Can't continue, sorry.")
```

### ⚙️ Hành vi của `panic()`:
1. **Dừng chương trình lập tức**: Luồng thực thi bị ngắt ngay tại dòng gọi `panic()`.
2. **Hiển thị thông điệp lỗi**: In chuỗi ký tự lỗi truyền vào hàm `panic()`.
3. **In Stack Trace (Vết ngăn xếp)**: Hiển thị danh sách các hàm đang chạy dở dang cùng số dòng và tên file cụ thể dẫn tới nơi xảy ra lỗi (ví dụ: dòng số 41). Điều này giúp lập trình viên lần ngược dấu vết lỗi cực kỳ nhanh chóng.

---

## 💡 3. Ví Dụ Về Lỗi Dẫn Đến Crash (Minh Họa)

Nếu chúng ta sửa mã nguồn của hàm `main` để gọi `panic()` khi lỗi đọc file xảy ra:

```go
func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		panic("Can't continue, sorry.") // Phát tín hiệu hoảng loạn (panic)
	}
    // ...
}
```

Khi chạy ứng dụng mà không có file `balance.txt`, giao diện CLI sẽ xuất hiện stack trace tương tự như sau:
```text
ERROR
failed to find balance file
-----------------
panic: Can't continue, sorry.

goroutine 1 [running]:
main.main()
        /home/thuong/Desktop/go-learning/learning/04-bank/bank.go:17 +0xf5
exit status 2
```

---

## ⚖️ 4. Khi Nào Nên Dùng `panic` vs Xử Lý Lỗi Thông Thường?

* **Không lạm dụng `panic`**: Đối với các lỗi thông thường, có thể đoán trước (như người dùng nhập sai số tiền, file dữ liệu chưa được tạo ở lần đầu chạy), ta nên trả về giá trị lỗi (`error`) và xử lý nhẹ nhàng như gán giá trị mặc định hoặc bắt người dùng nhập lại.
* **Chỉ dùng `panic` khi**: Lỗi xảy ra thuộc loại "bất khả kháng" và ứng dụng không thể làm gì khác ngoài việc đóng cửa ngay lập tức để tránh làm hỏng thêm dữ liệu (ví dụ: lỗi phần cứng ổ đĩa, lỗi không thể bind cổng mạng khi khởi động server).
*Trong ứng dụng ngân hàng Bank, chúng ta không cần thiết phải dùng đến `panic` vì việc gán giá trị mặc định là `1000` và cảnh báo lỗi là đã đủ bảo vệ chương trình hoạt động bình thường.*
