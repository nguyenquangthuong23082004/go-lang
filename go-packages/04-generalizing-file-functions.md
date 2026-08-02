# Bài 04: Tổng Quát Hóa (Generalizing) Các Hàm Xử Lý File

> [!TIP]
> *Trước khi đưa mã nguồn vào một Package riêng biệt, bước đầu tiên quan trọng là phải biến đổi các hàm từ mức độ chuyên biệt (chỉ dùng cho một mục đích) thành các hàm tổng quát (có thể tái sử dụng cho bất kỳ tệp tin và dữ liệu số thực nào).*

---

## 🔄 1. Các Bước Tổng Quát Hóa Hàm

Để một hàm có tính tái sử dụng cao trong Go, chúng ta cần loại bỏ các giá trị cố định (hardcoded) và tham số hóa chúng:

### A. Đổi tên hàm & Tham số hóa Tên File
* Đổi tên `getBalanceFromFile` thành `getFloatFromFile`.
* Đổi tên `writeBalanceToFile` thành `writeFloatToFile`.
* Thay vì sử dụng trực tiếp hằng số toàn cục `accountBalanceFile`, các hàm này sẽ nhận tham số truyền vào là `fileName string`.

### B. Thay đổi tên biến & Thông báo lỗi tổng quát
* Thay thế tên biến liên quan đến "balance" (như `balanceText`, `balance`) thành các tên gọi tổng quát hơn như `valueText`, `value`.
* Thay thế thông điệp báo lỗi từ `"failed to find balance file"` thành `"failed to find file"`.

---

## 💻 2. Mã Nguồn Hàm Tổng Quát Sau Khi Refactor

Dưới đây là định nghĩa mới của hai hàm xử lý tệp tin ở cuối file `bank.go`:

```go
func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}
```

---

## 🔌 3. Gọi Hàm Tổng Quát Trong Hàm `main`

Tại hàm `main()`, chúng ta gọi các hàm mới và truyền thêm tham số đường dẫn tệp `accountBalanceFile`:

* **Khi bắt đầu ứng dụng**:
  ```go
  var accountBalance, err = getFloatFromFile(accountBalanceFile)
  ```
* **Khi cập nhật số dư (Nạp/Rút tiền)**:
  ```go
  writeFloatToFile(accountBalance, accountBalanceFile)
  ```

*Với thay đổi này, hai hàm ghi/đọc file hiện tại hoàn toàn độc lập với khái niệm "số dư ngân hàng" (balance). Chúng ta có thể dùng chúng để đọc/ghi bất kỳ số thực float64 nào vào bất kỳ file `.txt` nào trong dự án.*
