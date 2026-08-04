# Bài 03: Tích Hợp Todo Vào Main & Vấn Đề Trùng Lặp Hàm

> [!NOTE]
> *Chúng ta sẽ tích hợp luồng nhập liệu, hiển thị và lưu trữ cho đối tượng Todo vào trong `main.go`, từ đó trực quan hóa sự lặp lại trong code.*

---

## 💻 1. Mã Nguồn Cập Nhật Tại `main.go`

Chúng ta import gói `todo` và cập nhật luồng xử lý chính trong `main.go` để chạy đồng thời cả ghi chú (Note) và công việc cần làm (Todo):

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	// Khởi tạo Todo
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Khởi tạo Note
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Xử lý Todo (Hiển thị & Lưu)
	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed:", err)
		return
	}
	fmt.Println("Saving the todo succeeded!")

	// Xử lý Note (Hiển thị & Lưu)
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return
	}
	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
```

---

## ⚠️ 2. Phân Tích Vấn Đề Trùng Lặp Mã Nguồn

Hãy quan sát kỹ hai khối lệnh xử lý hiển thị và lưu trữ dưới đây:

```go
// Khối xử lý Todo
todo.Display()
err = todo.Save()
if err != nil { ... }
fmt.Println("Saving the todo succeeded!")

// Khối xử lý Note
userNote.Display()
err = userNote.Save()
if err != nil { ... }
fmt.Println("Saving the note succeeded!")
```

### Điểm bất cập:
* **Thuật toán giống hệt nhau:** Cả hai khối lệnh đều thực hiện chính xác các bước: hiển thị thông tin -> lưu file -> kiểm tra lỗi -> in thông báo thành công.
* **Không thể viết hàm phụ trợ dùng chung:** Trong Go, nếu bạn muốn viết một hàm tiện ích ví dụ `saveData(data ???)` để đóng gói các bước trên, bạn sẽ gặp lỗi biên dịch vì kiểu dữ liệu của `todo` (`todo.Todo`) và `userNote` (`note.Note`) là hoàn toàn khác nhau.
* **Giải pháp:** Chúng ta cần định nghĩa một **Interface** chung đóng vai trò là một "hợp đồng hành vi" để hàm tiện ích có thể chấp nhận bất kỳ kiểu dữ liệu nào miễn là nó có đủ 2 phương thức `Display()` và `Save() error`.
