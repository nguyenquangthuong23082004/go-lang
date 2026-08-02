# Bài 35: Cú Pháp Vòng Lặp Điều Kiện (Thay Thế Cho Vòng Lặp While) trong Go

> [!NOTE]
> *Bên cạnh vòng lặp đếm số lần và vòng lặp vô hạn, Go còn cung cấp cú pháp vòng lặp kèm điều kiện logic. Vì Go không có từ khóa `while`, đây chính là giải pháp thay thế hoàn hảo cho vòng lặp `while` trong các ngôn ngữ lập trình khác.*

---

## 🔁 1. Cú Pháp Vòng Lặp Điều Kiện

Cú pháp này chỉ chứa một điều kiện duy nhất nằm sau từ khóa `for`:

```go
for <biểu_thức_điều_kiện> {
    // Đoạn code được lặp lại
}
```

Trong đó, `<biểu_thức_điều_kiện>` là một biểu thức so sánh trả về kiểu Boolean (`true` hoặc `false`) hoặc là một biến kiểu `bool`.

---

## ⚙️ 2. Cơ Chế Hoạt Động

1. Trước mỗi lượt lặp, Go sẽ đánh giá giá trị của biểu thức điều kiện.
2. Nếu điều kiện trả về `true`, toàn bộ khối lệnh trong thân `{}` của vòng lặp sẽ chạy.
3. Nếu điều kiện trả về `false`, vòng lặp sẽ kết thúc ngay lập tức và chương trình chuyển xuống dòng code tiếp theo dưới vòng lặp.

---

## 💡 3. Ví Dụ Thay Thế Vòng Lặp `while`

Để thấy rõ sự tương đồng với vòng lặp `while`, hãy xem ví dụ điều khiển ứng dụng Bank Bank:

```go
package main

import "fmt"

func main() {
	var choice int = 0

	// Vòng lặp chạy liên tục cho đến khi người dùng chọn 4 (Thoát)
	for choice != 4 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Checking balance...")
		} else if choice == 2 {
			fmt.Println("Depositing money...")
		} else if choice == 3 {
			fmt.Println("Withdrawing money...")
		}
	}

	fmt.Println("Goodbye! Program exited.")
}
```

*Trong ví dụ trên, chừng nào biến `choice` chưa phải là số `4`, điều kiện `choice != 4` vẫn luôn đúng (`true`) và vòng lặp tiếp tục chạy. Khi người dùng nhập `4`, điều kiện kiểm tra trở thành `false` và vòng lặp tự động dừng lại mà không cần dùng đến từ khóa `break`.*
