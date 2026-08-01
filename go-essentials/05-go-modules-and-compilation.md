# Bài 4: Khái Niệm Go Modules & Biên Dịch Ứng Dụng Thực Tế

> [!NOTE]
> *Trong bài này, chúng ta sẽ phân biệt hai khái niệm quan trọng: Package và Module, đồng thời học cách khởi tạo một Go Module để biên dịch thành công ứng dụng nhị phân.*

---

## 🆚 1. Phân Biệt: Package vs Module trong Go

Để giải quyết triệt để lỗi `missing main module`, chúng ta cần hiểu rõ sự khác biệt giữa hai khái niệm này:

| Khái niệm | Quy mô | Vai trò & Đặc điểm |
| :--- | :--- | :--- |
| **Package (Gói)** | Nhỏ hơn | Dùng để gom nhóm các file code có cùng chức năng (ví dụ: package `main` để khởi chạy, package `fmt` để in dữ liệu). |
| **Module (Mô-đun)** | Lớn hơn | Đại diện cho **cả dự án**. Một Module có thể chứa **nhiều Package khác nhau**. Một dự án Go thông thường sẽ tương đương với một Module. |

---

## 🛠️ 2. Khởi Tạo Go Module Bằng Lệnh `go mod init`

Để Go nhận diện dự án của bạn là một module và cho phép build, bạn bắt buộc phải khởi tạo nó thông qua terminal:

```bash
go mod init <tên_module_hoặc_đường_dẫn>
```

### 💡 Quy tắc đặt tên Module:
* **Khi phát hành thực tế:** Thường dùng đường dẫn kho lưu trữ mã nguồn của bạn để người khác có thể dễ dàng tải về (Ví dụ: `github.com/username/project-name`).
* **Khi thực hành/học tập:** Bạn có thể dùng bất kỳ tên giả lập nào, ví dụ:
  ```bash
  go mod init example.com/first-app
  ```

### 📄 File `go.mod` là gì?
Sau khi chạy lệnh trên, Go sẽ tự động tạo ra một file tên là **`go.mod`** trong thư mục dự án.
* Sự xuất hiện của file `go.mod` báo cho trình biên dịch Go biết rằng thư mục hiện tại cùng tất cả các thư mục con bên trong chính thức thuộc về Module có tên được khai báo.
* File này lưu trữ thông tin về phiên bản Go đang dùng cũng như các thư viện bên thứ ba (dependencies) mà dự án phụ thuộc vào.

---

## 🚀 3. Biên Dịch & Chạy Chương Trình Độc Lập

Khi đã có file `go.mod`, bạn chạy lại lệnh build:
```bash
go build
```
Lệnh sẽ chạy thành công mà không có lỗi. Kết quả thu được là một file thực thi nhị phân duy nhất mang tên trùng với tên module (ví dụ: `first-app`).

### 💻 Hướng dẫn chạy file thực thi sau khi build:
* **Trên Windows:** Bạn sẽ nhận được file `first-app.exe`. Bạn có thể kích đúp chuột vào file để chạy trực tiếp.
* **Trên Linux / macOS:** Bạn nhận được file chạy nhị phân không có đuôi. Hãy mở terminal tại thư mục chứa file và chạy:
  ```bash
  ./first-app
  ```
  *Kết quả hiển thị:* `Hello World` (giống hệt lệnh `go run`, nhưng có thể chạy độc lập trên máy không cài Go SDK).

---

## ⚠️ 4. Thử Nghiệm: Đổi Tên Package `main`

Nếu bạn mở file code chính và đổi dòng khai báo đầu tiên từ `package main` thành `package app`, sau đó lưu lại và chạy `go build`:
* **Kết quả:** Go build vẫn chạy qua nhưng **không có file thực thi nào được tạo ra**.
* **Giải thích:** Go Compiler không tìm thấy điểm bắt đầu chính (`package main` và hàm `func main()`). Nó hiểu package `app` chỉ là một thư viện hỗ trợ nên nó không đóng gói thành chương trình chạy độc lập.

> [!IMPORTANT]
> **Quy tắc vàng:** Tên package **`main`** là tên gói được giữ riêng (reserved package name) để làm điểm bắt đầu thực thi chương trình của Go. Bắt buộc phải có `package main` để tạo ra file thực thi chạy độc lập.
