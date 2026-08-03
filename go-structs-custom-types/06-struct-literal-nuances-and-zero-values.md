# Bài 06: Chi Tiết Về Khởi Tạo Struct & Giá Trị Mặc Định (Zero Values)

> [!NOTE]
> *Chúng ta sẽ đi sâu vào các quy tắc và lưu ý quan trọng khi khởi tạo Struct bằng Struct Literal: từ việc bỏ qua tên trường (Key), khởi tạo Struct rỗng, cho đến việc bỏ qua một vài trường nhất định.*

---

## ⚖️ 1. Lược Bỏ Tên Trường (Positional Struct Literal)

Nếu bạn truyền các giá trị vào Struct theo đúng thứ tự các trường được định nghĩa trong blueprint, Go cho phép bạn **lược bỏ tên trường (Key)**:

```go
// Viết ngắn gọn không có tên trường (Key)
var appUser User = User{
	userFirstName,
	userLastName,
	userBirthdate,
	time.Now(),
}
```

### ⚠️ Lưu ý nguy hiểm:
* **Bắt buộc đúng thứ tự:** Nếu đảo thứ tự (ví dụ truyền `userLastName` trước `userFirstName`), giá trị sẽ bị gán sai trường mà không có lỗi biên dịch.
* **Bắt buộc đầy đủ:** Bạn không thể bỏ sót bất kỳ trường nào khi sử dụng cách viết này. Nếu bỏ sót, Go sẽ báo lỗi biên dịch.
* **Khuyên dùng:** Nên hạn chế dùng cách này, thay vào đó hãy dùng cách viết đầy đủ `Key: Value` để mã nguồn rõ ràng, dễ bảo trì.

---

## ⚪ 2. Khởi Tạo Struct Rỗng (Empty Struct)

Bạn có thể tạo một Struct mà không truyền bất kỳ giá trị nào giữa cặp ngoặc nhọn `{}`:

```go
var appUser User = User{}
```

Lúc này, đối tượng `appUser` sẽ được khởi tạo với tất cả các trường nhận **giá trị mặc định (Zero Value)** của kiểu dữ liệu tương ứng:
* `string` nhận `""` (chuỗi rỗng).
* `int`, `float64` nhận `0`.
* `bool` nhận `false`.
* Các struct khác (như `time.Time`) nhận giá trị zero của chúng.

---

## 🧩 3. Lược Bỏ Một Vài Trường (Omitting Specific Fields)

Khi sử dụng cú pháp `Key: Value`, nếu bạn không có giá trị cho một số trường nào đó, bạn hoàn toàn có thể **bỏ qua không khai báo chúng**:

```go
// Khởi tạo và bỏ qua trường birthdate
var appUser User = User{
	firstName: userFirstName,
	lastName:  userLastName,
	createdAt: time.Now(),
	// birthdate không được khai báo
}
```

Khi đó:
* Các trường được khai báo (`firstName`, `lastName`, `createdAt`) nhận giá trị được truyền vào.
* Trường bị bỏ qua (`birthdate`) sẽ tự động nhận giá trị mặc định (**Zero Value**) là `""`.
