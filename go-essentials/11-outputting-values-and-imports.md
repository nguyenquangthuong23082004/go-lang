# Bài 10: In Dữ Liệu Ra Màn Hình & Cú Pháp Import Nhóm Trong Go

> [!NOTE]
> *Trong bài học này, chúng ta sẽ học cách hiển thị kết quả tính toán một cách rõ ràng bằng `fmt.Println`, tìm hiểu cú pháp import nhiều thư viện chuẩn của Go, và cách chạy nhanh toàn bộ module bằng lệnh `go run .`.*

---

## 📺 1. Hiển Thị Dữ Liệu Xuống Dòng Với `fmt.Println()`

Trong Go, để hiển thị thông tin ra cửa sổ dòng lệnh (Terminal), chúng ta có hai hàm phổ biến từ package `fmt`:

* **`fmt.Print()`**: In dữ liệu ra màn hình nhưng giữ nguyên con trỏ chuột ở cuối dòng (không tự động xuống dòng).
* **`fmt.Println()`**: In dữ liệu ra màn hình và **tự động chèn thêm ký tự xuống dòng** sau khi in xong. Điều này giúp các kết quả in liên tiếp không bị dính vào nhau và dễ đọc hơn.

### Ví dụ sử dụng:
```go
fmt.Println("Giá trị tương lai của khoản đầu tư", khoanDauTu, "là", giaTriTuongLai)
```

---

## 🔌 2. Cú Pháp Import Nhóm (Factored Import Statement)

Khi chương trình của bạn cần sử dụng từ 2 thư viện chuẩn trở lên (ví dụ ở đây là `fmt` và `math`), thay vì khai báo riêng lẻ nhiều dòng import:
```go
import "fmt"
import "math"
```

Go khuyến nghị gộp chung chúng vào một khối import sử dụng cặp ngoặc đơn `()`:
```go
import (
	"fmt"
	"math"
)
```

### 🚨 Một số quy tắc quan trọng:
* **Không dùng dấu phân tách:** Tuyệt đối không dùng dấu phẩy `,` hoặc dấu chấm phẩy `;` ở cuối mỗi dòng package bên trong dấu ngoặc.
* **Tự động định dạng:** Khi bạn lưu file trong VS Code, Go extension sẽ tự động định dạng các dòng import riêng lẻ thành dạng import nhóm này cho bạn.

---

## 🏃 3. Chạy Nhanh Dự Án Bằng Lệnh `go run .`

Trong giai đoạn phát triển (development), việc chạy liên tiếp hai lệnh `go build` rồi `./investment-calculator` sẽ tốn thời gian. Chúng ta có thể tối giản quy trình này.

Vì dự án của chúng ta đã được thiết lập dưới dạng một Go Module (đã có file `go.mod`), bạn chỉ cần chạy lệnh sau trong thư mục dự án:

```bash
go run .
```

* **Ý nghĩa:** Dấu chấm `.` đại diện cho thư mục hiện tại. Lệnh này yêu cầu Go tự động quét toàn bộ thư mục, tìm file thuộc `package main`, biên dịch tạm thời và khởi chạy hàm `main()` ngay lập tức.
* **Kết quả:**
  ```text
  Giá trị tương lai của khoản đầu tư 5000 là 5871.206874999999
  ```

---

## 🔍 5. Kiểm Chứng Kết Quả
Nếu nhập các thông số đầu vào tương tự (Số tiền: `5000`, Lợi suất: `5.5%`, Kỳ hạn: `3 năm`) lên các công cụ tính toán tài chính trực tuyến, bạn sẽ thấy kết quả hoàn toàn trùng khớp với chương trình Go của chúng ta. Chúc mừng bạn đã hoàn thành bài toán tính toán thực tế đầu tiên!
