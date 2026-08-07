# Bài 11: Thao Tác Với Map (Đọc, Thêm, Sửa, Xóa Dữ Liệu)

> [!NOTE]
> *Học các thao tác cơ bản làm việc với Map trong Go: Truy xuất dữ liệu theo Key, Thêm Key-Value mới (Map luôn là mảng động), Cập nhật dữ liệu, và Xóa phần tử khỏi Map bằng hàm `delete()`.*

---

## 📖 1. Đọc Giá Trị Theo Key (Reading Values)

Không dùng chỉ mục số `0, 1, 2...` như Array/Slice, bạn lấy giá trị từ Map bằng cách truyền **Key** vào trong cặp ngoặc vuông `[]`:

```go
url := websites["Amazon Web Services"]
fmt.Println(url) // Output: https://aws.com
```

> [!WARNING]
> Phải đảm bảo viết chính xác Tên Key (gồm cả hoa, thường và khoảng trắng). Viết sai tên Key sẽ dẫn đến lấy ra giá trị rỗng (Zero value của kiểu dữ liệu đó).

---

## ➕ 2. Thêm Hoặc Cập Nhật Dữ Liệu Trong Map

Khác với Array (độ dài cố định), **Map trong Go mặc định luôn luôn động (dynamic)**. Bạn có thể thêm hoặc sửa dữ liệu cực kỳ đơn giản:

### A. Thêm cặp Key-Value mới:
Chỉ cần gán giá trị cho một Key chưa từng tồn tại:
```go
websites["LinkedIn"] = "https://linkedin.com"
```

### B. Cập nhật (Sửa) giá trị của Key đã có:
Gán giá trị mới cho một Key đã tồn tại:
```go
websites["Google"] = "https://google.com.vn" // Gán đè URL mới
```

---

## 🗑️ 3. Xóa Cặp Key-Value Khỏi Map Với Hàm `delete()`

Go cung cấp hàm tích hợp sẵn **`delete()`** để xóa một phần tử ra khỏi Map.

### Cú pháp:
```go
delete(tên_map, "tên_key_cần_xóa")
```

### Ví dụ:
```go
// Xóa cặp key "Google" khỏi map websites
delete(websites, "Google")

fmt.Println(websites) 
// Google đã bị xóa hoàn toàn khỏi Map
```

---

## 📝 Tóm Tắt Thao Tác Với Map

```go
// 1. Tạo Map
m := map[string]string{"A": "Apple"}

// 2. Đọc (Read)
val := m["A"]

// 3. Thêm/Sửa (Insert/Update)
m["B"] = "Banana" 

// 4. Xóa (Delete)
delete(m, "A")
```
