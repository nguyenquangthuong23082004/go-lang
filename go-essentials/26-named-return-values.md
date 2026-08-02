# Bài 25: Đặt Tên Cho Giá Trị Trả Về (Named Return Values) trong Go

> [!NOTE]
> *Go cho phép bạn đặt tên trực tiếp cho các biến trả về ngay tại phần khai báo hàm (chữ ký hàm). Tính năng này mang lại một số đặc trưng thú vị và các quy ước viết code cần lưu ý.*

---

## 🛠️ 1. Khai Báo Named Return Values

Thay vì chỉ viết kiểu dữ liệu trả về là `(float64, float64)`, bạn có thể đặt tên biến cụ thể kèm kiểu dữ liệu cho chúng:

```go
func calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu float64) (fv float64, rfv float64) {
	fv = khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	rfv = fv / math.Pow(1+tyLeLamPhat/100, soNamDauTu)
	return fv, rfv
}
```

### ⚙️ Các đặc điểm cần lưu ý:
1. **Tự động khai báo:** Go sẽ tự động tạo sẵn các biến `fv` và `rfv` ở dạng mặc định (Zero Value) ngay khi bắt đầu thực thi hàm.
2. **Không dùng `:=`:** Trong thân hàm, bạn không sử dụng toán tử khai báo và gán nhanh `:=` nữa. Thay vào đó, bạn sử dụng toán tử gán thông thường `=` vì các biến này đã được khai báo ở đầu hàm rồi.

---

## 👻 2. Khái Niệm Naked Return (Trả Về Trống)

Khi hàm đã có tên biến trả về rõ ràng ở phần khai báo, Go cho phép bạn rút gọn câu lệnh trả về thành một chữ `return` duy nhất:

```go
func calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu float64) (fv float64, rfv float64) {
	fv = khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	rfv = fv / math.Pow(1+tyLeLamPhat/100, soNamDauTu)
	return // Naked Return - Go tự động hiểu là trả về fv và rfv
}
```

---

## 🚨 Lời Khuyên Thiết Kế (Best Practice)

Dù Naked Return giúp code ngắn gọn hơn một chút, nó không được khuyến khích sử dụng rộng rãi vì lý do sau:

* **Giảm khả năng đọc hiểu:** Trong các hàm lớn, phức tạp có nhiều khối lệnh điều kiện, việc chỉ viết `return` ở cuối khiến người đọc code rất khó biết chính xác dữ liệu nào đang được trả về. Họ buộc phải cuộn chuột lên đầu hàm để đối chiếu.
* **Quy ước tối ưu:** Bạn **nên đặt tên** cho biến trả về để làm tài liệu hướng dẫn (self-documenting), giúp người khác biết ý nghĩa của từng cột dữ liệu trả về. Nhưng ở cuối hàm, bạn **vẫn nên trả về tường minh** tên các biến đó:
  ```go
  return fv, rfv // Tường minh, dễ đọc và dễ bảo trì!
  ```
