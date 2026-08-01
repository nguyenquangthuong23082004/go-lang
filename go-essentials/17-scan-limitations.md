# Bài 16: Hạn Chế Của Hàm fmt.Scan() Trong Go

> [!NOTE]
> *Mặc dù `fmt.Scan()` là cách nhanh nhất và đơn giản nhất để nhận dữ liệu từ terminal, nó vẫn tồn tại một số hạn chế kỹ thuật đáng lưu ý khi xử lý chuỗi văn bản phức tạp.*

---

## 🚨 1. Hạn Chế Lớn Nhất: Không Nhận Được Nhiều Từ (Khoảng Trắng)

Hàm `fmt.Scan()` xử lý dữ liệu nhập vào dựa trên khoảng trắng (khoảng cách, phím tab, hoặc phím Enter) để làm ký tự phân tách các giá trị.

* **Vấn đề xảy ra:** Nếu bạn yêu cầu người dùng nhập họ và tên (Ví dụ: `Nguyen Van A`) và lưu vào một biến kiểu `string` duy nhất thông qua `fmt.Scan()`, hàm này sẽ **chỉ lấy chữ đầu tiên** (`Nguyen`). Các từ phía sau khoảng trắng (`Van A`) sẽ bị bỏ lại trong bộ đệm (input buffer) và có thể gây lỗi tự động điền ở các câu lệnh nhập liệu tiếp theo.
* **Tác động:** Rất khó để sử dụng `fmt.Scan()` để đọc một dòng văn bản hoàn chỉnh có chứa dấu cách.

---

## 💡 2. Giải Pháp Tạm Thời & Tương Lai

* **Hiện tại:** Trong bài học tính toán đầu tư này, chúng ta chỉ yêu cầu người dùng nhập các **con số riêng lẻ** (ví dụ: `5000`, `5.5`, `3`). Do các con số không chứa khoảng trắng, nên việc sử dụng `fmt.Scan()` hoàn toàn hoạt động hoàn hảo và không gặp bất kỳ lỗi nào.
* **Tương lai:** Trong các phần tiếp theo của khóa học, khi làm việc với các dự án yêu cầu nhập văn bản phức tạp hơn (như nhập chuỗi chứa khoảng trắng), chúng ta sẽ tìm hiểu giải pháp thay thế mạnh mẽ hơn bằng cách sử dụng package **`bufio`** (đọc toàn bộ dòng nhập cho đến khi gặp phím Enter) hoặc dùng hàm **`fmt.Scanln()`**.
