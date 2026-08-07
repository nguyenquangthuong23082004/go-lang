# Phần 09: Hàm Nâng Cao Trong Go (Functions Deep Dive)

> [!NOTE]
> *Chào mừng bạn đến với chương học chuyên sâu về **Hàm (Functions)** trong Go! Trong chương này, chúng ta sẽ khám phá các tính năng nâng cao mạnh mẽ liên quan đến Hàm.*

---

## 🎯 1. Các Chủ Đề Chính Sẽ Học Trong Chương

1. **Functions as Values (Hàm là một Giá trị / First-Class Functions):**
   - Truyền hàm làm đối số (argument) cho một hàm khác (Higher-Order Functions).
   - Lưu trữ hàm vào biến, Slice, hoặc Map.
   - Trả về một hàm từ một hàm khác.
   - Định nghĩa Custom Function Types.

2. **Anonymous Functions & Closures (Hàm Ẩn Danh & Closure):**
   - Tạo các hàm không có tên trực tiếp trong dòng lệnh (Inline Functions).
   - Khái niệm Closure: Hàm ẩn danh truy cập và "đóng đóng" (enclose) các biến ở phạm vi bên ngoài.

3. **Recursion (Đệ Quy):**
   - Kỹ thuật hàm tự gọi lại chính nó để giải quyết các bài toán chia để trị (giai thừa, cây dữ liệu,...).
   - Cơ chế hoạt động của Call Stack và điều kiện dừng (Base Case).

4. **Variadic Functions (Hàm Với Số Lượng Đối Số Bất Kỳ):**
   - Định nghĩa hàm chấp nhận số lượng tham số linh hoạt (`func sum(numbers ...int)`).
   - Truyền Slice vào Variadic Function bằng toán tử `...`.

---

## 📁 2. Cấu Trúc Thư Mục Dự Án

- Mã nguồn thực hành: [learning/09-functions-deep-dive/main.go](file:///home/uwal/Desktop/go-lang/learning/09-functions-deep-dive/main.go)
- Tài liệu lý thuyết chi tiết: Thư mục [go-functions-deep-dive](file:///home/uwal/Desktop/go-lang/go-functions-deep-dive/) *(sẽ cập nhật song song theo các bài học)*.
