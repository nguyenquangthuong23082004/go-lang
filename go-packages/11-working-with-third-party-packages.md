# Bài 11: Làm Việc Với Thư Viện Bên Thứ Ba (Third-Party Packages)

> [!NOTE]
> *Thư viện chuẩn của Go rất mạnh mẽ, nhưng đôi khi bạn cần các chức năng nâng cao hoặc muốn tiết kiệm thời gian viết code. Go sở hữu một hệ sinh thái cộng đồng rất lớn nơi bạn có thể tải và sử dụng các gói thư viện do người khác viết.*

---

## 🔍 1. Tìm Kiếm Thư Viện Bên Thứ Ba

Trang web chính thức để tìm kiếm và tra cứu tài liệu hướng dẫn của tất cả các gói Go trên thế giới là:
👉 **[pkg.go.dev](https://pkg.go.dev)**

Tại đây, bạn chỉ cần gõ từ khóa mong muốn (ví dụ: `random data` để tạo dữ liệu giả lập), hệ thống sẽ liệt kê các thư viện phù hợp kèm hướng dẫn sử dụng và liên kết đến kho chứa mã nguồn mở (thường là GitHub).

---

## 📥 2. Cách Cài Đặt Thư Viện Với `go get`

Để cài đặt một gói bên thứ ba vào dự án, chúng ta sử dụng lệnh `go get` đi kèm đường dẫn URL của package đó:

```bash
go get github.com/Pallinder/go-randomdata
```

### Điều gì xảy ra khi chạy lệnh này?
1. **Tải mã nguồn**: Go sẽ tải toàn bộ mã nguồn của thư viện về máy tính của bạn (được lưu trữ trong thư mục cache toàn cục của Go).
2. **Cập nhật `go.mod`**: File cấu hình module của dự án sẽ tự động thêm một dòng khai báo phụ thuộc (`require`):
   ```go
   require github.com/Pallinder/go-randomdata v1.2.0
   ```
3. **Tạo/Cập nhật file `go.sum`**: Đây là tệp tin tự động lưu trữ các mã băm bảo mật (cryptographic hashes) của thư viện nhằm đảm bảo tính toàn vẹn của mã nguồn, tránh việc thư viện bị sửa đổi mã độc hại khi bạn tải lại ở máy tính khác.

> [!TIP]
> **Quản lý dự án khi chia sẻ**: Khi bạn gửi dự án này cho người khác hoặc đẩy lên GitHub, bạn **không cần gửi thư mục chứa code của thư viện**. Người nhận chỉ cần chạy lệnh `go get` (không đi kèm tham số đường dẫn) hoặc chạy thẳng `go run .`. Go sẽ tự động đọc file `go.mod` và tải tất cả các thư viện cần thiết về máy của họ.

---

## 🔌 3. Import Và Sử Dụng Thư Viện

Chúng ta sử dụng đường dẫn đầy đủ của thư viện trong danh sách import:

```go
package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata" // Import thư viện bên thứ ba
)
```

Khi gọi hàm, chúng ta sử dụng tên package gốc (thường là từ cuối cùng của đường dẫn):

```go
fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())
```
*(Hàm `PhoneNumber()` của package `randomdata` sẽ sinh ra một số điện thoại ngẫu nhiên mỗi khi chương trình khởi chạy).*

---

## 🛠️ 4. Khắc Phục Lỗi "missing go.sum entry"

Khi bạn sao chép code hoặc tự chỉnh sửa file `go.mod` bằng tay để thêm thư viện mới mà chưa chạy lệnh tải chính thức, Go Compiler khi chạy lệnh `go run .` sẽ báo lỗi:

```text
bank.go:7:2: missing go.sum entry for module providing package github.com/Pallinder/go-randomdata (imported by example.com/bank); to add:
        go get example.com/bank
```

### Nguyên nhân:
Mã nguồn thư viện chưa được tải về máy của bạn và chưa có mã xác thực (checksum) được đăng ký lưu trữ trong tệp tin `go.sum` của dự án.

### Giải pháp:
Bạn cần mở Terminal ngay tại thư mục chứa file `go.mod` và chạy một trong hai lệnh sau để đồng bộ:

1. **`go mod tidy` (Khuyên dùng)**:
   Quét toàn bộ mã nguồn Go trong dự án, tự động tải các gói import còn thiếu, tạo/cập nhật cả `go.mod` lẫn `go.sum` một cách đồng bộ và xóa bỏ các gói không sử dụng nữa.
   ```bash
   go mod tidy
   ```

2. **`go get <đường-dẫn-gói>`**:
   Tải trực tiếp gói cụ thể được chỉ định, cập nhật `go.mod` và `go.sum` tương ứng.
   ```bash
   go get github.com/Pallinder/go-randomdata
   ```
