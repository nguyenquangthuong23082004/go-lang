# Bài 01: Giới Thiệu Về Structs (Cấu Trúc Dữ Liệu) Trong Go

> [!NOTE]
> *Cho đến thời điểm này của khóa học, chúng ta đã làm quen với các kiểu dữ liệu cơ bản như `string`, `int`, `float64`, và `bool`. Tuy nhiên, để xây dựng các ứng dụng thực tế phức tạp, Go cung cấp các kiểu dữ liệu nâng cao hơn. Trong phần này, chúng ta sẽ khám phá kiểu dữ liệu nâng cao đầu tiên và cực kỳ quan trọng: **Structs**.*

---

## 🔍 1. Struct Trong Go Là Gì?

Một **Struct (viết tắt của Structure - Cấu trúc)** là một kiểu dữ liệu tùy chỉnh (custom data type) cho phép chúng ta:
* **Gom nhóm các dữ liệu liên quan lại với nhau** thành một tập hợp (collection).
* **Định nghĩa cấu trúc rõ ràng** cho một đối tượng phức tạp trong thế giới thực mà các kiểu dữ liệu đơn lẻ (như chỉ dùng một chuỗi hoặc một số) không thể biểu diễn đầy đủ.

Ví dụ, để biểu diễn một người dùng (User), thay vì quản lý các biến rời rạc như `userName`, `userEmail`, `userAge`, ta có thể gom chúng vào một cấu trúc `User` duy nhất:

```go
type User struct {
    Name  string
    Email string
    Age   int
}
```

---

## 🎯 2. Các Nội Dung Sẽ Học Trong Phần Này

Trong chương này, chúng ta sẽ đi sâu vào tìm hiểu toàn bộ vòng đời và cách vận hành của Structs thông qua các chủ đề:

### A. Khởi tạo và Sử dụng Structs
* Tìm hiểu cú pháp định nghĩa một Struct và các cách khởi tạo (instantiate) một giá trị Struct cụ thể.
* Cách truy cập và thay đổi các trường (fields) bên trong một Struct.

### B. Các Trường Hợp Thực Tế Nên Dùng Structs
* Phân tích các bài toán thực tế và cách áp dụng Structs để giải quyết việc tổ chức dữ liệu một cách tối ưu.

### C. Đính Kèm Phương Thức (Methods) Vào Structs
* Structs không chỉ đơn thuần là nơi chứa dữ liệu. Chúng ta sẽ học cách gắn các hàm (functions) đặc biệt vào Structs, được gọi là **Methods (Phương thức)**.
* Cách hoạt động của Value Receivers và Pointer Receivers khi định nghĩa Method.

---

## 💡 3. Tại Sao Structs Lại Quan Trọng?

* **Tính Đóng Gói (Encapsulation):** Giúp giữ cho code sạch sẽ, dễ đọc và dễ bảo trì bằng cách nhóm các biến có liên quan logic lại với nhau.
* **Mô Hình Hóa Dữ Liệu:** Giúp ánh xạ trực tiếp các khái niệm ngoài đời thực (sản phẩm, hóa đơn, người dùng) vào trong mã nguồn một cách tự nhiên.
* **Nền Tảng Cho Lập Trình Hướng Đối Tượng (OOP) Trong Go:** Go không có `class`, nhưng Structs kết hợp với Methods chính là cách Go hiện thực hóa các khái niệm của OOP.
