# Bài 27: Khởi Đầu Cấu Trúc Điều Khiển & Khởi Tạo Dự Án Bank Application

> [!NOTE]
> *Để học cách kiểm soát luồng chạy của chương trình dựa trên các điều kiện khác nhau (Cấu trúc điều khiển - Control Structures), chúng ta sẽ khởi tạo một dự án thực tế mới: Ứng dụng ngân hàng mô phỏng (Bank Application).*

---

## 📋 1. Ý Tưởng Dự Án Bank Application

Chúng ta sẽ xây dựng một ứng dụng dòng lệnh (CLI) mô phỏng các giao dịch ngân hàng cơ bản:
1. Chào mừng người dùng và đưa ra menu lựa chọn:
   * **1. Check balance** (Kiểm tra số dư)
   * **2. Deposit money** (Nạp tiền)
   * **3. Withdraw money** (Rút tiền)
   * **4. Exit** (Thoát chương trình)
2. Nhận lựa chọn giao dịch từ người dùng (kiểu số nguyên `int`).
3. Thực hiện hành động tương ứng dựa trên lựa chọn đó.
4. Chạy chương trình trong một vòng lặp liên tục cho đến khi người dùng chọn thoát (Exit).
5. Ràng buộc kiểm tra các lỗi logic (Ví dụ: rút nhiều hơn số tiền hiện có, nạp/rút số tiền âm).

---

## 🛠️ 2. Khởi Tạo Dự Án Bank Application

1. **Tạo thư mục dự án mới:** Tạo thư mục `04-bank` nằm trong `/home/thuong/Desktop/go-learning/learning/`.
2. **Khởi tạo Module:**
   ```bash
   go mod init example.com/bank
   ```
3. **Tạo file code:** Tạo file `bank.go` và thiết lập cấu trúc cơ bản ban đầu.

---

## 💻 3. Mã Nguồn Cấu Hình Ban Đầu

Dưới đây là mã nguồn khởi động của file `/home/thuong/Desktop/go-learning/learning/04-bank/bank.go`:

```go
package main

import "fmt"

func main() {
	// Hiển thị lời chào và danh sách menu lựa chọn
	fmt.Println("Welcome to Go Bank.")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	// Sử dụng Print thay vì Println để con trỏ nhập liệu nằm trên cùng dòng
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("Your choice:", choice)
}
```

### Phân Tích Logic:
* Biến `choice` được khai báo ở kiểu số nguyên `int` để nhận giá trị `1, 2, 3, 4` đại diện cho các lựa chọn tương ứng.
* Hàm `fmt.Print("Your choice: ")` được sử dụng để giữ nguyên dấu nháy nhập liệu trên cùng một dòng, tăng tính thẩm mỹ và thân thiện cho giao diện CLI.
* Chương trình tạm thời sẽ thoát ngay sau khi hiển thị lại lựa chọn của bạn. Ở các bài tiếp theo, chúng ta sẽ bắt đầu áp dụng **Cấu trúc điều khiển (`if-else`, vòng lặp `for`)** để xử lý các nghiệp vụ ngân hàng này một cách hoàn chỉnh.
