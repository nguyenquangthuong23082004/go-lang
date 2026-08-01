# Bài 1: Go là gì và Tại sao nên chọn Go?

> [!NOTE]
> *Đây là phần giới thiệu đầu tiên giúp định hướng lý do Go được tạo ra và các bài toán thực tế mà Go giải quyết tốt trước khi đi vào chi tiết cú pháp ở những bài tiếp theo.*

---

## 1. Go (Golang) là gì?
* **Định nghĩa:** Go (hay Golang) là một ngôn ngữ lập trình mã nguồn mở (*Open-source*), được phát triển và công bố bởi **Google** (khởi động từ năm 2007, công bố chính thức năm 2009).
* **Mục tiêu:** Được Google tạo ra nhằm tối ưu hóa hiệu năng, khả năng mở rộng (*scalability*) và đơn giản hóa quá trình phát triển phần mềm quy mô lớn.
* **Vị trí trong lịch sử:** Go không quá mới như Rust, nhưng trẻ hơn nhiều so với Java, C++, hay PHP.

---

## 2. Các đặc tính nổi bật giúp Go thành công

### 🔑 Sự Đơn Giản & Rõ Ràng (Simplicity & Clarity)
Go được thiết kế loại bỏ các cú pháp rườm rà để code luôn dễ đọc và dễ bảo trì (Google hướng tới mục tiêu: *sau 5 năm quay lại vẫn dễ dàng hiểu được code đã viết*).
* **Ít keyword, ít "magic":** Hạn chế các tính năng phức tạp ẩn dưới ngôn ngữ.
* **Tự động định dạng:** Sử dụng công cụ `gofmt` tích hợp sẵn để thống nhất phong cách code của toàn cộng đồng.

**So sánh sự tối giản cú pháp:**

* **Trong PHP:**
  ```php
  class User extends Person implements LoginInterface
  {
  }
  ```
* **Trong Go:**
  ```go
  type User struct {
      Name string
  }
  ```

---

### 🔑 Dễ Đọc (Clarity)
Google muốn code Go cực kỳ rõ ràng để bất kỳ ai đọc vào cũng hiểu được ngay.
* Ít từ khóa (keywords).
* Ít "magic" (hạn chế các tính năng chạy ngầm phức tạp khó debug).
* Tự động định dạng code bằng công cụ `gofmt`.

---

### ⚡ Khả Năng Mở Rộng & Hiệu Năng Vượt Trội (Scalability & Performance)
* **Hiệu năng cao:** Chạy nhanh gần bằng C/C++, vượt trội hơn hẳn các ngôn ngữ thông dịch hoặc dynamic như PHP, Python, Ruby hay JavaScript (Node.js ở các tác vụ nặng về CPU).
* **Mở rộng dễ dàng:** Đảm bảo hệ thống hoạt động ổn định từ quy mô nhỏ (100 người dùng) đến cực lớn (hàng chục triệu người dùng).

> **Ví dụ so sánh hiệu năng xử lý request:**
> * **API viết bằng PHP:** ~500 request/giây
> * **API viết bằng Go:** ~3,000 request/giây
> *(Đây là số liệu minh họa, thực tế phụ thuộc nhiều vào cấu trúc ứng dụng)*

---

### 🔀 Xử Lý Đồng Thời Mạnh Mẽ (Concurrency)
Đây là điểm mạnh cốt lõi của Go. Go tích hợp cơ chế **Concurrency** cực kỳ gọn nhẹ thông qua các *Goroutines*.
* **Khái niệm:** Giúp một chương trình xử lý nhiều công việc cùng một lúc một cách song song mà không tốn nhiều tài nguyên hệ thống.
* **Ví dụ luồng xử lý thông thường (Tuần tự):**
  ```text
  User A -> Đọc DB -> Gửi Email -> Upload ảnh -> Lưu Log (Tốn thời gian chờ từng bước)
  ```
* **Xử lý song song trong Go:**
  ```go
  go SendEmail()
  go SaveLog()
  go UploadFile()
  // Chương trình chính vẫn tiếp tục chạy ngay lập tức mà không bị tắc nghẽn (non-blocking)
  ```

---

### 📦 Tích Hợp Sẵn Thư Viện Chuẩn (Batteries Included)
Go đi kèm với một bộ thư viện chuẩn vô cùng mạnh mẽ và đầy đủ, giúp lập trình viên không cần phụ thuộc quá nhiều vào các thư viện bên thứ ba.
* **Có sẵn:** HTTP Server, JSON parser, File system, mã hóa (Crypto), Templates, Quản lý thời gian, Context, và Testing.
* **So sánh việc cài đặt thư viện:**
  * **PHP:** `composer require ...`
  * **Node.js:** `npm install ...`
  * **Go:** Chỉ cần `import "net/http"` là dùng được ngay.

---

### 🛡️ Ngôn Ngữ Kiểu Tĩnh (Statically Typed)
Go bắt buộc xác định rõ kiểu dữ liệu tại thời điểm biên dịch, giúp phát hiện lỗi cực kỳ sớm (*catch many errors pretty early*).

* **Trong PHP (Kiểu động):**
  ```php
  $x = 5;
  $x = "hello"; // Hợp lệ, không báo lỗi khi lưu/chạy
  ```
* **Trong Go (Kiểu tĩnh):**
  ```go
  var x int = 5
  x = "hello" // Trình biên dịch (Compiler) báo lỗi ngay lập tức vì int ≠ string
  ```

* **Lợi ích:**
  ```go
  var age int
  age = "20"
  // Compiler báo lỗi: cannot use string as int
  // Bạn phát hiện và sửa lỗi ngay lập tức trước khi chạy ứng dụng
  ```

---

## 3. Go thường được dùng ở đâu?

| Lĩnh vực | Mô tả & Ví dụ |
| :--- | :--- |
| **Network & Web API** | Xây dựng Chat Server, TCP Server, HTTP APIs tốc độ cao.<br>*(Mô hình: Mobile -> Go API -> PostgreSQL)* |
| **Microservices** | Chia nhỏ hệ thống lớn thành các dịch vụ độc lập như *User Service, Order Service, Payment Service, Notification Service*. |
| **CLI Tools (Công cụ dòng lệnh)** | Các công cụ dòng lệnh cực kỳ nổi tiếng được viết bằng Go: `docker`, `kubectl`, `terraform`, `hugo`... |

---

## 📝 Tóm Tắt Bài Học
Giảng viên muốn bạn ghi nhớ 4 ý chính:
1. **Google sinh ra Go:** Hướng tới mã nguồn mở, đơn giản, dễ đọc và dễ bảo trì lâu dài.
2. **Hiệu năng và Concurrency cực mạnh:** Lựa chọn hoàn hảo cho backend, API và microservices.
3. **Statically Typed (Kiểu tĩnh):** Phát hiện lỗi lập trình ngay khi viết code thông qua Compiler thay vì đợi chạy mới thấy lỗi.
4. **Batteries Included:** Thư viện chuẩn cực tốt, hạn chế cài cắm thêm package bên ngoài giúp dự án gọn nhẹ và an toàn hơn.
