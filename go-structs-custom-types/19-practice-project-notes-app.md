# Bài 19: Dự Án Thực Hành - Ứng Dụng Ghi Chú (Notes App)

> [!NOTE]
> *Hướng dẫn liền mạch và toàn diện từ đầu đến cuối về cách xây dựng ứng dụng ghi chú (Notes App) bằng Go: từ việc dựng khung dự án, xử lý nhập chuỗi có dấu cách, tách package con, đến việc chuyển đổi sang JSON và lưu file.*

---

## 📂 1. Cấu Trúc Thư Mục Dự Án

Chúng ta xây dựng ứng dụng ghi chú với cấu trúc thư mục gồm gói chính `main` và gói con `note` phụ trách lưu trữ thông tin:

```
/learning/07-note/
├── go.mod
├── main.go  (Gói main - Điểm chạy chính)
└── note/
    └── note.go (Gói note - Định nghĩa Struct Note)
```

---

## ⌨️ 2. Đọc Chuỗi Nhiều Từ Có Khoảng Trắng (bufio & strings)

### A. Vấn đề của họ hàm Scan (`fmt.Scan` / `fmt.Scanln`)
Các hàm `Scan` mặc định coi khoảng trắng (space) làm ký tự phân tách đối số. Khi bạn nhập một chuỗi chứa khoảng trắng (ví dụ: `Learn Go`), hàm chỉ đọc từ đầu tiên (`Learn`), từ thứ hai (`Go`) sẽ bị giữ lại trong bộ đệm và tự động gán vào ô nhập tiếp theo khiến chương trình bị nhảy dòng.

### B. Giải pháp bằng `bufio` & `strings`
Chúng ta tạo một `Reader` lắng nghe trực tiếp từ luồng nhập chuẩn của bàn phím (`os.Stdin`) và đọc cho đến khi người dùng nhấn xuống dòng (`\n`). Sau đó sử dụng `strings.TrimSuffix` để dọn sạch ký tự xuống dòng ở cuối chuỗi:

```go
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// Đọc luồng nhập từ bàn phím
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // '\n' là kiểu dữ liệu Rune (nháy đơn)
	if err != nil {
		return ""
	}

	// Cắt bỏ ký tự xuống dòng ở cuối chuỗi (\n cho Linux/macOS, \r cho Windows)
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
```

---

## 🏗️ 3. Định Nghĩa Struct Note & Quy Tắc Struct Tags

Để lưu trữ thông tin ghi chú và ghi nó ra file dưới dạng **JSON**, chúng ta sử dụng package `"encoding/json"` của Go:

1. **Bắt buộc viết hoa tên trường (Exported Fields):** Package JSON hoạt động từ bên ngoài và chỉ có thể đọc các trường viết hoa chữ cái đầu (`Title`, `Content`, `CreatedAt`). Nếu để chữ thường, chuỗi JSON kết quả thu được sẽ trống rỗng `{}`.
2. **Sử dụng Struct Tags:** Để các thuộc tính viết hoa của Go hiển thị đẹp đẽ dưới dạng chữ thường khi chuyển thành JSON, chúng ta gắn thêm tag phía sau thuộc tính:

```go
type Note struct {
	Title     string    `json:"title"`      // Xuất thành khóa "title" trong JSON
	Content   string    `json:"content"`    // Xuất thành khóa "content" trong JSON
	CreatedAt time.Time `json:"created_at"` // Xuất thành khóa "created_at" trong JSON
}
```

---

## 🛠️ 4. Constructor Function & Các Methods

Để đóng gói dữ liệu và quản lý tập trung, toàn bộ logic xử lý dữ liệu của Note được cài đặt bên trong package `note`:

* **Constructor `New`:** Hàm khởi tạo nhận dữ liệu, tiến hành kiểm tra ràng buộc (Validation) để tránh tạo ghi chú rỗng, trả về kiểu giá trị thường `Note` và lỗi `error` (thay vì con trỏ vì struct này có kích thước nhỏ gọn).
* **Method `Display`:** Phương thức công khai in dữ liệu ghi chú ra màn hình theo định dạng.
* **Method `Save`:** Chuyển đổi đối tượng Note thành chuỗi JSON bằng `json.Marshal(n)` và ghi xuống ổ cứng thông qua `os.WriteFile` với tên file tự động hóa (chuyển chữ thường, thay khoảng trắng bằng gạch dưới `_`).

---

## 💻 5. Mã Nguồn Hoàn Chỉnh Của Dự Án

### 📄 File [note/note.go](file:///home/uwal/Desktop/go-lang/learning/07-note/note/note.go)
```go
package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonBytes, 0644)
}
```

### 📄 File [main.go](file:///home/uwal/Desktop/go-lang/learning/07-note/main.go)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

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
