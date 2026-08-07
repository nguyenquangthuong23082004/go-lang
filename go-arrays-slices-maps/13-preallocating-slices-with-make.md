# Bài 13: Tối Ưu Hóa Bộ Nhớ Của Slice Bằng Hàm `make()`

> [!NOTE]
> *Học cách khởi tạo Slice bằng hàm built-in `make()` để cấp phát trước dung lượng bộ nhớ (`len` và `cap`), giúp tối ưu hóa hiệu năng chương trình bằng cách tránh việc Go phải cấp phát lại mảng nhiều lần.*

---

## ⚡ 1. Vấn Đề Về Hiệu Năng Khi Dùng Slice Rỗng Trực Tiếp

Khi khởi tạo một Slice rỗng thông thường và liên tục `append()` phần tử mới:
```go
userNames := []string{}
userNames = append(userNames, "Max")
userNames = append(userNames, "Manuel")
```

**Cơ chế đằng sau:**
- Ban đầu `userNames` có `cap = 0`.
- Mỗi lần `append()` vượt quá `cap`, Go phải tạm dừng, **cấp phát một vùng nhớ mới**, copy toàn bộ dữ liệu từ mảng cũ sang mảng mới, rồi giải phóng mảng cũ.
- Nếu bạn cần thêm hàng nghìn/hàng triệu phần tử, việc tạo mảng liên tục này sẽ làm giảm hiệu năng ứng dụng.

---

## 🛠️ 2. Tối Ưu Bộ Nhớ Với Hàm `make()`

Nếu bạn ước lượng được số lượng phần tử sắp thêm vào, hàm **`make()`** cho phép bạn yêu cầu Go cấp phát trước một mảng có **sức chứa (`cap`)** tương ứng ngay từ đầu.

### Cú pháp:
```go
slice := make([]Kiểu_Dữ_Liệu, len, cap)
```
- **`len` (Length):** Số phần tử ban đầu (các ô này chứa giá trị mặc định/zero value).
- **`cap` (Capacity - Thường chọn tùy chọn):** Sức chứa tối đa được cấp phát bộ nhớ sẵn ở đằng sau hậu trường.

---

## 🧪 3. Các Trường Hợp Sử Dụng `make()`

### Trường hợp A: Truy cập theo chỉ mục (`len > 0`)
```go
// Tạo slice string có 2 phần tử rỗng ban đầu ""
userNames := make([]string, 2, 5)

// Gán giá trị trực tiếp cho 2 ô ban đầu theo chỉ mục (Index)
userNames[0] = "Max"
userNames[1] = "Manuel"

fmt.Println(userNames) // Output: [Max Manuel]
```

> [!CAUTION]
> Nếu bạn khởi tạo `make([]string, 2, 5)` mà lại dùng `append(userNames, "Max")`, giá trị `"Max"` sẽ bị đẩy vào vị trí **thứ 3** (sau 2 phần tử rỗng ban đầu) -> Kết quả ra `["", "", "Max"]`.

### Trường hợp B: Cấp phát trước bộ nhớ để `append()` dần (`len = 0`)
Nếu bạn muốn dùng `append()` mà không có ô rỗng thừa ở đầu, hãy đặt `len = 0`:

```go
// len = 0 (chưa có phần tử nào), cap = 5 (đã đặt trước 5 chỗ trong RAM)
userNames := make([]string, 0, 5)

// Tối ưu! append() 5 phần tử đầu tiên sẽ KHÔNG bị cấp phát lại bộ nhớ
userNames = append(userNames, "Max")
userNames = append(userNames, "Manuel")

fmt.Println(userNames) // Output: [Max Manuel]
```

---

## 📝 Tóm Tắt Khi Nào Dùng `make()`

1. **Ứng dụng bình thường / Dữ liệu nhỏ:** Khai báo Slice literals (`[]string{...}`) hoặc slice rỗng là đủ, không cần tối ưu quá sớm.
2. **Xử lý hiệu năng / Biết trước số lượng:** Dùng `make([]T, 0, cap)` khi nạp dữ liệu lớn (như đọc file, query database 10.000 dòng) để tăng tốc độ thực thi chương trình.
