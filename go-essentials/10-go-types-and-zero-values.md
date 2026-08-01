# Bài 9: Các Kiểu Dữ Liệu Cơ Bản & Giá Trị Mặc Định (Zero Values) trong Go

> [!NOTE]
> *Go là ngôn ngữ kiểm soát kiểu dữ liệu tĩnh. Điều này nghĩa là mọi biến đều phải có một kiểu dữ liệu xác định và nếu bạn không gán giá trị ban đầu cho biến, Go sẽ tự động gán cho nó một giá trị mặc định đặc trưng.*

---

## 🎨 1. Các Kiểu Dữ Liệu Cơ Bản (Basic Types)

Dưới đây là các kiểu dữ liệu phổ biến nhất mà bạn sẽ sử dụng hàng ngày:

* **`int`:** Số nguyên không có phần thập phân (ví dụ: `-10`, `0`, `45`).
* **`float64`:** Số thực có dấu phẩy thập phân (ví dụ: `1.5`, `-9.85`, `100.0`).
* **`string`:** Chuỗi ký tự (văn bản). Được tạo bằng dấu nháy kép `"Hello World"` hoặc dấu backticks (phím huyền) `` `Hi everyone` ``.
* **`bool`:** Kiểu Boolean chỉ nhận hai giá trị `true` (đúng) hoặc `false` (sai).

---

## 🔍 2. Các Kiểu Dữ Liệu "Ngách" (Niche Basic Types)

Đây là các kiểu dữ liệu chuyên sâu hơn, giúp tối ưu hóa bộ nhớ hoặc giải quyết các bài toán đặc thù:

* **`uint`:** Số nguyên không dấu (Unsigned Integer), chỉ nhận giá trị lớn hơn hoặc bằng 0 (ví dụ: `0`, `100`).
* **`int32`:** Số nguyên có dấu 32-bit, giới hạn phạm vi lưu trữ từ $-2.147.483.648$ đến $2.147.483.647$.
* **`rune`:** Là tên gọi khác (alias) của kiểu `int32`. Nó đại diện cho một ký tự Unicode đơn lẻ và được viết trong dấu nháy đơn (Ví dụ: `'a'`, `'ñ'`, `'世'`).
* **`uint32`:** Số nguyên không dấu 32-bit, nhận giá trị từ $0$ đến $4.294.967.295$.
* **`int64`:** Số nguyên có dấu 64-bit, lưu trữ các số siêu lớn (ví dụ ID hệ thống, timestamp nano giây).
* *Ngoài ra còn có các kiểu như `int8`, `uint8`, `int16`, `uint16` tương tự với phạm vi nhỏ hơn.*

---

## 🌀 3. Giá Trị Mặc Định (Zero Values / Null Values)

Không giống như một số ngôn ngữ khác nơi biến chưa khởi tạo sẽ mang giá trị `undefined` hoặc `null` dẫn đến lỗi Runtime, Go tự động gán giá trị mặc định (**Zero Value**) tương ứng với kiểu của biến đó.

Ví dụ:
```go
var age int // age sẽ tự động mang giá trị 0
```

### Bảng giá trị mặc định (Zero Values):

| Kiểu dữ liệu | Giá trị mặc định (Zero Value) |
| :--- | :--- |
| **`int` / `int32` / `int64` / `uint`** | `0` |
| **`float64` / `float32`** | `0.0` |
| **`string`** | `""` *(chuỗi rỗng)* |
| **`bool`** | `false` |

---

## 📚 4. Danh Sách Bài Học Tiếp Theo Trong Phần Cốt Lõi (Go Essentials)
* Bài 10: In dữ liệu ra màn hình (Outputting Values)
* Bài 11: Chuyển đổi kiểu & gán kiểu dữ liệu hiển thị (Explicit Type Assignment)
* Bài 12: Các cách khai báo biến khác (Alternative Variable Declarations)
* Bài 13: Làm việc với hằng số (Constants)
* Bài 14: Nhận dữ liệu nhập từ người dùng bằng `fmt.Scan()`
* Bài 15: Định dạng chuỗi văn bản nâng cao (Formatting Strings)
* Bài 16: Đi sâu tìm hiểu về Hàm (Functions)
* Bài 17: Cấu trúc điều khiển (if-else, switch, for-loop)
* Bài 18: Đọc/Ghi file & Xử lý lỗi (File I/O & Error Handling)
