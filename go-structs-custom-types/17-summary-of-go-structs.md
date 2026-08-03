# Bài 17: Tổng Kết Về Structs (Cấu Trúc Dữ Liệu) Trong Go

> [!NOTE]
> *Tổng kết toàn bộ các kiến thức cốt lõi đã học trong chương Structs và Custom Types.*

---

## 📋 1. Tóm Tắt Các Kiến Thức Cốt Lõi

Trong chương này, chúng ta đã tìm hiểu toàn diện về **Structs** - một trong những tính năng nền tảng và được sử dụng nhiều nhất khi lập trình Go:

1. **Khái niệm & Vai trò:** Struct cho phép chúng ta tự định nghĩa một kiểu dữ liệu phức tạp bằng cách gom nhóm nhiều trường dữ liệu và các hành vi (hàm) liên quan lại với nhau thành một thực thể duy nhất.
2. **Khởi tạo & Hàm khởi tạo (Constructor):**
   * Sử dụng cú pháp Struct Literal (khuyên dùng dạng `Key: Value`).
   * Sử dụng quy ước đặt tên hàm khởi tạo là **`New(...)`** giúp tập trung logic tạo đối tượng và tích hợp tính năng kiểm tra dữ liệu đầu vào (Validation).
3. **Đóng gói & Phân chia Package (Casing):**
   * Go sử dụng quy tắc viết hoa chữ cái đầu để xác định phạm vi hiển thị (Visibility).
   * Để Struct, các trường (fields), và phương thức (methods) có thể truy cập từ package bên ngoài, chúng bắt buộc phải viết hoa chữ cái đầu.
   * Để ẩn thuộc tính (Private), chúng ta viết thường chữ cái đầu tiên của trường dữ liệu.
4. **Phương thức (Methods):**
   * Đính kèm hàm trực tiếp vào Struct bằng **Receiver Argument** (Tham số nhận).
   * **Pointer Receiver (`*User`):** Bắt buộc khi method cần thay đổi dữ liệu của Struct gốc hoặc khi Struct lớn cần tối ưu bộ nhớ.
   * **Value Receiver (`User`):** Sử dụng cho các thao tác chỉ đọc (Read-only) với Struct có kích thước nhỏ.
5. **Nhúng Struct (Struct Embedding):**
   * Go không sử dụng cơ chế kế thừa (Inheritance) mà sử dụng **Composition** thông qua nhúng Struct ẩn danh.
   * Các trường và phương thức của Struct được nhúng sẽ tự động được quảng bá (promoted) trực tiếp lên Struct cha.

---

## 🧠 2. Lưu Ý Quan Trọng Về Hiệu Năng Bộ Nhớ

* **Trả về con trỏ (`*User`):** Giúp tránh việc sao chép (copy) Struct trong bộ nhớ mỗi khi trả về từ hàm khởi tạo. Khuyên dùng cho các Struct phức tạp và có kích thước lớn.
* **Trả về giá trị thường (`Admin`):** Phù hợp với các Struct nhỏ, đơn giản. Việc sao chép các Struct này đôi khi còn hiệu quả và tối ưu hơn so với việc cấp phát vùng nhớ Heap cho con trỏ.
* **Cơ chế tự động giải tham chiếu:** Go tự động chuyển đổi ngầm giữa dạng giá trị và con trỏ khi bạn gọi các phương thức, giúp mã nguồn luôn ngắn gọn và trực quan.
