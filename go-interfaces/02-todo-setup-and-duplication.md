# Bài 02: Khởi Tạo Gói Todo & Sự Trùng Lặp Logic

> [!NOTE]
> *Chúng ta sẽ chuẩn bị môi trường để học Interfaces bằng cách tạo thêm một package `todo` độc lập. Từ đó, ta sẽ thấy sự trùng lặp trong việc gọi các phương thức giống nhau như `Display()` và `Save()` của hai đối tượng khác nhau.*

---

## 📂 1. Cấu Trúc Dự Án Mới Với Gói Todo

Để mở rộng ứng dụng ghi chú hiện tại thành một ứng dụng quản lý cả công việc cần làm (Todo), chúng ta tạo thêm thư mục `todo` bên trong dự án thực hành `/learning/07-note/`:

```
/learning/07-note/
├── go.mod
├── main.go
├── note/
│   └── note.go (Định nghĩa ghi chú - Note)
└── todo/
    └── todo.go (Định nghĩa công việc - Todo)
```

---

## 📝 2. Mã Nguồn Gói Todo (`todo/todo.go`)

Struct `Todo` có cấu trúc đơn giản hơn `Note` (không có tiêu đề và ngày tạo, chỉ có thuộc tính nội dung công việc):

```go
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json" // Tên file cố định cho Todo

	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonBytes, 0644)
}
```

---

## ⚠️ 3. Phân Tích Sự Trùng Lặp & Cơ Hội Cho Interface

Nhìn vào hai Struct `Note` và `Todo`, ta thấy mặc dù cấu trúc dữ liệu của chúng khác nhau (một bên chứa `Title`, `Content`, `CreatedAt`; một bên chỉ chứa `Text`), chúng đều chia sẻ chung hai hành vi:

1. **Hành vi hiển thị:** Đều có phương thức `Display()`.
2. **Hành vi lưu trữ:** Đều có phương thức `Save() error`.

Nếu trong hàm `main.go`, chúng ta muốn viết một hàm chung chịu trách nhiệm nhận vào bất kỳ dữ liệu nào của người dùng (dù là `Note` hay `Todo`), tiến hành in ra màn hình và lưu vào file, chúng ta sẽ gặp khó khăn do tính chất Tĩnh (Statically typed) của Go:

```go
// Chúng ta bắt buộc phải viết 2 hàm riêng biệt vì kiểu dữ liệu khác nhau:
func printAndSaveNote(n note.Note) { ... }
func printAndSaveTodo(t todo.Todo) { ... }
```

Việc này gây trùng lặp code rất nhiều. Để giải quyết vấn đề này, chúng ta cần một cơ chế cho phép định nghĩa một kiểu dữ liệu chung đại diện cho *"bất kỳ thứ gì có thể Display và Save"*. Đó chính là lúc **Interfaces** xuất hiện để giải quyết bài toán Đa hình (Polymorphism).
