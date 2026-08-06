# Bài 05: Sử Dụng Interface Làm Kiểu Dữ Liệu & Implicit Implementation

> [!NOTE]
> *Chúng ta sẽ học cách dùng Interface làm kiểu dữ liệu cho tham số của hàm (`saveData`), đồng thời tìm hiểu tính chất triển khai ngầm (Implicit Implementation) độc đáo của Go.*

---

## 📤 1. Sử Dụng Interface Làm Tham Số Hàm

Mặc dù một Interface không tự chứa bất kỳ mã thực thi logic nào, nó vẫn là một **Kiểu dữ liệu (Type)** hoàn chỉnh trong Go. Điều này có nghĩa là bạn có thể sử dụng Interface ở bất kỳ đâu yêu cầu khai báo kiểu dữ liệu.

Chúng ta định nghĩa hàm tiện ích `saveData` nhận tham số `data` thuộc kiểu giao diện `saver`:

```go
func saveData(data saver) error {
	// u.Save() hoặc t.Save() tùy thuộc vào biến truyền vào
	err := data.Save() 
	
	if err != nil {
		fmt.Println("Saving the data failed:", err)
		return err
	}
	fmt.Println("Saving the data succeeded!")
	return nil
}
```

Ý nghĩa: Tham số `data saver` là một lời cam kết rằng bất kỳ giá trị nào được truyền vào hàm `saveData` đều bắt buộc phải có phương thức `Save() error`. Nhờ đó, Go biên dịch thành công dòng lệnh `data.Save()` mà không cần quan tâm `data` cụ thể là Struct `Note` hay `Todo`.

---

## 🔄 2. Gọi Hàm Tiện Ích Trong `main.go`

Nhờ có hàm tiện ích sử dụng Interface, chúng ta rút gọn được mã nguồn trong hàm `main()`, loại bỏ sự lặp lại của việc kiểm tra lỗi ghi file:

```go
func main() {
	// ... Khởi tạo todo và userNote

	todo.Display()
	err = saveData(todo) // Truyền Todo vào hàm nhận saver
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote) // Truyền Note vào hàm nhận saver
	if err != nil {
		return
	}
}
```

---

## 🌟 3. Cơ Chế Triển Khai Ngầm (Implicit Implementation) Trong Go

Tại sao chúng ta có thể truyền trực tiếp `todo` và `userNote` vào hàm `saveData` mà trình biên dịch Go không báo lỗi? 

Ở các ngôn ngữ như Java, C# hay TypeScript, bạn bắt buộc phải khai báo tường minh:
```java
// Java/C# bắt buộc phải chỉ định implements
public class Todo implements Saver { ... }
```

Nhưng trong Go:
* **Không cần từ khóa `implements`:** Go sử dụng cơ chế kiểm tra kiểu cấu trúc (structural typing), thường được gọi vui là *"Duck Typing"* (*"Nếu nó đi như vịt, kêu như vịt, thì nó là con vịt"*).
* **Kiểm tra tự động:** Khi bạn truyền biến `todo` kiểu `todo.Todo` vào hàm `saveData`, trình biên dịch Go sẽ tự động kiểm tra xem Struct `todo.Todo` có phương thức nào tên là `Save() error` hay không.
* Vì cả `todo.Todo` và `note.Note` đều tự định nghĩa phương thức `Save() error` từ trước, Go coi như chúng **tự động thỏa mãn và triển khai ngầm** Interface `saver`.

### Lợi ích:
Cơ chế triển khai ngầm giúp mã nguồn Go cực kỳ linh hoạt (loose coupling). Bạn có thể dễ dàng định nghĩa một Interface mới ở gói `main` để gom nhóm các Struct từ các bên thứ ba (thư viện ngoài) lại với nhau, miễn là các Struct đó có sẵn phương thức khớp với chữ ký của Interface, mà bạn không cần phải sửa đổi mã nguồn gốc của các Struct đó.
