# Bài 6: Dự Án Đầu Tiên - Bộ Tính Toán Đầu Tư (Investment Calculator)

> [!NOTE]
> *Đã đến lúc tạm biệt chương trình "Hello World" đơn giản để bắt tay vào xây dựng một chương trình thực tế hơn: Bộ tính toán giá trị tương lai của một khoản đầu tư. Thông qua dự án này, bạn sẽ bắt đầu làm quen với cách sử dụng biến, kiểu dữ liệu và hằng số.*

---

## 🎯 1. Bài Toán Thực Tế: Tính Giá Trị Tương Lai Của Khoản Đầu Tư

Chương trình của chúng ta sẽ yêu cầu người dùng cung cấp các dữ liệu đầu vào sau:
1. **Số tiền đầu tư ban đầu** (Investment Amount).
2. **Lợi suất hàng năm mong đợi** (Expected Annual Return Rate).
3. **Thời hạn đầu tư** (Investment Horizon - số năm nắm giữ khoản đầu tư).

Từ các thông số này, chương trình sẽ tính toán và hiển thị ra **giá trị tương lai của khoản đầu tư** sau khi kết thúc kỳ hạn.

---

## 🛠️ 2. Chuẩn Bị Không Gian Làm Việc

Chúng ta sẽ tiếp tục phát triển trên thư mục dự án hiện tại. Trước khi viết code, hãy thực hiện đổi tên các file để cấu trúc dự án trở nên rõ ràng và chuyên nghiệp hơn:

* **Đổi tên file chạy chính:** Đổi tên file từ `app.go` (hoặc `main.go`) thành **`investment_calculator.go`**.
* **Đổi tên Module:** Mở file `go.mod` và cập nhật tên module thành **`investment-calculator`**:
  ```text
  module investment-calculator

  go 1.26.5
  ```

---

## 💡 3. Quy Tắc Đặt Tên File & Module Trong Go

Để viết code chuyên nghiệp như một Go Developer thực thụ, bạn nên tuân thủ các quy chuẩn đặt tên sau:

### 📄 Đối với tên file (.go):
* **Khuyên dùng:** Sử dụng **dấu gạch dưới `_`** để phân tách giữa các từ nếu tên file dài (ví dụ: `investment_calculator.go`).
* **Có thể dùng:** Dấu gạch ngang `-` (ví dụ: `investment-calculator.go`).
* **Tuyệt đối tránh:** Sử dụng **khoảng trắng (blanks/spaces)** trong tên file vì nó sẽ gây lỗi khi chạy lệnh trên Terminal.

### 🌐 Đối với tên Module (Module Path):
* Tên module thường được cấu trúc giống như một đường dẫn URL (ví dụ: `github.com/username/project`). Do đó, việc sử dụng **dấu gạch ngang `-`** trong tên module là hoàn toàn bình thường và rất phổ biến (ví dụ: `investment-calculator`).

---

## 🏁 Bước Tiếp Theo
Sau khi đã đổi tên file và module thành công, chúng ta đã sẵn sàng bắt tay vào viết những dòng code đầu tiên để khai báo các biến lưu trữ thông tin đầu tư!
