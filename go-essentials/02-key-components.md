# Bài 1: Các Thành Phần Cốt Lõi Trong Chương Trình Go

> [!NOTE]
> *Để bắt đầu đi sâu vào Go, hãy cùng phân tích lại chương trình "Hello World" đã viết. Tuy đơn giản, chương trình này chứa đựng đầy đủ các thành phần cơ bản sẽ xuất hiện trong mọi dự án Go của bạn.*

---

## 🔍 Phân Tích Dòng Lệnh Thực Thi Đầu Tiên

Hãy bắt đầu với hàm `main` và câu lệnh bên trong nó:

```go
func main() {
    fmt.Print("Hello World")
}
```

### 1. Hàm (Function) & Lệnh Thực Thi
Dòng code `fmt.Print(...)` thực chất là một **lời gọi hàm (Function Call)**.
* **Hàm là gì?** Bạn có thể hiểu đơn giản hàm giống như một **lệnh (command)** mà chương trình sẽ thực hiện.
* **Chức năng của `Print`:** Khi bạn di chuột qua hàm này trong VS Code, trình soạn thảo sẽ hiển thị mô tả kỹ thuật:
  > *`Print formats using the default formats for its operands and writes to standard output.`*
  
  *Hiểu một cách đơn giản:* Hàm này có nhiệm vụ hiển thị dòng chữ ra **đầu ra chuẩn (standard output)** — trong trường hợp này chính là cửa sổ dòng lệnh (Terminal/Command Line) nơi bạn chạy lệnh `go run`.

---

### 2. Package `fmt` (Thư Viện Tích Hợp Sẵn)
Lệnh `Print` không tự nhiên tồn tại độc lập mà được cung cấp bởi package **`fmt`**.
* **Đặc điểm:** Đây là thư viện chuẩn được tích hợp sẵn trong Go core. Bạn không cần tải hay cài đặt thêm bất kỳ bên thứ ba nào để sử dụng.
* Go có rất nhiều package tích hợp sẵn như thế này để cung cấp các tính năng tiện ích hỗ trợ lập trình viên.

---

### 3. Giá Trị Chuỗi (String Value)
Nội dung nằm trong ngoặc `("Hello World")` truyền vào hàm `Print` được gọi là một **giá trị (value)**, cụ thể là kiểu **chuỗi ký tự (string)** - đại diện cho dữ liệu dạng văn bản.

#### ⚠️ Quy tắc sử dụng dấu nháy trong Go:
* **Bắt buộc:** Phải sử dụng **dấu nháy kép `" "`** để khai báo chuỗi (ví dụ: `"Hello World"`).
* **Khác biệt ngôn ngữ:** Trong JavaScript hay PHP, bạn có thể dùng dấu nháy đơn `' '` để viết chuỗi. Tuy nhiên, trong Go điều này **bị cấm hoàn toàn** (dấu nháy đơn chỉ dành cho ký tự đơn - Rune).
* **Dấu backtick (`` ` ``):** Ngoài nháy kép, bạn cũng có thể dùng dấu backtick để khai báo chuỗi nhiều dòng (sẽ học chi tiết ở bài sau). Nhưng nháy kép vẫn là lựa chọn phổ biến nhất.

---

## 🎬 Hàm `main` - Trái Tim Của Chương Trình

Hàm `main` cũng là một hàm giống như `Print`, nhưng điểm khác biệt quan trọng:
* **Hàm `Print`:** Là hàm được viết sẵn trong thư viện mà chúng ta **gọi ra để dùng**.
* **Hàm `main`:** Là hàm do **chính chúng ta tự định nghĩa**.
* **Cơ chế:** Khi chương trình Go chạy, Go Engine sẽ tự động tìm đến hàm `main` này để kích hoạt chạy toàn bộ mã nguồn bên trong nó.

---

## 📝 Tóm Tắt Bài Học
| Thành phần | Vai trò chính |
| :--- | :--- |
| **`func main()`** | Hàm chính khởi chạy ứng dụng do lập trình viên định nghĩa. |
| **`fmt.Print()`** | Hàm hiển thị dữ liệu ra Terminal có sẵn từ thư viện chuẩn. |
| **`"Hello World"`** | Giá trị dạng chuỗi văn bản (bắt buộc dùng dấu nháy kép). |
