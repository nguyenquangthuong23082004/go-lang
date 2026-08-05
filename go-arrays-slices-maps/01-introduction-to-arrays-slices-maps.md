# Bài 01: Giới Thiệu Về Arrays, Slices Và Maps Trong Go

> [!NOTE]
> *Chúng ta đã tìm hiểu về kiểu dữ liệu **Struct** – một công cụ mạnh mẽ để gom nhóm các mẩu dữ liệu có kiểu khác nhau thành một giá trị duy nhất. Trong chương này, chúng ta sẽ khám phá các kiểu dữ liệu tích hợp sẵn (built-in) khác trong Go cũng có nhiệm vụ tổ chức và gom nhóm dữ liệu liên quan, bao gồm: **Arrays (Mảng)**, **Slices (Lát cắt)** và **Maps (Bảng băm/Ánh xạ)**. Chúng ta cũng sẽ so sánh khi nào nên dùng các kiểu dữ liệu này và khi nào nên ưu tiên Struct.*

---

## 🔍 1. Gom Nhóm Dữ Liệu Trong Go: Struct vs. Collections

Trước khi đi sâu vào chi tiết, hãy cùng điểm qua sự khác biệt cốt lõi về mục đích sử dụng:
* **Struct**: Dùng để gom nhóm các dữ liệu **khác kiểu** (như tên kiểu `string`, tuổi kiểu `int`) đại diện cho một đối tượng cụ thể (ví dụ: một User, một Product).
* **Arrays, Slices, Maps (Collections)**: Dùng để gom nhóm danh sách các phần tử **cùng kiểu dữ liệu** (ví dụ: danh sách các số nguyên, danh sách các chuỗi, danh sách các Structs).

---

## 🎯 2. Các Nội Dung Sẽ Học Trong Phần Này

Chúng ta sẽ lần lượt đi qua các khái niệm quan trọng để làm chủ cách quản lý danh sách dữ liệu trong Go:

### A. Arrays (Mảng Tĩnh)
* Định nghĩa mảng và cách khai báo.
* Các giới hạn của Array trong Go (kích thước cố định, không thể thay đổi sau khi khai báo).
* Khi nào nên dùng Array.

### B. Slices (Mảng Động/Lát Cắt)
* Slice là gì và nó giải quyết các hạn chế của Array như thế nào.
* Cách hoạt động bên dưới của Slice (con trỏ trỏ tới Array gốc, độ dài `len`, sức chứa `cap`).
* Deep dive (Đi sâu) vào cách làm việc với Slices: các hàm `append`, cắt lát mảng (slicing), sao chép mảng, v.v.

### C. Maps (Bảng băm / Ánh xạ Key-Value)
* Định nghĩa Map trong Go và cách tổ chức dữ liệu theo cặp Khóa - Giá trị (Key-Value).
* Cách thêm, sửa, xóa và truy xuất phần tử trong Map.
* Các điểm lưu ý đặc biệt về Map (kiểu tham chiếu, không có thứ tự cố định).

### D. So Sánh & Lựa Chọn Kiểu Dữ Liệu Phù Hợp
* Khi nào chọn Struct?
* Khi nào chọn Array/Slice?
* Khi nào chọn Map?

---

## 💡 3. Tại Sao Việc Tổ Chức Dữ Liệu Lại Quan Trọng?
Trong lập trình thực tế, chúng ta hiếm khi làm việc với các biến đơn lẻ. Việc quản lý một danh sách bài viết, danh sách đơn hàng hay ánh xạ giữa ID người dùng và thông tin chi tiết đòi hỏi các cấu trúc dữ liệu linh hoạt, tối ưu về hiệu năng và an toàn bộ nhớ. Go cung cấp các kiểu dữ liệu tích hợp này với hiệu năng rất cao, giúp chúng ta xử lý các tập dữ liệu lớn một cách dễ dàng.
