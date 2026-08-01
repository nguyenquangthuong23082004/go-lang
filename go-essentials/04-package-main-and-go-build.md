# Bài 3: Tại Sao Tên Package Là "main" & Tìm Hiểu Về "go build"

> [!NOTE]
> *Trong bài này, chúng ta sẽ giải mã lý do tại sao phải dùng tên gói đặc biệt là `main`, sự khác biệt giữa môi trường phát triển và môi trường thực tế (production), cũng như bước đầu tiếp cận quy trình build ứng dụng.*

---

## 🌟 1. Ý Nghĩa Của Tên Package Đặc Biệt "main"

Về lý thuyết, bạn có thể đặt bất cứ tên gì cho package của mình (giống như package `fmt` của thư viện chuẩn). Tuy nhiên, tên **`main`** có vai trò cực kỳ đặc biệt:

* **Điểm bắt đầu:** Tên `main` báo cho trình biên dịch Go biết rằng đây là **điểm khởi chạy chính (main entry point)** của cả ứng dụng.
* **Tạo file thực thi:** Go chỉ tạo ra một file thực thi chạy độc lập (executable binary) nếu file đó thuộc về `package main` và có chứa hàm `func main()`. Nếu đặt tên khác (ví dụ: `package app`), Go Compiler sẽ hiểu đây chỉ là một thư viện hỗ trợ (library/package bổ trợ) chứ không phải là chương trình chạy chính.

---

## ⚙️ 2. Sự Khác Biệt Giữa Phát Triển (Development) và Đóng Gói (Production)

Trong lập trình thực tế, chúng ta có hai cách chạy mã nguồn Go:

### 🛠️ Cách 1: Sử dụng `go run` (Tiện lợi khi lập trình)
* **Mục đích:** Viết code đến đâu chạy thử đến đó để kiểm tra lỗi nhanh.
* **Cơ chế:** Biên dịch code ra một file nhị phân tạm thời trong thư mục hệ thống, chạy file đó rồi tự động xóa đi sau khi hoàn thành.
* **Hạn chế:** Chỉ chạy được trên máy tính đã cài đặt sẵn Go SDK (Go Compiler).

### 🚀 Cách 2: Sử dụng `go build` (Khi hoàn thiện và phát hành)
* **Mục đích:** Biên dịch toàn bộ mã nguồn thành một file chạy nhị phân duy nhất (Ví dụ: `app.exe` trên Windows hoặc file nhị phân không đuôi trên Linux/macOS).
* **Ưu điểm vượt trội:** Người dùng cuối (hoặc máy chủ triển khai) **không cần cài đặt Go** vẫn có thể mở file này chạy trực tiếp.
* **Phân phối:** Bạn chỉ cần gửi file chạy này đi mà không cần gửi kèm mã nguồn gốc.

---

## 🚨 Lỗi "Fails to Find a Main Module" khi chạy `go build`

Khi bạn thử chạy lệnh `go build` trong thư mục dự án của mình, bạn có thể gặp phải lỗi tương tự như sau:
```text
go: cannot find main module; see 'go help modules'
```

### Tại sao lỗi này xảy ra?
Kể từ các phiên bản Go hiện đại, Go quản lý dự án và các thư viện liên quan thông qua một hệ thống gọi là **Go Modules**. Lỗi trên xuất hiện vì Go không xác định được dự án của bạn nằm trong module nào (chưa được khởi tạo).

Chúng ta sẽ tìm hiểu cách khởi tạo một module và giải quyết triệt để lỗi này trong bài học tiếp theo!
