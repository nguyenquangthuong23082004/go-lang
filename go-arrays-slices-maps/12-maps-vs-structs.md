# Bài 12: So Sánh Chi Tiết: Maps vs Structs (Khi Nào Dùng Gì?)

> [!NOTE]
> *Cả **Map** và **Struct** đều lưu trữ dữ liệu dưới dạng Khóa - Giá trị (Key-Value). Vậy sự khác biệt cốt lõi là gì và khi nào nên chọn Map, khi nào nên chọn Struct?*

---

## ⚖️ 1. Bảng So Sánh Tổng Quan (Maps vs Structs)

| Tiêu chí | Map 🗺️ | Struct 🏢 |
| :--- | :--- | :--- |
| **Bản chất** | Là một **tập hợp linh hoạt (collection)** của nhiều giá trị cùng/khác loại. Có thể coi Map là **Array dùng nhãn (labels) thay cho số (index)**. | Khai báo một **thực thể dữ liệu (Data Entity)** có cấu trúc cố định và định hình sẵn. |
| **Tính co giãn (Độ linh hoạt)** | **Rất động:** Có thể thêm/xóa Key-Value thoải mái lúc runtime. | **Cố định:** Các trường (fields) được chốt hạ khi định nghĩa. Không thể thêm/xóa trường lúc chạy code. |
| **Tên Key** | Nhãn tùy ý (chuỗi có khoảng trắng như `"Amazon Web Services"`, số, struct,...). | Tên định danh biến chuẩn trong Go (không chứa khoảng trắng, ví dụ: `Price`, `Title`). |
| **Kiểu dữ liệu của Key & Value** | Tất cả các Key phải cùng kiểu, tất cả các Value phải cùng kiểu (VD: `map[string]string`). | Các trường trong Struct có thể mang các **kiểu dữ liệu hoàn toàn khác nhau** (`string`, `float64`, `int`, `bool`,...). |

---

## 🎯 2. Khi Nào Nên Dùng Struct?

**Dùng Struct khi bạn cần mô tả một Đối tượng / Thực thể dữ liệu cụ thể (Data Entity):**
- Có các thuộc tính cố định, biết trước khi viết code.
- Mỗi thuộc tính có kiểu dữ liệu khác nhau.
- Ví dụ: `Product` (id `string`, price `float64`), `User` (name `string`, age `int`), `Order`, `Invoice`.

```go
type Product struct {
    ID    string
    Title string
    Price float64
}
// Không thể tự dưng thêm product.DiscountCode lúc runtime nếu Struct chưa định nghĩa!
```

---

## 🗺️ 3. Khi Nào Nên Dùng Map?

**Dùng Map khi bạn cần quản lý một Tập Hợp (Collection) các dữ liệu linh hoạt:**
- Cần mở rộng hoặc thu nhỏ danh sách liên tục lúc runtime (thêm/xóa key).
- Cần nhãn đặt tên tùy biến (Key) để tìm kiếm thay vì dùng chỉ mục số `0, 1, 2...` của Array.
- Các dữ liệu có cùng bản chất/kiểu dữ liệu (ví dụ: Danh sách URL của các công ty, bảng điểm của các học sinh).

```go
websites := map[string]string{
    "Google": "https://google.com",
    "Amazon Web Services": "https://aws.com",
}
// Dễ dàng thêm 100 công ty mới lúc chạy chương trình!
websites["Facebook"] = "https://facebook.com"
```

---

## 💡 Tóm Tắt Trong 1 Câu Nhanh

> - **Struct**: Định hình một **đối tượng cụ thể** với cấu trúc cố định (Giống như một Form mẫu).
> - **Map**: Danh sách **mảng động dùng nhãn** tự do thay cho chỉ mục số (Giống như Từ điển).
