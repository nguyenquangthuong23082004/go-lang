# Bài 06: Giải Thích Đường Dẫn Import (Import Paths) trong Go

> [!NOTE]
> *Khi làm việc với các gói (packages) nội bộ tự tạo trong Go, chúng ta luôn phải sử dụng các đường dẫn import trông có vẻ phức tạp như `example.com/bank/fileops`. Bài học này giúp bạn hiểu rõ bản chất tại sao Go lại yêu cầu cú pháp này.*

---

## 📐 1. Công Thức Xác Định Đường Dẫn Import

Mọi đường dẫn import gói nội bộ trong Go được cấu thành từ hai phần chính:

$$\text{Đường dẫn Import} = \text{[Tên Module gốc]} + \text{[/Đường dẫn đến thư mục con của gói]}$$

### Phân tích cụ thể dự án Bank:
1. **Tên Module gốc (`example.com/bank`)**: Được cấu hình trong file `go.mod` nằm ở thư mục dự án `/04-bank/`:
   ```go
   module example.com/bank
   ```
2. **Đường dẫn thư mục con (`/fileops`)**: Thư mục chứa gói `fileops` nằm ngay trong dự án.

$\Rightarrow$ Đường dẫn import đầy đủ là: `"example.com/bank/fileops"`

---

## 🌐 2. Tại Sao Lại Cần Tiền Tố Tên Miền (Ví dụ: `example.com`)?

Go được thiết kế để quản lý các package một cách phi tập trung (decentralized). Go không sử dụng một kho lưu trữ tập trung như npm (của Node.js) hay PyPI (của Python). Thay vào đó, Go tải trực tiếp mã nguồn từ internet.

Do đó, tên module được quy ước đặt theo đường dẫn URL của kho chứa mã nguồn trực tuyến (thường là GitHub hoặc Gitlab):
* **Trong dự án thực tế**:
  ```go
  module github.com/ten-tai-khoan/ten-du-an
  ```
  Nếu bạn publish gói này lên GitHub, bất kỳ ai trên thế giới cũng có thể tải về thông qua lệnh `go get github.com/ten-tai-khoan/ten-du-an`.
* **Trong môi trường học tập/thử nghiệm**:
  Do chúng ta chỉ chạy offline cục bộ (local), nên ta có thể dùng bất kỳ tên miền giả lập nào như `example.com/bank` hay chỉ đơn giản là `bank`. Tuy nhiên, việc đặt tên chứa tên miền giả lập giúp chúng ta làm quen với phong cách code thực tế chuẩn của Go.

---

## ⚙️ 3. Điều Gì Xảy Ra Nếu Đổi Tên Module Trong `go.mod`?

Nếu bạn mở file `go.mod` và đổi tên module từ `example.com/bank` thành `my-awesome-bank`:
```go
module my-awesome-bank
```

Thì trong file `bank.go`, bạn **bắt buộc** phải thay đổi đường dẫn import tương ứng:
```go
import (
	"fmt"
	"my-awesome-bank/fileops" // Thay đổi để khớp với go.mod mới
)
```
Nếu không sửa đổi đồng bộ, Go Compiler sẽ không thể tìm thấy gói và báo lỗi biên dịch ngay lập tức.
