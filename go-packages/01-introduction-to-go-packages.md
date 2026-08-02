# Bài 01: Giới Thiệu Về Package & Cách Tổ Chức Dự Án Trong Go

> [!NOTE]
> *Sau khi nắm vững các cú pháp cốt lõi của Go, chúng ta sẽ tiến thêm một bước sâu hơn: Tìm hiểu về **Packages (Gói)**. Đây là cách Go giúp bạn chia nhỏ code, tổ chức thư mục dự án và tái sử dụng mã nguồn một cách khoa học.*

---

## 📦 1. Go Package Là Gì?

Trong Go, một **Package (Gói)** là một tập hợp các file mã nguồn Go (có đuôi `.go`) nằm trong **cùng một thư mục** và có cùng khai báo tên gói ở đầu file (ví dụ: `package main` hoặc `package fmt`).

Cho đến nay, chúng ta đã tiếp xúc với:
* **Package `main`**: Gói đặc biệt chứa hàm `main()`, đóng vai trò là điểm khởi chạy (entry point) của mọi ứng dụng Go có thể thực thi.
* **Các Package chuẩn của Go (Standard Library)**: Như `fmt` (để in/quét), `os` (giao tiếp hệ điều hành), `errors` (quản lý lỗi), `strconv` (chuyển đổi chuỗi).

---

## 🎯 2. Các Nội Dung Sẽ Học Trong Phần Này

Khi dự án phình to, việc viết hàng nghìn dòng code trong duy nhất một file `main.go` là điều bất khả thi và cực kỳ khó bảo trì. Trong phần này, chúng ta sẽ học cách giải quyết vấn đề đó thông qua các chủ đề:

### A. Chia nhỏ mã nguồn ra nhiều file trong cùng một Package
* Cách tách các hàm tiện ích sang các file `.go` khác nhau nhưng vẫn giữ nguyên khai báo `package main`.
* Cách biên dịch và chạy ứng dụng Go khi có nhiều file liên kết với nhau.

### B. Tổ chức dự án thành nhiều Package tùy chỉnh (Custom Packages)
* Cách tạo các thư mục con và khai báo gói riêng biệt (ví dụ: gói `filemanager`, gói `calculator`).
* Các quy tắc đóng gói và quản lý tầm vực hiển thị (Scope) của các biến/hàm giữa các gói.

### C. Import và sử dụng các Package tùy chỉnh
* Cách nhập khẩu (import) gói nội bộ thông qua đường dẫn module được định nghĩa trong `go.mod`.
* Cách sử dụng các hàm được xuất khẩu (Exported Functions) từ gói này sang gói khác.
