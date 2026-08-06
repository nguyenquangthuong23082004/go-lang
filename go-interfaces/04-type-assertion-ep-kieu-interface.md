# Bài 04: Kỹ Thuật Type Assertion Trong Go

> [!NOTE]
> *Tìm hiểu về Type Assertion - cú pháp thay thế giúp kiểm tra và ép kiểu một cách chính xác cho một Interface rỗng (`interface{}`) về một kiểu dữ liệu cụ thể.*

---

## 📋 1. Khái Niệm Type Assertion

**Type Assertion** là một cú pháp trong Go cho phép bạn truy cập vào giá trị thực tế đằng sau một giá trị kiểu Interface (như `interface{}` hoặc `any`). 

Thay vì sử dụng Type Switch để kiểm tra nhiều kiểu cùng một lúc, Type Assertion được dùng khi bạn muốn kiểm tra xem một giá trị **có phải là một kiểu cụ thể nào đó hay không** và lấy ra giá trị đã được ép kiểu đó để sử dụng.

---

## 🧠 2. Cú Pháp Type Assertion

Cú pháp của Type Assertion sử dụng dấu chấm `.` theo sau là cặp ngoặc đơn chứa kiểu dữ liệu mong muốn:

```go
concreteValue, ok := value.(TypeName)
```

Trong đó:
* **`value`**: Biến có kiểu dữ liệu là Interface (ví dụ: `interface{}`).
* **`TypeName`**: Kiểu dữ liệu cụ thể mà bạn muốn kiểm tra (ví dụ: `int`, `string`, `float64`, hoặc một Struct cụ thể).
* **`concreteValue`**: Biến nhận giá trị thực tế sau khi ép kiểu.
  * Nếu thành công: Nó chứa giá trị gốc với kiểu dữ liệu `TypeName`.
  * Nếu thất bại: Nó chứa giá trị mặc định (zero value) của kiểu `TypeName`.
* **`ok`**: Biến kiểu logic (`bool`) báo kết quả:
  * `true`: Nếu `value` thực sự thuộc kiểu `TypeName`.
  * `false`: Nếu `value` không phải kiểu `TypeName`.

---

## 💻 3. Ví Dụ Minh Họa

Chúng ta có thể viết lại hàm `printSomething` sử dụng cú pháp Type Assertion thay thế cho Type Switch:

```go
package main

import "fmt"

func printSomething(value interface{}) {
	// 1. Kiểm tra xem có phải kiểu int hay không
	intVal, ok := value.(int)
	if ok {
		// Lúc này Go biết chắc chắn intVal có kiểu int
		// Bạn có thể cộng trừ toán học bình thường: intVal + 10
		fmt.Println("Integer:", intVal)
		return // Kết thúc hàm sớm
	}

	// 2. Kiểm tra xem có phải kiểu float64 hay không
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	// 3. Kiểm tra xem có phải kiểu string hay không
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
}
```

---

## ⚠️ 4. Rủi Ro Khi Không Kiểm Tra Giá Trị `ok` (Panic)

Go cho phép bạn viết Type Assertion dạng rút gọn chỉ nhận về 1 giá trị:

```go
concreteValue := value.(int)
```

> [!CAUTION]
> Cách viết rút gọn này cực kỳ nguy hiểm. Nếu `value` **không phải** là kiểu `int`, chương trình sẽ ngay lập tức bị **crash (panic)** tại thời điểm chạy. 
> 
> Vì vậy, luôn luôn khuyến khích sử dụng cú pháp nhận về 2 giá trị (`concreteValue, ok := value.(int)`) và kiểm tra điều kiện `if ok` trước khi xử lý để đảm bảo ứng dụng chạy an toàn.

---

## ⚖️ 5. So Sánh Type Switch Và Type Assertion

| Tiêu chí | Type Switch | Type Assertion |
| :--- | :--- | :--- |
| **Cú pháp** | `switch v := value.(type)` | `v, ok := value.(Type)` |
| **Mục đích** | Kiểm tra và phân nhánh cho **nhiều kiểu dữ liệu** khác nhau. | Kiểm tra **một kiểu dữ liệu cụ thể** duy nhất. |
| **Độ gọn gàng** | Rất gọn gàng khi có nhiều loại case khác nhau. | Dẫn đến lồng ghép nhiều câu lệnh `if` nếu kiểm tra nhiều kiểu. |
| **Tính ứng dụng** | Thích hợp cho các hàm xử lý chung (Log, Print, Format). | Rất hữu ích khi xây dựng REST API để kiểm tra kiểu dữ liệu đầu vào cụ thể. |
