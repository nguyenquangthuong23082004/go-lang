# Bài 02: Functions As First-Class Values (Hàm Là Một Giá Trị) & Custom Function Types

> [!NOTE]
> *Trong Go, **Hàm (Function) là một "Công dân hạng nhất" (First-Class Citizen)**. Điều này có nghĩa là bạn có thể đối xử với Hàm giống hệt như các giá trị thông thường khác (`int`, `string`, `slice`): lưu trữ vào biến, truyền làm tham số cho hàm khác, hoặc định nghĩa kiểu dữ liệu cho hàm.*

---

## 🌟 1. Khái Niệm "Functions as First-Class Values"

Lập trình viên thường quen với việc truyền số, chuỗi, hoặc slice làm đối số cho hàm. Trong Go, bạn cũng có thể **truyền một Hàm làm đối số (parameter) cho một Hàm khác**:

- Hàm nhận vào một hàm khác làm tham số được gọi là **Higher-Order Function**.
- Thao tác này giúp tránh lặp lại code (DRY - Don't Repeat Yourself) và tạo ra các hàm xử lý tổng quát (Generic/Reusable Utility Functions).

---

## 💻 2. Ví Dụ Cụ Thể: Hàm Biến Đổi Mảng (`transformNumbers`)

Giả sử bạn cần nhân đôi (`double`) hoặc nhân ba (`triple`) từng phần tử trong một danh sách số. Thay vì viết 2 hàm lặp mảng riêng biệt, ta tạo một hàm xử lý tổng quát `transformNumbers`:

```go
package main

import "fmt"

// Định nghĩa Custom Type cho kiểu Hàm (Function Type Alias)
// Hàm này nhận vào 1 số int và trả về 1 số int
type transformFn func(int) int

func main() {
    numbers := []int{1, 2, 3, 4}

    // Truyền tên hàm `double` (KHÔNG CÓ cặp ngoặc `()`) làm đối số
    doubled := transformNumbers(&numbers, double)
    fmt.Println("Doubled Numbers:", doubled) // Output: [2 4 6 8]

    // Truyền tên hàm `triple` làm đối số
    tripled := transformNumbers(&numbers, triple)
    fmt.Println("Tripled Numbers:", tripled) // Output: [3 6 9 12]
}

// Higher-Order Function: Nhận tham số `transform` là một hàm có kiểu `transformFn`
func transformNumbers(numbers *[]int, transform transformFn) []int {
    dNumbers := []int{}

    for _, val := range *numbers {
        // Thực thi hàm transform được truyền vào
        dNumbers = append(dNumbers, transform(val))
    }

    return dNumbers
}

// Các hàm Helper matching với kiểu transformFn (func(int) int)
func double(number int) int {
    return number * 2
}

func triple(number int) int {
    return number * 3
}
```

---

## ⚠️ 3. Lưu Ý Cực Kỳ Quan Trọng Khi Truyền Hàm

Khi truyền hàm làm đối số:
- **ĐÚNG:** Truyền tên hàm **KHÔNG CÓ DẤU NGHỆCH `()`** (Ví dụ: `double`).  
  👉 Nghĩa là bạn đang truyền **chính bản thân hàm đó** đi như một giá trị.
- **SAI:** Truyền tên hàm **CÓ DẤU NGHỆCH `()`** (Ví dụ: `double(5)`).  
  👉 Nghĩa là bạn đang **thực thi hàm** đó ngay lập tức và truyền **kết quả trả về** (`10`) vào hàm khác.

---

## 🎨 4. Custom Function Types (`type ... func(...) ...`)

Cú pháp khai báo kiểu hàm có thể rất dài nếu hàm có nhiều tham số. Go cho phép bạn tạo một **Type Alias cho kiểu Hàm**:

```go
// Thay vì gõ func(int, string, map[string]int) (int, error) nhiều lần
type complexFn func(int, string, map[string]int) (int, error)

// Bạn sử dụng ngắn gọn:
func process(fn complexFn) {
    // ...
}
```

---

## 📝 Tóm Tắt

1. **Hàm là giá trị:** Bạn có thể truyền hàm dưới dạng tham số mà không có cặp dấu `()`.
2. **Cú pháp định nghĩa tham số hàm:** `func(kiểu_tham_số) kiểu_trả_về`.
3. **Custom Function Type:** Dùng `type TênKiểu func(...) ...` để rút gọn kiểu hàm giúp code dễ đọc và tái sử dụng tốt hơn.
