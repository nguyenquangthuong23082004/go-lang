# Bài 05: Hạn Chế Của Interface Rỗng Và Sự Cần Thiết Của Generics

> [!NOTE]
> *Tìm hiểu về những hạn chế lớn của `interface{}` (hoặc `any`) khi giải quyết bài toán viết code linh hoạt nhưng vẫn cần kiểm soát hành vi (ví dụ: hàm cộng hai giá trị), dẫn đến sự ra đời của Generics trong Go.*

---

## 📋 1. Bài Toán Thiết Kế Hàm `add()` Linh Hoạt

Giả sử chúng ta muốn viết một hàm `add(a, b)` để cộng/ghép hai giá trị lại với nhau. Tuy nhiên, hàm này không được cố định kiểu dữ liệu (như chỉ nhận `int` hay `float64`). Thay vào đó:
* Nó có thể cộng hai số nguyên (`int`).
* Nó có thể cộng hai số thực (`float64`).
* Nó có thể nối hai chuỗi ký tự (`string`).

---

## 🧠 2. Hạn Chế Khi Sử Dụng Interface Rỗng (`interface{}`)

Một giải pháp tự nhiên là sử dụng kiểu `interface{}` (hoặc `any`) cho cả tham số đầu vào và đầu ra để chấp nhận mọi kiểu dữ liệu:

```go
func add(a interface{}, b interface{}) interface{} {
    return a + b // LỖI BIÊN DỊCH!
}
```

### Tại sao lại bị lỗi biên dịch?
Trình biên dịch của Go sẽ báo lỗi: **`invalid operation: a + b (operator + not defined on interface)`**. 
Nguyên nhân là vì kiểu `interface{}` chấp nhận **mọi thứ** (bao gồm cả Struct, Array, Map, Boolean...), mà toán tử cộng `+` chỉ được định nghĩa trên một số kiểu dữ liệu nhất định (số và chuỗi). Cách tiếp cận này quá linh hoạt và không an toàn.

---

## 🛠️ 3. Giải Pháp Tạm Thời Bằng Type Assertion (Workaround)

Để giải quyết lỗi trên bằng các kiến thức hiện có, chúng ta bắt buộc phải sử dụng **Type Assertion** để ép kiểu thủ công từng trường hợp:

```go
func add(a interface{}, b interface{}) interface{} {
	// 1. Kiểm tra nếu cả hai đều là int
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}

	// 2. Kiểm tra nếu cả hai đều là float64
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	// 3. Kiểm tra nếu cả hai đều là string
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)
	if aIsString && bIsString {
		return aString + bString
	}

	// Trả về nil nếu kiểu dữ liệu không hợp lệ hoặc không khớp nhau
	return nil
}
```

---

## ⚠️ 4. Những Điểm Yếu Của Giải Pháp Ép Kiểu Thủ Công

Mặc dù hàm trên hoạt động được, nó bộc lộ 3 nhược điểm cực kỳ lớn:

1. **Quá nhiều code rườm rà (Boilerplate Code):** 
   Chúng ta phải viết đi viết lại các đoạn logic kiểm tra ép kiểu tương tự nhau cho từng kiểu dữ liệu muốn hỗ trợ. Nếu sau này cần hỗ trợ thêm kiểu `int64` hay `float32`, lượng code sẽ phình to rất nhanh.

2. **Mất an toàn kiểu ở kết quả trả về:** 
   Vì hàm phải trả về `interface{}`, nên người gọi hàm (caller) sau khi nhận kết quả lại phải tiếp tục sử dụng Type Assertion một lần nữa để lấy về kiểu thực tế (`int`, `float64` hoặc `string`) trước khi có thể tính toán tiếp.

3. **Thiếu kiểm soát ràng buộc kiểu dữ liệu:**
   Chúng ta mong muốn hai tham số truyền vào phải **cùng kiểu với nhau** (cùng là `int` hoặc cùng là `string`). Nhưng với `interface{}`, chúng ta hoàn toàn có thể truyền `add(5, "Hello")`. Chương trình sẽ compile thành công nhưng âm thầm trả về `nil` lúc chạy (run-time), gây khó khăn cho việc gỡ lỗi.

---

## 🚀 5. Giải Pháp Tối Ưu: Sự Ra Đời Của Generics (Kiểu Chung)

Để giải quyết triệt để vấn đề này, Go đã giới thiệu tính năng **Generics (Kiểu chung)** từ phiên bản 1.18. Generics cho phép chúng ta:
1. Xác định kiểu dữ liệu thực tế tại thời điểm gọi hàm thay vì cố định lúc viết hàm.
2. Giữ nguyên tính an toàn kiểu dữ liệu (Type Safety) lúc compile-time.
3. Tránh việc viết code trùng lặp (Boilerplate).

---

## 🎨 6. Cú Pháp Khai Báo Generics Trong Go

