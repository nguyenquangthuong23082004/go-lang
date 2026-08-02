# Bài 02: Chia Nhỏ Code Thành Nhiều File Trong Cùng Một Package

> [!NOTE]
> *Khi dự án phát triển lớn hơn, việc gom toàn bộ code vào một file duy nhất sẽ gây khó khăn cho việc đọc và tìm kiếm. Go cho phép chia nhỏ mã nguồn thành nhiều file trong cùng một Package một cách dễ dàng.*

---

## 🗂️ 1. Quy Tắc Đồng Nhất Tên Gói (Package Name)

Tất cả các file nằm trong cùng một thư mục dự án thông thường bắt buộc phải khai báo **cùng một tên gói** ở ngay dòng đầu tiên.

Ví dụ trong thư mục `/04-bank/`:
* File `bank.go` bắt đầu bằng: `package main`
* File mới `communication.go` cũng bắt đầu bằng: `package main`

---

## 🔗 2. Chia Sẻ Hàm Và Tầm Vực (Scope) Trong Cùng Gói

Khi hai file thuộc cùng một gói (ở đây là `package main`):
* Các hàm, biến toàn cục, hằng số hoặc struct được khai báo ở file này sẽ **tự động khả dụng** và có thể gọi được ở file kia.
* Bạn **không cần** import file chứa hàm và **không cần** thêm bất kỳ tiền tố nào khi gọi hàm.

Ví dụ: Hàm `presentOptions()` định nghĩa trong `communication.go`:
```go
func presentOptions() {
    // ...
}
```
Được gọi trực tiếp trong vòng lặp của file `bank.go`:
```go
for {
    presentOptions() // Gọi trực tiếp không cần tiền tố hay import
    // ...
}
```

---

## ⚠️ 3. Lưu Ý Quan Trọng: Import Không Được Chia Sẻ

Mặc dù các hàm và biến được tự động chia sẻ, **các gói import thì không**:
* Nếu file `communication.go` sử dụng thư viện `fmt`, nó bắt buộc phải tự khai báo `import "fmt"` ở đầu file của mình.
* Việc file `bank.go` đã import `fmt` không có nghĩa là `communication.go` được dùng ké mà không khai báo.

---

## ⚙️ 4. Cách Chạy Dự Án Nhiều File

Khi dự án có nhiều file mã nguồn, nếu bạn chạy lệnh chỉ định cụ thể một file:
```bash
go run bank.go
```
Bạn sẽ gặp lỗi biên dịch: `# command-line-arguments ... undefined: presentOptions` vì Go không biên dịch file `communication.go` đi kèm.

### Giải pháp:
Bạn phải chạy lệnh biên dịch toàn bộ thư mục hiện tại:
```bash
go run .
```
Ký tự dấu chấm `.` báo cho Go biết cần phải tìm kiếm, liên kết và chạy toàn bộ các file `.go` có trong thư mục hiện hành.
