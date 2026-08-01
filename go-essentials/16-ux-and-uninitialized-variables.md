# Bài 15: Cải Thiện Trải Nghiệm Người Dùng (UX) & Khai Báo Biến Không Khởi Tạo

> [!NOTE]
> *Để chương trình không giống như bị "treo" khi chờ nhập liệu, chúng ta cần cung cấp các chỉ dẫn rõ ràng. Bài học này hướng dẫn cách cải thiện giao diện tương tác (UX) bằng `fmt.Print` và làm rõ quy tắc khai báo biến trống.*

---

## 📺 1. Tối Ưu Hóa Trải Nghiệm (UX) Bằng `fmt.Print()`

Khi chạy một chương trình dòng lệnh tương tác, nếu màn hình chỉ hiển thị một dấu nháy trống mà không có hướng dẫn gì, người dùng sẽ bối rối. 

Chúng ta giải quyết vấn đề này bằng cách in ra lời nhắc trước khi gọi lệnh `fmt.Scan`. Để giao diện trông tự nhiên nhất, chúng ta sử dụng **`fmt.Print()`** thay vì `fmt.Println()`:

* **`fmt.Println()`**: Tự động xuống dòng ➔ Lời nhắc ở dòng trên, ô nhập liệu nhảy xuống dòng dưới.
* **`fmt.Print()`**: Không xuống dòng ➔ Ô nhập liệu nằm ngay sát lời nhắc trên cùng một dòng.

### Ví dụ so sánh:
```go
fmt.Print("Nhập số tiền đầu tư: ")
fmt.Scan(&investmentAmount)
```
*Giao diện terminal hiển thị:*
```text
Nhập số tiền đầu tư: 5000_ (người dùng nhập ngay tại đây)
```

---

## 🚨 2. Quy Tắc Khai Báo Biến Trống (Không Khởi Tạo Giá Trị)

Vì giá trị của các biến đầu vào (`investmentAmount`, `expectedReturnRate`, `years`) sẽ được nhập động từ bàn phím, chúng ta không cần gán giá trị cứng ban đầu cho chúng. 

Tuy nhiên, do Go là ngôn ngữ **kiểu dữ liệu tĩnh (statically typed)**, bạn không được phép khai báo biến trống một cách mập mờ.

### ❌ Cách viết sai (Lỗi biên dịch):
```go
var years
// Hoặc:
years :=
```
*Lý do:* Go không thể biết biến này dùng để chứa kiểu dữ liệu gì (chuỗi, số nguyên hay số thực) để cấp phát vùng nhớ phù hợp.

###  Cách viết đúng:
Khi không gán giá trị khởi tạo, bạn **bắt buộc phải ghi rõ kiểu dữ liệu**:
```go
var years float64
```
Lúc này, Go sẽ khởi tạo biến đó ở dạng **Zero Value** của kiểu tương ứng (đối với kiểu `float64` là `0.0`) trước khi được hàm `Scan` ghi đè giá trị mới vào ô nhớ.

---

## 💻 3. Mã Nguồn Chương Trình Hoàn Chỉnh

Dưới đây là cấu trúc chương trình sau khi được tối ưu hóa toàn bộ phần nhập liệu và hiển thị:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	
	// Khai báo các biến đầu vào không khởi tạo giá trị
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	// Giao diện nhập liệu thân thiện
	fmt.Print("Nhập số tiền đầu tư: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Nhập lãi suất kỳ vọng (%): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Nhập số năm đầu tư: ")
	fmt.Scan(&years)

	// Tính toán kết quả
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Hiển thị kết quả đầu ra
	fmt.Println("Giá trị tương lai danh nghĩa:", futureValue)
	fmt.Println("Giá trị thực tế sau lạm phát (trừ đi lạm phát 2.5%):", futureRealValue)
}
```
