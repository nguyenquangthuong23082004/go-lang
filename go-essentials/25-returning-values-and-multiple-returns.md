# Bài 24: Giá Trị Trả Về (Return Values) & Trả Về Nhiều Giá Trị Trong Go

> [!NOTE]
> *Để tối ưu hóa mã nguồn, chúng ta có thể chuyển logic tính toán phức tạp ra một hàm riêng. Go hỗ trợ tính năng trả về giá trị (Return) và đặc biệt là khả năng trả về nhiều giá trị cùng một lúc.*

---

## 🌎 1. Phạm Vi Hoạt Động Của Biến (Variable Scope)

Khi tổ chức code thành nhiều hàm, bạn cần hiểu rõ biến đó thuộc phạm vi nào:

* **Phạm vi cục bộ (Local Scope):** Biến/hằng số được khai báo bên trong một hàm chỉ tồn tại và sử dụng được trong phạm vi hàm đó.
* **Phạm vi toàn cục (Package-level Scope):** Biến/hằng số khai báo bên ngoài mọi hàm (ở cấp độ file) có thể được truy cập bởi bất kỳ hàm nào trong cùng file đó.
  * *Ví dụ:* Hằng số `tyLeLamPhat` không thay đổi trong suốt chương trình, nên ta có thể đưa lên phạm vi toàn cục để tất cả các hàm cùng sử dụng.
  * *Lưu ý:* Không nên đưa các biến nhập liệu động (`khoanDauTu`, `years`) lên toàn cục vì chúng sẽ mang giá trị mặc định (`0.0`), dễ gây lỗi tính toán sai thời điểm. Thay vào đó, hãy giữ chúng ở hàm `main()` và truyền dưới dạng tham số đầu vào.

---

## 🛠️ 2. Khai Báo Hàm Trả Về Giá Trị (Return Values)

Vì Go là ngôn ngữ kiểm soát kiểu tĩnh (statically typed), bạn phải khai báo rõ kiểu dữ liệu mà hàm sẽ trả về trước khi mở ngoặc nhọn `{}`.

### A. Hàm trả về một giá trị duy nhất:
Ghi kiểu dữ liệu trả về ngay sau ngoặc đơn tham số:
```go
func calculateFutureValue(khoanDauTu, laiSuatHangNam, soNamDauTu float64) float64 {
	fv := khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	return fv
}
```

### B. Hàm trả về nhiều giá trị (Multiple Returns):
Bọc các kiểu dữ liệu trả về trong cặp ngoặc đơn `()` và phân tách bằng dấu phẩy `,`:
```go
func calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu float64) (float64, float64) {
	fv := khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	rfv := fv / math.Pow(1+tyLeLamPhat/100, soNamDauTu)
	return fv, rfv // Trả về đồng thời hai giá trị
}
```

---

## 🏃 3. Gọi Hàm Và Hứng Nhiều Giá Trị Trả Về

Khi gọi hàm trả về nhiều giá trị, ở vế trái toán tử `:=` bạn cần chuẩn bị các biến hứng tương ứng ngăn cách bởi dấu phẩy:

```go
giaTriTuongLai, giaTriThucTe := calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu)
```

* Giá trị trả về thứ nhất (`fv`) được gán cho biến hứng thứ nhất (`giaTriTuongLai`).
* Giá trị trả về thứ hai (`rfv`) được gán cho biến hứng thứ hai (`giaTriThucTe`).

Tính năng này giúp bạn viết code ngắn gọn hơn rất nhiều mà không cần tạo cấu trúc phức tạp hay viết nhiều hàm nhỏ lẻ chỉ để lấy một giá trị.
