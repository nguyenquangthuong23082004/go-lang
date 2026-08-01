# Bài 3: Thiết Lập Môi Trường Phát Triển Go Trên VS Code

> [!NOTE]
> *Để viết code Go hiệu quả, chúng ta sẽ sử dụng trình soạn thảo **VS Code** kết hợp với **Go Extension**. Dưới đây là hướng dẫn từng bước để thiết lập không gian làm việc của bạn.*

---

## 🛠️ Quy Trình Thiết Lập Không Gian Làm Việc

### 📁 Bước 1: Tạo thư mục làm việc
Trước hết, hãy tạo một thư mục trống trên máy tính để chứa mã nguồn dự án của bạn (ví dụ: `go-learning`).

### 💻 Bước 2: Mở thư mục bằng VS Code
Mở thư mục vừa tạo bằng VS Code:
* Mở VS Code -> Chọn **File** -> **Open Folder...** -> Chọn thư mục vừa tạo.

### 📄 Bước 3: Tạo file mã nguồn Go đầu tiên
Tạo một file mới bên trong thư mục đó. 
* Tên file có thể tùy chọn (ví dụ: `app.go` hoặc `main.go`) nhưng bắt buộc phải có phần mở rộng là `.go`.
* *Theo quy ước chuẩn trong các dự án Go, file chạy chính thường được đặt tên là `main.go`.*

---

## 🔌 Cài Đặt Go Extension Cho VS Code

Để VS Code hỗ trợ viết code Go tối ưu nhất, hãy cài đặt Extension chính thức từ Go Team:

1. **Nhận diện tự động:** Khi bạn mở một file `.go`, VS Code thường sẽ hiển thị một thông báo gợi ý dưới góc phải màn hình: *"Install Go Extension?"* -> Hãy chọn **Install**.
2. **Cài đặt thủ công:**
   * Mở danh sách Extensions bằng tổ hợp phím `Ctrl + Shift + X` (hoặc `Cmd + Shift + X` trên macOS).
   * Tìm kiếm từ khóa: `Go`
   * Chọn và cài đặt extension được phát hành bởi **Go Team** (nhà phát triển: `golang.go`).

---

## ❓ Go Extension Giúp Gì Cho Bạn?

> [!IMPORTANT]
> *Lưu ý quan trọng: Extension này **không** có nhiệm vụ chạy chương trình Go (việc đó do **Go Compiler** đảm nhận). Nhiệm vụ của nó là cung cấp các tính năng hỗ trợ viết code cực kỳ tiện lợi.*

| Tính năng | Mô tả chi tiết & Ví dụ |
| :--- | :--- |
| **🎨 Tô màu cú pháp (Syntax Highlighting)** | Giúp mã nguồn dễ nhìn, rõ ràng hơn.<br>Ví dụ: Các từ khóa như `package`, `import`, `func` sẽ có màu sắc riêng biệt. |
| **💡 Gợi ý code (Autocomplete)** | Khi bạn gõ tên thư viện kèm dấu chấm (ví dụ: `fmt.`), extension sẽ tự động hiển thị danh sách các hàm khả dụng như `Println`, `Printf`, `Sprint`, `Errorf`... |
| **🚨 Phát hiện lỗi tức thì** | Nếu gõ sai chính tả (ví dụ: `fmt.Prinln()`), extension sẽ lập tức gạch dưới màu đỏ để bạn sửa ngay mà không cần đợi lúc biên dịch. |
| **🧹 Tự động định dạng (Format)** | Tự động canh lề chuẩn khi lưu file (`Ctrl + S`) thông qua công cụ `gofmt`. <br> *Trước khi lưu:*<br>```go\nfunc main(){\nfmt.Println("Hello")\n}```<br> *Sau khi lưu:*<br>```go\nfunc main() {\n    fmt.Println("Hello")\n}``` |
| **🐞 Gỡ lỗi (Debugging)** | Hỗ trợ đặt điểm dừng (breakpoints), chạy debug từng dòng code, theo dõi sự thay đổi giá trị của biến tương tự như khi debug các ngôn ngữ khác (PHP, JS). |

---

## 📦 Cài Đặt Các Công Cụ Bổ Trợ (Go Tools)

Lần đầu tiên mở file Go sau khi cài đặt extension, VS Code sẽ hiển thị một thông báo yêu cầu cài đặt thêm các công cụ bổ trợ (Go Tools):
> *"Install any other tools?"*

Bạn hãy nhấp vào **Install All** để tự động cài đặt toàn bộ các công cụ hữu ích dưới đây:
* **`gopls`** (Language Server): Bộ não đứng sau tính năng autocomplete, phát hiện lỗi và điều hướng code.
* **`goimports`** (Import tool): Tự động thêm hoặc xóa các gói `import` khi bạn sử dụng hoặc không dùng đến chúng trong code.
* **`dlv`** (Delve Debugger): Công cụ hỗ trợ gỡ lỗi và phân tích tiến trình chạy của Go.
* **`staticcheck`** (Linter): Giúp phân tích code tĩnh để phát hiện các lỗi logic tiềm ẩn hoặc các đoạn code chưa tối ưu.
