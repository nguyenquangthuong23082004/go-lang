# Bài 14: Nhận Dữ Liệu Nhập Từ Người Dùng & Con Trỏ (Pointers) Cơ Bản

> [!NOTE]
> *Để ứng dụng trở nên thực tế và linh hoạt hơn, chúng ta cần cho phép người dùng tự nhập các thông số đầu tư từ bàn phím. Bài học này hướng dẫn bạn cách sử dụng hàm `fmt.Scan` và tìm hiểu sơ bộ về khái niệm con trỏ thông qua ký tự `&`.*

---

## 🆚 1. Biến (Variables) vs Hằng Số (Constants) Khi Nhập Liệu

Sự khác biệt cốt lõi giữa biến và hằng số được thể hiện rõ ràng khi chúng ta làm việc với dữ liệu đầu vào:

* **Biến (`var`):** Có thể gán lại giá trị mới (reassigned). Bạn có thể khai báo biến mà không cần gán giá trị khởi tạo ban đầu (biến sẽ tự động nhận giá trị mặc định - Zero Value) và chờ người dùng nhập giá trị mới để ghi đè vào.
* **Hằng số (`const`):** Không thể thay đổi giá trị. Bắt buộc phải gán giá trị ngay khi khai báo. Do đó, **không thể dùng hằng số để nhận dữ liệu từ người dùng**.

---

## 📥 2. Nhận Dữ Liệu Từ Terminal Bằng `fmt.Scan()`

Package `fmt` cung cấp hàm **`Scan()`** để đọc dữ liệu người dùng gõ vào terminal.

### 📍 Khái niệm Con trỏ & Ký tự và (`&`)
Nếu bạn viết `fmt.Scan(investmentAmount)`, chương trình sẽ không hoạt động như mong muốn. Để hàm `Scan` có thể trực tiếp sửa đổi giá trị lưu trong ô nhớ của biến gốc, bạn phải truyền vào một **con trỏ (pointer)** trỏ tới biến đó.

* **Cú pháp:** Thêm dấu và **`&`** ngay trước tên biến.
  ```go
  fmt.Scan(&investmentAmount)
  ```
* **Lưu ý:** Con trỏ là một chủ đề rất quan trọng và nâng cao trong Go (sẽ được học riêng ở một phần sau). Hiện tại, bạn chỉ cần nhớ quy tắc: **Luôn đặt dấu `&` trước tên biến khi truyền vào `fmt.Scan`**.

---

## 🛠️ 3. Cập Nhật Mã Nguồn Sử Dụng `fmt.Scan`

Hãy xem cách khai báo biến trống và điền giá trị thông qua nhập liệu:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	
	// Khai báo biến trống (Go tự động gán giá trị mặc định 0.0)
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Nhập số tiền đầu tư: ")
	fmt.Scan(&investmentAmount) // Nhận dữ liệu và gán vào ô nhớ của investmentAmount

	fmt.Print("Nhập lãi suất kỳ vọng (%): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Nhập số năm đầu tư: ")
	fmt.Scan(&years)

	// Thực hiện tính toán dựa trên dữ liệu người dùng vừa nhập
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Giá trị tương lai danh nghĩa:", futureValue)
	fmt.Println("Giá trị thực tế sau lạm phát:", futureRealValue)
}
```

---

## 🚦 4. Cơ Chế Chờ Nhập Liệu (Blocking)

Khi chạy chương trình bằng lệnh `go run .`:
1. Chương trình sẽ in ra dòng chữ `"Nhập số tiền đầu tư: "` và dừng lại (bị block).
2. Go Engine sẽ chờ người dùng gõ một con số vào terminal (ví dụ: `2000`) và nhấn **Enter**.
3. Sau khi nhận được giá trị, chương trình mới đi tiếp đến dòng lệnh tiếp theo để hỏi tham số thứ hai.