Chúng ta biến đổi hàm `add` thông thường thành một hàm Generic bằng cách thêm **ngoặc vuông `[...]`** ngay sau tên hàm và trước danh sách tham số:

```go
func add[T any](a T, b T) T {
    return a + b // Vẫn sẽ báo lỗi nếu T là 'any'
}
```

Trong đó:
* **`T`**: Là một **Type Parameter** (Tham số kiểu) - đóng vai trò là một placeholder (trình giữ chỗ) đại diện cho kiểu dữ liệu sẽ được truyền vào sau này. Theo quy ước, ta thường đặt tên là `T` (viết tắt của Type), nhưng bạn có thể đặt bất kỳ tên nào tùy ý.
* **`any`**: Là một **Type Constraint** (Ràng buộc kiểu) - quy định những kiểu dữ liệu nào được phép truyền vào thay thế cho `T`. Ở đây, `any` có nghĩa là chấp nhận mọi kiểu dữ liệu.

### Vấn đề khi dùng ràng buộc `any`:
Nếu đặt ràng buộc là `any`, Go sẽ tiếp tục báo lỗi tại dòng `return a + b` vì `any` vẫn quá rộng và không đảm bảo kiểu truyền vào có hỗ trợ toán tử `+` (ví dụ: không thể cộng hai Struct với nhau).

---

## 🔒 7. Giới Hạn Phạm Vi Kiểu Dữ Liệu (Union Type Constraints)

Để giải quyết lỗi trên, thay vì dùng `any`, chúng ta có thể giới hạn danh sách các kiểu dữ liệu cụ thể được phép sử dụng cho `T`. Chúng ta liệt kê chúng và phân tách nhau bằng ký tự gạch đứng **`|`** (gọi là Union Type):

```go
// Định nghĩa danh sách kiểu dữ liệu được phép sử dụng cho T
func add[T int | float64 | string](a T, b T) T {
    // Bây giờ Go biết chắc chắn T chỉ có thể là int, float64 hoặc string.
    // Cả ba kiểu này đều hỗ trợ toán tử '+', nên phép cộng này hoàn toàn hợp lệ!
    return a + b 
}
```

Bằng cách này, chúng ta đã giải quyết được vấn đề lỗi biên dịch toán tử `+` mà không cần viết các câu lệnh kiểm tra kiểu dài dòng.

---

## ⚡ 8. Cách Sử Dụng Và Cơ Chế Tự Động Suy Luận Kiểu (Type Inference)

Khi sử dụng hàm Generic `add`, chúng ta có thể gọi nó như sau:

```go
func main() {
    // Gọi hàm với tham số kiểu int
    resultInt := add(10, 20)
    fmt.Println(resultInt + 5) // Hợp lệ! Go hiểu resultInt có kiểu int
    
    // Gọi hàm với tham số kiểu float64
    resultFloat := add(1.5, 2.5)
    fmt.Println(resultFloat) // Hợp lệ!
    
    // Gọi hàm với tham số kiểu string
    resultStr := add("Hello, ", "Go")
    fmt.Println(resultStr) // Hợp lệ!
}
```

### Cơ chế hoạt động của Go:
1. **Type Inference (Suy luận kiểu):** Khi bạn truyền `add(10, 20)`, Go nhìn vào kiểu dữ liệu của đối số truyền vào (`10` và `20` là `int`) và tự động suy luận rằng `T` trong lượt gọi này sẽ là kiểu `int`.
2. **Compile-time Safety:** Vì `T` là `int`, kiểu trả về của hàm cũng tự động được xác định là `int`. Biến `resultInt` nhận giá trị kiểu `int` thực tế, do đó bạn có thể tính toán tiếp (`resultInt + 5`) mà không cần bất kỳ bước ép kiểu (Type Assertion) nào khác.
3. **Đồng nhất kiểu dữ liệu:** Nếu bạn cố tình truyền sai kiểu dữ liệu khác nhau, ví dụ: `add(10, "Hello")`, trình biên dịch Go sẽ báo lỗi ngay lập tức vì cả hai tham số đầu vào đều được định nghĩa chung một kiểu `T`.

---

## 🎯 9. Khi Nào Nên Sử Dụng Generics?

Mặc dù Generics rất mạnh mẽ, bạn không nên lạm dụng nó cho mọi hàm trong dự án:
* **Nên dùng:** Khi viết các thư viện dùng chung (Utility Libraries), các cấu trúc dữ liệu dùng chung (như Stack, Queue, LinkedList tự định nghĩa), hoặc các hàm xử lý dữ liệu tổng quát (như hàm tìm min/max, lọc mảng) mà bạn không biết trước hoặc muốn hỗ trợ nhiều kiểu dữ liệu khác nhau nhưng cùng chung thuật toán xử lý.
* **Không nên dùng:** Đối với các hàm xử lý logic nghiệp vụ (business logic) thông thường. Hãy luôn ưu tiên kiểu dữ liệu cụ thể hoặc Interface chuẩn để giữ cho code đơn giản và trực quan nhất.

