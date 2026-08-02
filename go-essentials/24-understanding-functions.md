# Bài 23: Đi Sâu Tìm Hiểu Về Hàm (Functions) Trong Go

> [!NOTE]
> *Hàm là một khối mã nguồn được đóng gói dùng để thực hiện một tác vụ cụ thể và chỉ chạy khi được gọi (code-on-demand). Việc chia nhỏ chương trình thành các hàm giúp mã nguồn sạch sẽ, dễ đọc và dễ tái sử dụng.*

---

## 🧩 1. Hàm (Function) Là Gì?

Hàm giúp bạn gom nhóm các dòng code thực hiện chung một mục đích lại với nhau.

* **Hàm đặc biệt `main()`:** Là điểm khởi chạy của mọi ứng dụng Go. Nó được tự động gọi bởi hệ thống khi chương trình bắt đầu chạy.
* **Hàm thư viện chuẩn (như `fmt.Println`):** Là các hàm đã được đội ngũ phát triển Go viết sẵn. Toàn bộ mã nguồn thư viện này đều mở (nằm trong thư mục cài đặt `src/` của Go trên máy tính), bạn có thể dễ dàng truy cập và đọc cách họ tối ưu hóa code.

---

## 🛠️ 2. Khú Pháp Khai Báo Hàm Tự Định Nghĩa (Custom Functions)

Các hàm tự định nghĩa thường được viết ở bên ngoài và nằm phía dưới hàm `main()`. Cú pháp sử dụng từ khóa **`func`**:

```go
func <tên_hàm>(<tham_số_1> <kiểu_1>, <tham_số_2> <kiểu_2>) {
    // Thân hàm chứa logic thực thi
}
```

### Ví dụ thực tế:
```go
func outputText(text string) {
	fmt.Print(text)
}
```

### 🚨 Quy tắc khai báo tham số (Parameters):
1. **Tên trước, Kiểu sau:** Tên tham số viết trước, kiểu dữ liệu viết sau (Ví dụ: `text string`).
2. **Khai báo nhiều tham số:** Phân cách các tham số bằng dấu phẩy `,` (Ví dụ: `func printInfo(name string, age int)`).
3. **Cú pháp viết tắt (cùng kiểu):** Nếu các tham số liên tiếp có cùng kiểu dữ liệu, bạn có thể gộp tên lại và chỉ ghi kiểu dữ liệu ở cuối:
   ```go
   func addNumbers(a, b, c float64) { ... }
   ```
4. **Phạm vi biến (Scope):** Tham số truyền vào chỉ có giá trị sử dụng bên trong cặp ngoặc nhọn `{}` của hàm đó.

---

## 🏃 3. Gọi Hàm Và Ràng Buộc Tham Số

Sau khi định nghĩa, bạn có thể gọi hàm đó bất cứ lúc nào trong chương trình bằng cách truyền đối số (arguments) vào cặp ngoặc đơn `()`:

```go
func main() {
    outputText("Nhập số tiền đầu tư: ") // Hợp lệ
}
```

### 🚨 Lỗi truyền thiếu đối số:
Go kiểm soát số lượng và kiểu dữ liệu của đối số truyền vào rất nghiêm ngặt. Nếu gọi hàm mà không truyền đủ số lượng đối số như khai báo:
```go
outputText() // LỖI: not enough arguments in call to outputText
```
Trình biên dịch sẽ lập tức báo lỗi để ngăn ngừa lỗi runtime phát sinh.
