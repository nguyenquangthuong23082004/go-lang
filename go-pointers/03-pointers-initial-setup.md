# Bài 03: Khởi Tạo Dự Án & Ví Dụ Truyền Bản Sao (Pass by Value)

> [!NOTE]
> *Để chuẩn bị thực hành con trỏ, chúng ta sẽ bắt đầu bằng cách dựng cấu trúc dự án mới và viết mã nguồn truyền biến theo giá trị (bản sao) thông thường để làm hệ quy chiếu so sánh.*

---

## 📂 1. Cấu Trúc Thư Mục Dự Án Mới

Tạo một thư mục dự án độc lập mới tại `/learning/05-pointers/` với hai tệp tin:
1. **`go.mod`**: Khai báo tên module của dự án:
   ```go
   module example.com/pointers
   
   go 1.21
   ```
2. **`pointers.go`**: Triển khai mã nguồn cơ bản sử dụng package `main`.

---

## 💻 2. Mã Nguồn Ban Đầu (`pointers.go`)

Mã nguồn truyền thống chưa sử dụng con trỏ:
```go
package main

import "fmt"

func main() {
	age := 32 // Biến thông thường

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
```

---

## 🔍 3. Cơ Chế Sao Chép Bộ Nhớ (Memory Copy) Trong Đoạn Code Trên

1. **Khởi tạo**: Dòng `age := 32` lưu số `32` tại một địa chỉ ô nhớ (ví dụ: `0xAA`).
2. **Gọi hàm**: Khi gọi `getAdultYears(age)`, Go không truyền ô nhớ `0xAA` vào hàm. Thay vào đó, Go tạo một bản sao của số `32` và lưu tại ô nhớ mới (ví dụ: `0xBB`) rồi đưa cho hàm `getAdultYears` xử lý.
3. **Giải phóng**: Khi hàm `getAdultYears` thực thi xong và trả về kết quả, ô nhớ sao chép `0xBB` không còn được sử dụng nữa. Bộ dọn rác (Garbage Collector) của Go sẽ quét và giải phóng ô nhớ này trong tương lai gần.
