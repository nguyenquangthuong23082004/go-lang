# Bài 43: Tổng Kết Phần Học - Go Essentials (Các Kiến Thức Cốt Lõi)

> [!NOTE]
> *Chúc mừng bạn đã hoàn thành phần học Go Essentials! Đây là phần học nền tảng vô cùng quan trọng, trang bị đầy đủ các viên gạch đầu tiên để bạn tự tin xây dựng các dự án Go phức tạp hơn.*

---

## 📊 Bảng Tổng Hợp Kiến Thức Cốt Lõi

| Chủ đề / Khái niệm | Cú pháp ví dụ | Mô tả / Ghi chú |
| :--- | :--- | :--- |
| **Khai báo Biến** | `var age int = 22`<br>`score := 9.5` | Dùng `:=` để khai báo nhanh và tự suy luận kiểu dữ liệu (chỉ dùng trong hàm). |
| **Khai báo Hằng số** | `const pi = 3.14` | Giá trị cố định, không thể thay đổi trong suốt quá trình chạy. |
| **Kiểu dữ liệu cơ bản**| `int`, `float64`, `string`, `bool` | Go là ngôn ngữ tĩnh (statically typed), các kiểu dữ liệu rất rõ ràng. |
| **Ép kiểu dữ liệu** | `float64(intValue)` | Phải ép kiểu tường minh khi thực hiện tính toán giữa các kiểu khác nhau. |
| **Nhập/Xuất CLI** | `fmt.Println(...)`<br>`fmt.Scan(&var)` | Dùng con trỏ `&` để truyền địa chỉ bộ nhớ khi nhận dữ liệu từ bàn phím. |
| **Định dạng chuỗi** | `fmt.Printf("%.2f", val)` | Định dạng hiển thị. Dùng `fmt.Sprintf` nếu muốn lưu chuỗi vào biến. |
| **Hàm (Functions)** | `func calc(a, b float64) (float64, error)` | Hỗ trợ gộp kiểu tham số, trả về nhiều giá trị, đặt tên giá trị trả về. |
| **Cấu trúc rẽ nhánh** | `if / else if / else`<br>`switch / case / default` | Khối `switch` trong Go không tự động chạy tràn (no fallthrough). |
| **Vòng lặp (Loops)** | `for i := 0; i < n; i++`<br>`for condition { ... }`<br>`for { ... }` | Go chỉ có duy nhất vòng lặp `for`. Hỗ trợ từ khóa `break` và `continue`. |
| **Lưu trữ Tệp tin** | `os.WriteFile(...)`<br>`os.ReadFile(...)` | Ghi/đọc tệp tin ở dạng byte slice (`[]byte`). Phân quyền file `0644`. |
| **Xử lý Lỗi** | `if err != nil { ... }`<br>`errors.New(...)` | Trả về lỗi như một giá trị thông thường. Dùng `nil` để báo không có lỗi. |
| **Ngắt khẩn cấp** | `panic("Fatal error")` | Dừng ứng dụng lập tức và in vết ngăn xếp (Stack Trace) để gỡ lỗi. |

---

## 🚀 Bước Tiếp Theo

Với nền tảng vững chắc này, chúng ta đã sẵn sàng bước sang các chương tiếp theo để học về:
1. **Con trỏ (Pointers):** Cách quản lý ô nhớ hiệu quả hơn.
2. **Cấu trúc dữ liệu nâng cao:** Structs, Arrays, Slices, Maps.
3. **Lập trình hướng đối tượng trong Go:** Interfaces.
4. **Xử lý đồng thời (Concurrency):** Goroutines & Channels.
