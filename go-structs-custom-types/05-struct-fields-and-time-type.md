# Bài 05: Thêm Trường Kiểu Dữ Liệu Khác & Sử Dụng Package Time

> [!NOTE]
> *Chúng ta sẽ tìm hiểu cách thêm các kiểu dữ liệu từ package bên ngoài làm trường của Struct (ví dụ `time.Time`), đồng thời tìm hiểu quy tắc dấu phẩy bắt buộc ở dòng cuối cùng của Struct trong Go.*

---

## 🕒 1. Sử Dụng Kiểu Dữ Liệu time.Time Trong Struct

Trường của một Struct có thể có bất kỳ kiểu dữ liệu nào, bao gồm cả các Struct được định nghĩa từ các package khác. Ví dụ, để ghi nhận thời gian tạo tài khoản người dùng, ta sử dụng kiểu dữ liệu `Time` của package `time`:

```go
import "time"

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time // Sử dụng struct Time của package time
}
```

---

## ⚠️ 2. Quy Tắc Dấu Phẩy Bắt Buộc (Trailing Comma Rule)

Khi khởi tạo một Struct trên nhiều dòng trong Go, bạn **bắt buộc phải có dấu phẩy (`,`) ở cuối mỗi dòng**, kể cả dòng chứa trường cuối cùng:

```go
var appUser User = User{
	firstName: userFirstName,
	lastName:  userLastName,
	birthdate: userBirthdate,
	createdAt: time.Now(), // Bắt buộc phải có dấu phẩy ở đây!
}
```

### Tại sao lại cần quy tắc này?
Go tự động thêm các dấu chấm phẩy (`;`) vào cuối các dòng lệnh khi biên dịch để giữ cho cú pháp gọn gàng. Dấu phẩy ở cuối dòng báo cho trình biên dịch biết rằng khối lệnh khai báo Struct chưa kết thúc và tránh việc chèn nhầm dấu chấm phẩy gây lỗi biên dịch.

---

## 📦 3. Hàm Trả Về Struct (Functions Returning Structs)

Trong đoạn code trên, `time.Now()` là một hàm có sẵn trong package `time` trả về một đối tượng (instance) của struct `time.Time`.
* Thay vì chỉ có thể tự tay khởi tạo struct bằng cú pháp ngoặc nhọn `{...}`, chúng ta hoàn toàn có thể khởi tạo một struct thông qua việc gọi một hàm trả về struct đó.
* Trong các bài học tiếp theo, chúng ta cũng sẽ tự viết các hàm đóng vai trò sản xuất (Factory Functions) để khởi tạo các Struct tùy chỉnh của mình.
