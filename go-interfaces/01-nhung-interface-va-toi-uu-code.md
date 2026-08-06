# Bài 01: Nhúng Interface (Embedded Interfaces) và Tối Ưu Hóa Code Trong Go

> [!NOTE]
> *Tìm hiểu cách giải quyết vấn đề trùng lặp mã nguồn (code duplication) bằng cách sử dụng Interface, quy tắc đặt tên Interface, và kỹ thuật nhúng Interface (Embedded Interfaces) trong Go.*

---

## 📋 1. Vấn Đề Trùng Lặp Code (Code Duplication)

Khi làm việc với các Struct khác nhau như `Note` (Ghi chú) và `Todo` (Công việc cần làm), chúng ta thường thực hiện các hành động tương tự nhau:
1. Hiển thị dữ liệu lên màn hình console bằng phương thức `Display()`.
2. Lưu dữ liệu vào file system bằng phương thức `Save()`.
3. Xử lý lỗi nếu có phát sinh trong quá trình lưu dữ liệu.

Ví dụ về đoạn code trùng lặp ban đầu trong hàm `main`:
```go
// Đối với Note
userNote.Display()
err = userNote.Save()
if err != nil {
    // Xử lý lỗi...
}

// Đối với Todo
todo.Display()
err = todo.Save()
if err != nil {
    // Xử lý lỗi...
}
```

Để tối ưu hóa và tái sử dụng code, chúng ta muốn tạo một hàm dùng chung gọi là **`outputData()`** để thực hiện cả hai hành động hiển thị và lưu này.

---

## 🧠 2. Thiết Kế Interface Cho Hàm Dùng Chung

Hàm `outputData()` cần nhận vào một tham số đại diện cho cả `Note` và `Todo`. Tuy nhiên:
* Chúng ta không thể sử dụng kiểu dữ liệu cụ thể (`Note` hoặc `Todo`) vì hàm này cần hoạt động linh hoạt cho cả hai kiểu.
* Chúng ta cũng không thể sử dụng Interface `saver` hiện tại (chỉ chứa duy nhất method `Save()`) vì Interface này không đảm bảo đối tượng truyền vào sẽ có phương thức `Display()`.

### Giải Pháp Thử Nghiệm 1: Tạo Interface Riêng Biệt
Ta có thể định nghĩa một Interface mới chuyên cho việc hiển thị dữ liệu:
```go
type displayer interface {
    Display()
}
```

Tuy nhiên, nếu hàm `outputData()` sử dụng kiểu tham số là `displayer`, Go sẽ báo lỗi biên dịch khi ta cố gắng gọi phương thức `Save()` trên tham số đó, bởi vì kiểu `displayer` không đảm bảo có phương thức `Save()`.

### Giải Pháp Thử Nghiệm 2: Kết Hợp Các Phương Thức Vào Một Interface
Chúng ta cần một Interface yêu cầu **cả hai phương thức**: `Save()` và `Display()`. Ta định nghĩa Interface **`outputtable`**:
```go
type outputtable interface {
    Save() error
    Display()
}
```

Bây giờ, hàm `outputData` có thể được định nghĩa như sau:
```go
func outputData(data outputtable) error {
    data.Display()
    return data.Save()
}
```

---

## 🏷️ 3. Quy Tắc Đặt Tên Interface Trong Go

Trong Go, việc đặt tên Interface tuân theo các quy ước phổ biến sau:

1. **Quy tắc đuôi `-er` (Single-method Interface):**
   * Nếu một Interface chỉ định nghĩa **một phương thức duy nhất**, tên của nó thường kết thúc bằng hậu tố `er`.
   * Ví dụ: `saver` (chứa method `Save`), `displayer` (chứa method `Display`), `Reader`, `Writer`.

2. **Quy tắc mô tả đặc tính (Multi-method Interface):**
   * Khi Interface định nghĩa **nhiều phương thức**, ta không sử dụng hậu tố `er` nữa.
   * Thay vào đó, ta đặt tên mô tả **đặc tính hoặc khả năng** của kiểu dữ liệu mà Interface đó yêu cầu.
   * Ví dụ: **`outputtable`** (thể hiện khả năng có thể xuất dữ liệu ra ngoài - vừa hiển thị được vừa lưu được).

---

## 🧩 4. Kỹ Thuật Nhúng Interface (Embedded Interfaces)

Thay vì sao chép thủ công các khai báo phương thức từ các Interface khác vào `outputtable`, Go hỗ trợ **Nhúng Interface** (tương tự như nhúng Struct).

### Cách 1: Nhúng Hoàn Toàn Các Interface Khác
Chúng ta có thể nhúng cả hai Interface `saver` và `displayer` vào trong `outputtable`:
```go
type outputtable interface {
    saver
    displayer
}
```
*Cơ chế hoạt động:* Go sẽ tự động hiểu rằng bất kỳ kiểu dữ liệu nào muốn triển khai `outputtable` đều phải đáp ứng đầy đủ tất cả các phương thức có trong cả `saver` và `displayer`.

### Cách 2: Kết Hợp Nhúng Và Khai Báo Trực Tiếp
Nếu Interface `displayer` ít khi được sử dụng độc lập ở những nơi khác, ta có thể bỏ qua nó và định nghĩa `outputtable` bằng cách nhúng `saver` đồng thời khai báo trực tiếp phương thức `Display()`:
```go
type outputtable interface {
    saver
    Display()
}
```
Cách viết này giúp Code ngắn gọn hơn, tránh việc tạo quá nhiều Interface đơn lẻ không cần thiết trong chương trình.

---

## 🛠️ 5. Áp Dụng Thực Tế Vào Code

Với Interface `outputtable` đã được thiết kế bằng cách nhúng, chúng ta có thể tái cấu trúc hàm `main` và định nghĩa hàm `outputData` như sau:

```go
package main

import (
	// ... các packages khác ...
)

// Khai báo các interface cần thiết
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	// 1. Đối với Note
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Gọi hàm outputData để hiển thị và lưu Note
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Lỗi khi xử lý note:", err)
		return
	}

	// 2. Đối với Todo
	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Gọi hàm outputData để hiển thị và lưu Todo
	err = outputData(userTodo)
	if err != nil {
		fmt.Println("Lỗi khi xử lý todo:", err)
		return
	}
}

// Hàm dùng chung nhận vào bất kỳ dữ liệu nào thỏa mãn interface outputtable
func outputData(data outputtable) error {
	data.Display()
	return data.Save()
}
```

> [!TIP]
> * Nhờ tính năng **Implicit Implementation** (Triển khai ẩn) của Go, cả `Note` và `Todo` đều tự động thỏa mãn Interface `outputtable` mà không cần khai báo từ khóa kế thừa hay triển khai tường minh nào.
> * Việc nhúng Interface giúp tăng khả năng tái sử dụng mã nguồn và giữ cho các Interface nhỏ gọn, dễ quản lý hơn.
