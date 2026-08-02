# Bài 07: Sử Dụng Con Trỏ Làm Tham Số Cho Hàm (Pointers & Functions)

> [!IMPORTANT]
> *Để truyền địa chỉ ô nhớ vào một hàm, chúng ta cấu hình tham số đầu vào của hàm là một kiểu con trỏ. Trong Go, mọi phép toán số học trực tiếp trên địa chỉ con trỏ đều bị cấm.*

---

## 🚫 1. Không Có Phép Toán Số Học Trên Con Trỏ (No Pointer Arithmetic)

Trong một số ngôn ngữ như C/C++, lập trình viên có thể cộng/trừ trực tiếp trên địa chỉ con trỏ (ví dụ: `ptr + 1` để nhảy sang ô nhớ tiếp theo). Tuy nhiên, việc này rất dễ gây ra các lỗi bảo mật nghiêm trọng nếu truy cập sai phân vùng bộ nhớ.

Go **ngăn chặn hoàn toàn** điều này:
* Nếu `age` là một con trỏ kiểu `*int`, biểu thức `age - 18` là **không hợp lệ** và sẽ bị báo lỗi biên dịch lập tức.
* **Giải pháp**: Phải giải tham chiếu để lấy giá trị số thực tế trước, sau đó mới tính toán: `*age - 18`.

---

## 💻 2. Định Nghĩa Hàm Nhận Tham Số Con Trỏ

Dưới đây là hàm `getAdultYears` sau khi được nâng cấp để nhận con trỏ `*int`:

```go
func getAdultYears(age *int) int {
	return *age - 18 // Giải tham chiếu lấy giá trị rồi trừ đi 18
}
```

---

## 🔌 3. Cách Gọi Hàm Với Tham Số Con Trỏ

Trong hàm `main()`, chúng ta có hai cách để truyền đối số con trỏ vào hàm:

### Cách 1: Truyền biến con trỏ đã khai báo
```go
agePointer := &age
adultYears := getAdultYears(agePointer) // Truyền trực tiếp con trỏ
```

### Cách 2: Truyền địa chỉ bằng toán tử `&` trực tiếp
```go
adultYears := getAdultYears(&age) // Lấy địa chỉ và truyền ngay lập tức
```

---

## ⚠️ 4. Khuyến Nghị Về Hiệu Năng (Performance Guidelines)

Tránh việc lạm dụng con trỏ đối với các kiểu dữ liệu nguyên thủy cơ bản (như `int`, `float64`, `bool`, `string` ngắn):
* Việc sao chép một số nguyên `32` trong bộ nhớ RAM diễn ra cực kỳ nhanh và tốn rất ít tài nguyên.
* Nếu cố gắng dùng con trỏ cho số nguyên, bạn sẽ tốn thêm bộ nhớ để lưu biến con trỏ và tốn thời gian để CPU thực hiện giải tham chiếu (indirection overhead).
* **Lời khuyên**: Chỉ dùng con trỏ khi dữ liệu truyền vào là một cấu trúc dữ liệu lớn (như Struct nhiều trường) hoặc khi bạn thực sự muốn hàm chỉnh sửa giá trị biến gốc.
