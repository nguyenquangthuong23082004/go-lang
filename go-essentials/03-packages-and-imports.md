# Bài 2: Tìm Hiểu Về Packages & Import Trong Go

> [!NOTE]
> *Trong bài học này, chúng ta sẽ làm rõ khái niệm khai báo gói (package) và câu lệnh import. Đây là nền tảng cốt lõi giúp tổ chức, quản lý và tái sử dụng mã nguồn trong Go.*

---

## 📦 1. Khai Báo Gói (Package Clause) là gì?

Ở đầu mọi file mã nguồn Go, dòng đầu tiên luôn là khai báo package (ví dụ: `package main`).

### 🚨 Tại sao đây là bắt buộc?
Nếu bạn xóa dòng khai báo này, Go Extension trong VS Code sẽ lập tức báo lỗi:
> *`expected a package clause, but found import...`*

Điều này có nghĩa trình biên dịch Go luôn yêu cầu khai báo gói là câu lệnh đầu tiên trong file trước khi thực hiện bất kỳ lệnh import hay viết logic nào khác.

---

## 📂 2. Ý Tưởng Đằng Sau Việc Sử Dụng Package

Lập trình Go khuyến khích tổ chức và chia nhỏ mã nguồn thành các gói logic khác nhau để quản lý dự án hiệu quả:
* **Tối thiểu:** Mỗi dự án Go phải có ít nhất **một package**.
* **Đa file:** Một package có thể bao gồm **nhiều file mã nguồn khác nhau**. Ví dụ: Bạn có thể tạo `main.go` và `helper.go` và cả hai đều khai báo `package main` ở đầu file.
* **Tái sử dụng:** Khi chia dự án thành nhiều package (ví dụ: package A, package B), bạn có thể xuất (export) tính năng ở package B để sử dụng (import) vào package A. Điều này giúp các file mã nguồn luôn ngắn gọn và sạch sẽ.

---

## 🔌 3. Câu Lệnh `import` & Thư Viện Chuẩn (Standard Library)

Khi chúng ta viết:
```go
import "fmt"
```
Chúng ta đang yêu cầu trình biên dịch Go import (nhập) gói thư viện tên là `fmt` vào file hiện tại để sử dụng các hàm của nó (như hàm `Print`).

### 📚 Thư viện chuẩn (Standard Library) của Go:
* `fmt` không phải là gói do chúng ta tự viết mà là một phần của thư viện chuẩn khổng lồ tích hợp sẵn đi kèm khi cài đặt Go Compiler.
* Bạn có thể sử dụng các gói này ở mọi dự án Go mà không cần tải thêm bất kỳ thư viện bên thứ ba nào.
* Tra cứu danh sách các package tích hợp sẵn của Go tại trang tài liệu chính thức: 👉 [Go Standard Library Docs](https://pkg.go.dev/std)

---

## ❓ Câu Hỏi Đặt Ra
Việc chia package giúp gom nhóm code rất tốt. Nhưng tại sao ở chương trình đầu tiên này, chúng ta bắt buộc phải đặt tên package là **`main`** mà không phải một tên gọi nào khác? Chúng ta sẽ giải đáp câu hỏi này trong bài học tiếp theo.
