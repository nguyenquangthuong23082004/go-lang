# Bài 05: Tạo Package Tự Định Nghĩa & Xuất Khẩu (Exporting)

> [!IMPORTANT]
> *Để tạo một gói tùy chỉnh có thể sử dụng từ các phần khác của ứng dụng, chúng ta cần tuân thủ quy tắc khai báo thư mục, định nghĩa xuất khẩu (Exporting) thông qua cách viết hoa tên hàm, và import thông qua đường dẫn gốc của module.*

---

## 📂 1. Quy Trình Khởi Tạo Package Tùy Chỉnh

1. **Tạo cấu trúc thư mục con**:
   Chúng ta tạo thư mục `/fileops/` nằm bên trong thư mục dự án `/04-bank/`.
2. **Khai báo tên Package**:
   Tạo file `/fileops/fileops.go` và khai báo gói ở dòng đầu tiên:
   ```go
   package fileops
   ```

---

## 🔑 2. Quy Tắc Tầm Vực: Xuất Khẩu (Exported) vs Không Xuất Khẩu (Unexported)

Khác với Java hay C# sử dụng các từ khóa `public`, `private`, Go quyết định tầm vực truy cập của một hàm/biến/struct dựa trên **chữ cái đầu tiên** của tên định danh đó:

* **Viết hoa chữ cái đầu (Exported / Public)**: Hàm hoặc biến sẽ được xuất khẩu ra ngoài và có thể truy cập được từ các gói khác.
  * *Ví dụ:* `WriteFloatToFile`, `GetFloatFromFile` (Bắt đầu bằng chữ cái viết hoa `W` và `G`).
* **Viết thường chữ cái đầu (Unexported / Private)**: Hàm hoặc biến chỉ có giá trị nội bộ bên trong nội tại gói đó, bên ngoài không thể nhìn thấy và gọi được.
  * *Ví dụ:* `writeFloatToFile`, `getFloatFromFile`.

---

## 🔌 3. Cách Import Và Sử Dụng Package Tùy Chỉnh

Để import gói tiện ích nội bộ của dự án, đường dẫn import phải được bắt đầu bằng **tên module** đã khai báo trong file `go.mod`.

### Bước 1: Khởi tạo module trong `go.mod`
File `go.mod` nằm ở thư mục gốc `/04-bank/` khai báo tên module:
```go
module example.com/bank

go 1.21
```

### Bước 2: Khai báo import trong `bank.go`
Sử dụng đường dẫn đầy đủ `tên_module/thư_mục_gói`:
```go
import (
	"fmt"
	"example.com/bank/fileops" // Import gói fileops tự tạo
)
```

### Bước 3: Sử dụng hàm thông qua tiền tố tên gói
Truy cập các hàm public bằng cú pháp `<tên_gói>.<TênHàmViếtHoa>`:
```go
// 1. Đọc file
accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

// 2. Ghi file
fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
```

*Nhờ việc tách gói này, mã nguồn trong file `bank.go` giờ đây vô cùng gọn gàng, và logic giao tiếp tệp tin `fileops` có thể được tái sử dụng một cách linh hoạt.*
