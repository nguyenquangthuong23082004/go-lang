# Bài 11: Khai Báo Kiểu Dữ Liệu Tường Minh (Explicit Type Assignment) trong Go

> [!NOTE]
> *Việc lạm dụng ép kiểu cục bộ (`float64(biến)`) sẽ làm cho mã nguồn trở nên dài dòng và khó đọc. Bài học này hướng dẫn bạn cách tối ưu hóa bằng cách khai báo kiểu dữ liệu một cách tường minh ngay từ đầu.*

---

## 🧠 1. Cơ Chế Mặc Định: Tự Động Suy Luận Kiểu (Type Inference)

Khi bạn viết code mà không khai báo kiểu dữ liệu:
```go
var investmentAmount = 1000 // Go tự suy luận kiểu: int
var expectedReturnRate = 5.5 // Go tự suy luận kiểu: float64
```
Go sẽ tự động quét qua giá trị bên phải dấu bằng để gán kiểu dữ liệu tương ứng.

### 🚨 Vấn đề phát sinh:
Nếu bạn có một biến tỷ lệ lợi suất được viết là `5` (thay vì `5.5`), Go sẽ suy luận kiểu của biến này là `int`. Tuy nhiên, tỷ lệ lợi suất về mặt logic phải là kiểu số thực (`float64`) vì nó có thể nhận các giá trị lẻ như `5.25`, `5.75`. Việc Go tự suy luận thành `int` sẽ gây ra lỗi kiểu dữ liệu khi thực hiện phép chia hoặc nhân với các số thực khác.

---

## 🔄 2. Giải Pháp: Khai Báo Kiểu Tường Minh (Explicit Type Assignment)

Để ghi đè lên cơ chế tự động của Go, bạn có thể khai báo trực tiếp kiểu dữ liệu mong muốn ngay sau tên biến theo cú pháp:
```go
var <tên_biến> <kiểu_dữ_liệu> = <giá_trị>
```

### Ví dụ áp dụng:
```go
var investmentAmount float64 = 1000
var years float64 = 10
```

* **Cơ chế hoạt động:** Trình biên dịch Go sẽ hiểu và lưu trữ các số `1000` và `10` trong bộ nhớ dưới dạng số thực `float64` (tương đương `1000.0` và `10.0`), cho phép chúng tham gia trực tiếp vào các phép toán số thực mà không cần ép kiểu.

---

## 🚀 3. Tối Ưu Hóa Công Thức Tính Toán

Sau khi đã khai báo tường minh kiểu dữ liệu cho tất cả các biến đầu vào, chúng ta có thể loại bỏ hoàn toàn các hàm ép kiểu cục bộ `float64()` trong công thức, giúp mã nguồn trở nên vô cùng gọn gàng và dễ đọc:

### ❌ Code trước khi tối ưu (Ép kiểu cục bộ):
```go
var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
```

###  Code sau khi tối ưu (Khai báo tường minh):
```go
var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
```

---

## 📝 4. Khi Nào Nên Dùng Cách Nào?

| Phương pháp | Trường hợp sử dụng tốt nhất |
| :--- | :--- |
| **Khai báo kiểu tường minh** | Khi biến số **chỉ phục vụ** cho một công thức toán học yêu cầu kiểu dữ liệu cụ thể (ví dụ: chỉ tính toán số thực trong dự án này). |
| **Ép kiểu cục bộ (`float64()`)** | Khi biến số đó bắt buộc phải giữ kiểu số nguyên `int` để thực hiện một số phép toán số nguyên ở chỗ khác (ví dụ: đếm số vòng lặp, chỉ mục mảng) nhưng cần tham gia phép tính số thực tại một vị trí duy nhất. |
