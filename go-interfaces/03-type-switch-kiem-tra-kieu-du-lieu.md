# Bài 03: Kiểm Tra Kiểu Dữ Liệu Với Type Switch Trong Go

> [!NOTE]
> *Tìm hiểu về Type Switch - một phiên bản đặc biệt của câu lệnh `switch` trong Go giúp xác định kiểu dữ liệu thực tế và thực hiện các hành vi xử lý khác nhau dựa trên kiểu dữ liệu của một Interface rỗng (`interface{}`).*

---

## 📋 1. Tại Sao Cần Type Switch?

Như đã tìm hiểu ở bài trước, khi một hàm nhận tham số kiểu `interface{}` (hoặc `any`), nó có thể chấp nhận bất kỳ giá trị nào. 
Tuy nhiên, trong thực tế, chúng ta thường chỉ muốn xử lý một số kiểu dữ liệu nhất định, và với mỗi kiểu chúng ta lại có các hành động xử lý riêng biệt.

Để giải quyết vấn đề này, Go cung cấp cú pháp **Type Switch** giúp kiểm tra kiểu dữ liệu thực tế (dynamic type) của một giá trị lúc chạy chương trình (run-time).

---

## 🧠 2. Cú Pháp Type Switch

Cú pháp của Type Switch tương tự như `switch-case` thông thường, nhưng biểu thức kiểm tra ở đầu câu lệnh sẽ sử dụng cú pháp đặc biệt là **`.(type)`**:

```go
switch value.(type) {
case kiểu_dữ_liệu_1:
    // Code xử lý cho kiểu dữ liệu 1
case kiểu_dữ_liệu_2:
    // Code xử lý cho kiểu dữ liệu 2
default:
    // Code xử lý mặc định nếu không khớp kiểu nào ở trên (hoặc bỏ qua)
}
```

---

## 💻 3. Ví Dụ Minh Họa

Chúng ta định nghĩa hàm `printSomething` để lọc và hiển thị thông tin tùy theo kiểu dữ liệu nhận được:

```go
package main

import "fmt"

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float 64:", value)
	case string:
		fmt.Println("String:", value)
	default:
		// Không định nghĩa code xử lý ở đây -> bỏ qua các kiểu dữ liệu khác
	}
}
```

### Cách Hoạt Động Trong Hàm `main`:

```go
func main() {
	printSomething(1)     // Output: Integer: 1 (Khớp case int)
	printSomething(1.5)   // Output: Float 64: 1.5 (Khớp case float64)
	printSomething("Go")  // Output: String: Go (Khớp case string)

	// Giả sử ta truyền vào một đối tượng kiểu Todo (Struct)
	userTodo, _ := todo.New("Học Go")
	printSomething(userTodo) 
	
	// Kết quả: Không có gì xảy ra và chương trình KHÔNG BỊ LỖI/CRASH.
	// Trình biên dịch chấp nhận truyền vào (vì là interface{}), 
	// nhưng khi chạy, Type Switch không khớp với case nào nên lặng lẽ bỏ qua.
}
```

---

## 💡 4. Mở Rộng: Gán Và Trích Xuất Giá Trị Thực (Type Switch Guard)

Cú pháp `switch value.(type)` phía trên chỉ giúp rẽ nhánh theo kiểu dữ liệu, bản thân biến `value` bên trong các case vẫn mang kiểu `interface{}`. 

Nếu bạn muốn biến đó tự động chuyển đổi sang kiểu dữ liệu cụ thể tương ứng trong mỗi case để thực hiện các thao tác chuyên biệt (ví dụ: cộng trừ toán học với kiểu `int`), hãy sử dụng cú pháp gán sau:

```go
func processValue(value interface{}) {
    // Sử dụng cú pháp gán để tạo biến mới (ví dụ: concreteValue) chứa giá trị đã được ép kiểu tự động
    switch concreteValue := value.(type) {
    case int:
        // Trong case này, concreteValue có kiểu là int (có thể cộng trừ)
        fmt.Println("Integer cộng thêm 10:", concreteValue + 10)
    case string:
        // Trong case này, concreteValue có kiểu là string (có thể dùng len(), nối chuỗi...)
        fmt.Println("Độ dài chuỗi:", len(concreteValue))
    }
}
```

---

## 📋 5. Tóm Tắt Đặc Điểm Của Type Switch

* **Linh hoạt & An toàn:** Cho phép viết các hàm nhận mọi kiểu dữ liệu nhưng vẫn kiểm soát tốt các kiểu dữ liệu thực tế muốn xử lý.
* **Không crash chương trình:** Nếu kiểu dữ liệu truyền vào không khớp với bất kỳ `case` nào (và không có `default`), Go sẽ bỏ qua một cách an toàn mà không gây crash chương trình hay báo lỗi biên dịch.
* **Cú pháp đặc biệt `.(type)`:** Chỉ dùng được duy nhất trong cấu trúc câu lệnh `switch`. Bạn không thể sử dụng `value.(type)` ở ngoài câu lệnh `switch` này.
