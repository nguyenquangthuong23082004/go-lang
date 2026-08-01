# Bài 8: Kiểu Dữ Liệu & Ép Kiểu (Type Conversion) Trong Go

> [!NOTE]
> *Go là ngôn ngữ kiểu tĩnh cực kỳ nghiêm ngặt. Bài học này giải thích chi tiết cơ chế hoạt động của kiểu dữ liệu, cách thực hiện ép kiểu và cách sử dụng thư viện toán học để tính lũy thừa.*

---

## 📊 1. Kiểu Số Nguyên (`int`) vs Kiểu Số Thực (`float64`)

Mọi giá trị trong chương trình Go đều có một kiểu dữ liệu cụ thể. Khi di chuột qua các biến trong VS Code, bạn sẽ thấy kiểu dữ liệu được tự động suy luận:

* **Kiểu `int` (Integer):** Dành cho số nguyên không có phần thập phân (Ví dụ: `1000`, `10`, `-5`).
* **Kiểu `float64`:** Dành cho số thực có dấu phẩy thập phân (Ví dụ: `5.5`, `3.14`, `1000.0`). *Số 64 nghĩa là nó sử dụng 64-bit bộ nhớ để lưu trữ độ chính xác cao.*

### 🚫 Quy tắc bắt buộc:
Go **không tự động ép kiểu ẩn** (implicit conversion) khi tính toán. Bạn không thể thực hiện phép toán (nhân, chia, cộng, trừ) trực tiếp giữa một số `int` và một số `float64`. 

---

## 🔄 2. Chuyển Đổi Kiểu Dữ Liệu Chủ Động (Type Conversion)

Để tính toán giữa `int` và `float64`, bạn phải chủ động ép kiểu (convert) giá trị về cùng một dạng.

* **Cú pháp:** Sử dụng tên kiểu dữ liệu mục tiêu như một hàm bọc ngoài giá trị cần đổi:
  ```go
  float64(investmentAmount)
  ```
* **Cơ chế:** Lệnh này chuyển đổi giá trị số nguyên thành số thực (ví dụ `1000` thành `1000.0`) trong phạm vi phép tính đó. Biến gốc vẫn giữ nguyên kiểu dữ liệu ban đầu trong bộ nhớ.

---

## 🧮 3. Tính Lũy Thừa Bằng Package `math`

Trong Go, **không có toán tử lũy thừa trực tiếp** như ký tự `^` hay `**` ở một số ngôn ngữ khác. Thay vào đó, bạn phải sử dụng hàm `Pow` từ gói thư viện chuẩn **`math`**.

### 🔌 Bước 1: Import thư viện `math` ở đầu file
```go
import "math"
```

### 🛠️ Bước 2: Sử dụng hàm `math.Pow(x, y)`
Hàm `math.Pow` nhận vào 2 tham số và yêu cầu cả hai đều phải thuộc kiểu **`float64`**:
* Tham số thứ nhất: Cơ số ($x$)
* Tham số thứ hai: Số mũ ($y$)

Vì biến số năm (`years`) được khai báo ban đầu là kiểu `int`, ta phải ép kiểu nó sang `float64` khi truyền vào hàm:
```go
math.Pow(1 + expectedReturnRate / 100, float64(years))
```

---

## 📈 4. Công Thức Tính Toàn Diện & Hợp Lệ

Dưới đây là dòng code tính toán hoàn chỉnh không còn lỗi biên dịch:

```go
var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
```

* **Cơ chế suy luận kiểu:** Vì kết quả của biểu thức bên phải là `float64`, Go sẽ tự động thiết lập kiểu dữ liệu cho biến `futureValue` là **`float64`**.
