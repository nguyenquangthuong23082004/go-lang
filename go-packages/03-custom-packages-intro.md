# Bài 03: Tổ Chức Dự Án Thành Nhiều Gói (Custom Packages)

> [!NOTE]
> *Khi dự án phát triển và trở nên phức tạp, việc gộp tất cả các file vào cùng gói `package main` không còn là giải pháp tốt nhất. Chia mã nguồn thành các **Package tùy chỉnh (Custom Packages)** giúp phân tách chức năng rõ ràng hơn và cho phép tái sử dụng mã nguồn cho các dự án khác.*

---

## 💡 1. Tại Sao Nên Sử Dụng Nhiều Package?

* **Phân tách trách nhiệm (Separation of Concerns):** Giữ cho logic chính của chương trình (ví dụ: giao diện người dùng, nghiệp vụ ngân hàng) tách biệt hoàn toàn với logic kỹ thuật (như ghi/đọc dữ liệu ra ổ đĩa).
* **Tái sử dụng mã nguồn (Reusability):** Một gói tiện ích (utility package) xử lý file có thể được copy sang các dự án Go khác và hoạt động ngay lập tức mà không cần chỉnh sửa logic chính.
* **Độc lập và mô-đun hóa (Modularity):** Giúp kiểm thử (unit testing) dễ dàng hơn vì mỗi gói đảm nhận một nhiệm vụ riêng biệt.

---

## 📂 2. Cấu Trúc Gói Trong Go

Để định nghĩa một Package mới trong Go:
1. **Tạo một thư mục con** nằm trong thư mục dự án (Ví dụ: Thư mục `/fileops/`).
2. Tên của Package theo quy ước sẽ **trùng với tên thư mục con** đó.
3. Tất cả các file `.go` bên trong thư mục con này sẽ khai báo gói tương ứng ở dòng đầu tiên:
   ```go
   package fileops
   ```

---

## 🗺️ 3. Kế Hoạch Phân Tách Cho Bank Application

Trong các bài tiếp theo, chúng ta sẽ tiến hành tách hai hàm xử lý tệp tin từ `bank.go` ra một gói độc lập:
* Hàm `writeBalanceToFile(balance float64)`
* Hàm `getBalanceFromFile() (float64, error)`

Các hàm này làm nhiệm vụ chung là đọc/ghi dữ liệu số thực (`float64`) ra tệp tin, hoàn toàn có thể tái sử dụng được ở bất cứ đâu. Chúng ta sẽ đặt chúng vào một gói mới có tên là **`fileops`**.
