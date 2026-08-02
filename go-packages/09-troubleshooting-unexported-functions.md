# Bài 09: Khắc Phục Lỗi "Undefined" Khi Sử Dụng Hàm Từ Package Khác

> [!NOTE]
> *Khi mới bắt đầu làm việc với các gói tự tạo trong Go, một lỗi cực kỳ phổ biến mà bạn sẽ gặp phải là trình biên dịch báo lỗi hàm gọi từ gói khác là `undefined` (chưa được định nghĩa), mặc dù bạn chắc chắn đã định nghĩa hàm đó. Bài này giúp bạn cách chẩn đoán và khắc phục lỗi này.*

---

## 🔍 1. Hiện Tượng Lỗi

Khi bạn import gói `fileops` thành công vào file `bank.go`:
```go
import "example.com/bank/fileops"
```
Và gọi hàm trong hàm `main()`:
```go
accountBalance, err := fileops.getFloatFromFile(accountBalanceFile)
```
Trình biên dịch của Go sẽ báo lỗi đỏ:
```text
fileops.getFloatFromFile undefined (cannot refer to unexported identifier fileops.getFloatFromFile)
```

---

## 🧠 2. Nguyên Nhân Bản Chất

Lỗi này phát sinh do quy tắc đóng gói nghiêm ngặt của Go:
* Hàm `getFloatFromFile` trong gói `fileops.go` được bắt đầu bằng chữ cái **`g` viết thường**.
* Do đó, Go Compiler coi đây là một **định danh không xuất khẩu (unexported identifier)** hay hàm nội bộ (private).
* Bất kỳ nỗ lực nào gọi hàm này từ package ngoài (ở đây là package `main` của file `bank.go`) đều bị chặn đứng bởi trình biên dịch.

---

## 🛠️ 3. Giải Pháp Khắc Phục

Để sửa lỗi này, chúng ta cần thực hiện đổi tên (refactor) đồng bộ ở cả nơi định nghĩa và nơi gọi hàm:

### Bước 1: Sửa tên hàm trong gói nguồn (`fileops/fileops.go`)
Đổi chữ cái đầu tiên từ viết thường sang **viết hoa**:
```diff
-func getFloatFromFile(fileName string) (float64, error) {
+func GetFloatFromFile(fileName string) (float64, error) {
     // ...
 }
```

### Bước 2: Sửa tên hàm ở nơi gọi (`bank.go`)
Đổi tên gọi tương ứng với chữ cái viết hoa:
```diff
-accountBalance, err := fileops.getFloatFromFile(accountBalanceFile)
+accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
```

*Tương tự áp dụng cho hàm `writeFloatToFile` $\rightarrow$ `WriteFloatToFile`. Sau khi đổi tên, lỗi biên dịch `undefined` sẽ hoàn toàn biến mất và chương trình hoạt động bình thường.*
