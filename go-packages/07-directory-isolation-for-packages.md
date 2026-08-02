# Bài 07: Tách Biệt Thư Mục Cho Mỗi Package & Giải Quyết Lỗi Gói Trùng Lặp

> [!IMPORTANT]
> *Một trong những quy tắc nghiêm ngặt nhất của Go về cấu trúc mã nguồn là: **Mỗi thư mục chỉ được phép chứa duy nhất một Package**. Bài viết này sẽ phân tích lỗi biên dịch phát sinh khi vi phạm quy tắc này và giải pháp cô lập thư mục.*

---

## 🚫 1. Lỗi Biên Dịch Khi Chứa Nhiều Package Trong Cùng Thư Mục

Giả sử bạn tạo file `fileops.go` chứa khai báo:
```go
package fileops
```
Nhưng bạn đặt file này chung một thư mục gốc `/04-bank/` với file `bank.go` (đang khai báo `package main`). 

Khi chạy ứng dụng, Go Compiler sẽ lập tức báo lỗi và dừng biên dịch:
```text
found packages main and fileops in /home/thuong/Desktop/go-learning/learning/04-bank
```
**Quy tắc của Go:** Tất cả các file Go nằm chung một thư mục vật lý bắt buộc phải có cùng một tên package (ngoại trừ các file kiểm thử kết thúc bằng `_test.go`). Bạn không thể định nghĩa nhiều package song song trong cùng một thư mục.

---

## 📂 2. Cô Lập Gói Bằng Thư Mục Con (Subfolder)

Để giải quyết lỗi trên và tạo ra một package độc lập mới, ta bắt buộc phải tạo một thư mục con tương ứng:

1. **Tạo thư mục con**: Tạo thư mục `/fileops/` bên trong `/04-bank/`.
2. **Di chuyển file**: Di chuyển file `fileops.go` vào trong thư mục `/fileops/` này.

Cấu trúc thư mục chuẩn sau khi cô lập:
```text
04-bank/
├── bank.go             (package main)
├── communication.go    (package main)
├── go.mod              (module example.com/bank)
└── fileops/
    └── fileops.go      (package fileops)
```
*Lưu ý: Tên của file `.go` bên trong gói con không bắt buộc phải trùng tên với thư mục (bạn có thể đặt là `file.go` hay `helper.go`), nhưng dòng khai báo đầu file bắt buộc phải trùng với tên gói (tên thư mục) là `package fileops`.*

---

## 🔑 3. Cách Kết Nối Lại Khi Bị Lỗi "Hàm Không Tồn Tại" (Undefined)

Khi di chuyển các hàm `getFloatFromFile` và `writeFloatToFile` sang gói con `fileops`, file `bank.go` sẽ báo lỗi undefined do tầm vực gói `main` không tự động nhìn thấy gói khác. 

Chúng ta giải quyết lỗi kết nối này bằng hai bước:
1. **Xuất khẩu hàm (Exporting)**: Viết hoa chữ cái đầu tiên của các hàm trong `fileops.go` để chuyển chúng thành Hàm Public:
   * `getFloatFromFile` $\rightarrow$ `GetFloatFromFile`
   * `writeFloatToFile` $\rightarrow$ `WriteFloatToFile`
2. **Import & Gọi hàm qua tiền tố**: Thêm gói con vào danh sách import của `bank.go` và gọi hàm kèm tiền tố `fileops.`:
   ```go
   import "example.com/bank/fileops"
   
   // ...
   fileops.GetFloatFromFile(accountBalanceFile)
   ```
