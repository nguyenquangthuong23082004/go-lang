# Bài 12: Tổng Kết Phần Học - Go Packages (Quản Lý Gói & Thư Viện)

> [!NOTE]
> *Chúc mừng bạn đã hoàn thành phần học về Go Packages! Việc hiểu và làm chủ cách tổ chức mã nguồn bằng package giúp bạn xây dựng các dự án Go chuyên nghiệp, dễ bảo trì và có khả năng mở rộng tốt.*

---

## 📊 Bảng Tổng Kết Kiến Thức Về Packages

| Khái niệm cốt lõi | Quy tắc & Cách thực hiện | Ví dụ / Lệnh |
| :--- | :--- | :--- |
| **Nhiều file trong một Gói** | Các file chung thư mục phải khai báo cùng tên package. Chia sẻ hàm/biến trực tiếp. Không dùng chung import. | Chạy lệnh: `go run .` |
| **Gói tùy chỉnh (Custom Package)** | Tạo thư mục con riêng biệt. Đặt tên package trùng tên thư mục con. | Gói con `/fileops/` khai báo `package fileops` |
| **Quy tắc viết hoa (Casing Rule)** | Viết HOA đầu từ là Public (gói khác dùng được). Viết thường đầu từ là Private (chỉ dùng nội bộ). | `fileops.WriteFloatToFile` (Public)<br>`fileops.writeFloatToFile` (Private) |
| **Đường dẫn Import nội bộ** | Bắt đầu bằng tên module trong `go.mod`, nối tiếp bằng đường dẫn thư mục con. | `import "example.com/bank/fileops"` |
| **Thư viện bên thứ ba** | Tìm kiếm trên [pkg.go.dev](https://pkg.go.dev). Cài đặt tải về qua `go get`. | `go get github.com/Pallinder/go-randomdata` |
| **Quản lý Dependencies** | `go.mod` quản lý tên gói & phiên bản. `go.sum` lưu mã băm bảo mật. | Đồng bộ hóa: `go mod tidy` |

---

## 💡 Các Điểm Cần Lưu Ý Khi Làm Việc Với Package

1. **Một thư mục - Một package:** Đây là quy tắc bất di dịch của Go. Nếu cố tình đặt 2 package khác tên nhau trong cùng một thư mục vật lý, Go Compiler sẽ báo lỗi.
2. **Quy tắc casing áp dụng cho mọi thứ:** Không chỉ hàm, quy tắc viết hoa để export còn áp dụng cho biến toàn cục, hằng số, tên Struct, và cả các trường (fields) bên trong Struct.
3. **Mã băm trong `go.sum`:** Không bao giờ tự chỉnh sửa file `go.sum`. Hãy để Go tự động tính toán thông qua lệnh `go mod tidy`.

*Bây giờ bạn đã sở hữu một nền tảng lập trình Go cực kỳ vững chắc cùng khả năng tổ chức dự án mô-đun hóa tối ưu. Hãy tiếp tục khám phá các phần tiếp theo của ngôn ngữ tuyệt vời này!*
