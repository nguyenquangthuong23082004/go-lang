# Bài 30: Thực Hành Tự Giải Quyết Nhánh Rút Tiền (Withdraw Money)

> [!IMPORTANT]
> *Để củng cố kiến thức về `else if` và các toán tử gán nhanh, thử thách đặt ra là tự viết thêm chức năng Rút tiền (Withdraw) bằng cách trừ bớt số tiền trong tài khoản.*

---

## 🛠️ 1. Lời Giải Thách Thức: Nhánh Rút Tiền

Chúng ta thêm một nhánh `else if` mới kiểm tra xem `choice == 3` hay không:

```go
else if choice == 3 {
	fmt.Print("Your withdrawal: ")
	var withdrawalAmount float64
	fmt.Scan(&withdrawalAmount)
	
	accountBalance -= withdrawalAmount // Trừ số dư tài khoản bằng toán tử -=
	fmt.Println("Balance updated! New amount:", accountBalance)
}
```

---

## 💾 2. Khái Niệm Lưu Trữ Tạm Thời (In-Memory State)

Khi chạy thử chương trình:
1. Bạn chọn **Nạp tiền (2)** thêm `$200` $\rightarrow$ Số dư mới là `$1200`.
2. Chương trình kết thúc.
3. Bạn chạy lại chương trình và chọn **Rút tiền (3)** bớt `$200` $\rightarrow$ Số dư hiện tại báo về vẫn là `$800` (giảm từ mức gốc `$1000.0` ban đầu chứ không phải từ `$1200` trước đó).

### ⚙️ Tại sao dữ liệu cũ bị mất?
* Khi chương trình kết thúc, tất cả các vùng nhớ chứa biến (như `accountBalance`) sẽ bị hệ điều hành giải phóng.
* Mỗi lần chạy lệnh `go run .`, chương trình khởi chạy lại từ đầu và gán lại mức mặc định `$1000.0`.
* **Giải pháp:** Để giữ lại số tiền giữa các lần chạy, chúng ta buộc phải lưu số dư vào một file dữ liệu (`.txt`, `.json`) hoặc lưu vào cơ sở dữ liệu (Database). Chúng ta sẽ học cách xử lý việc này trong các bài học sau.
