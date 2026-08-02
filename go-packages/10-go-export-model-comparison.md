# Bài 10: Cơ Chế Export Của Go Và So Sánh Với Các Ngôn Ngữ Khác

> [!NOTE]
> *Để hiểu sâu sắc về thiết kế độc đáo của Go, chúng ta sẽ liên hệ cơ chế xuất khẩu (Export) của nó với các gói thư viện chuẩn (Standard Library) và so sánh trực tiếp với mô hình từ khóa của các ngôn ngữ lập trình phổ biến khác như JavaScript.*

---

## 🔍 1. Liên Hệ Với Thư Viện Chuẩn (Go Standard Library)

Hãy nhớ lại các đoạn mã chúng ta đã viết từ những bài học đầu tiên:
* **`fmt.Println()`**: Tại sao chữ `P` trong `Println` lại viết hoa?
* **`fmt.Scan()`**: Tại sao chữ `S` trong `Scan` lại viết hoa?
* **`strconv.ParseFloat()`**: Tại sao chữ `P` và `F` lại viết hoa?

**Câu trả lời:** 
Bởi vì `fmt` hay `strconv` là các Package độc lập được viết bởi nhóm phát triển Go. Để các lập trình viên bên ngoài (như chúng ta trong package `main`) có thể import và sử dụng được các hàm này, bắt buộc tên hàm phải được định nghĩa viết hoa chữ cái đầu. Nếu họ định nghĩa là `fmt.println()`, trình biên dịch sẽ báo lỗi ngay lập tức.

---

## ⚖️ 2. So Sánh Mô Hình Export: Go vs JavaScript (ES6)

Để thấy rõ sự khác biệt trong tư duy thiết kế, hãy so sánh Go với JavaScript/TypeScript:

### Trong JavaScript/TypeScript (Dùng từ khóa tường minh)
Bạn cần thêm từ khóa `export` trước hàm hoặc biến muốn công khai:
```javascript
// fileops.js
export function writeFloatToFile(value, fileName) {
    // ...
}
```
Và ở file khác, bạn import cụ thể hàm đó:
```javascript
import { writeFloatToFile } from './fileops.js';
```

### Trong Go (Dựa hoàn toàn vào Casing)
Không cần bất kỳ từ khóa nào, chỉ cần viết hoa chữ cái đầu tiên:
```go
// fileops.go
func WriteFloatToFile(value float64, fileName string) {
    // ...
}
```
Và sử dụng tiền tố gói để gọi trực tiếp sau khi import thư mục:
```go
import "example.com/bank/fileops"
// ...
fileops.WriteFloatToFile(...)
```

---

## 🧠 3. Ưu Điểm Của Mô Hình Casing Trong Go

1. **Hiển thị API rõ ràng**: Khi nhìn vào mã nguồn của một gói, bạn biết ngay thực thể nào là một phần của giao diện công khai (Public API) để tích hợp, thực thể nào là chi tiết triển khai nội bộ (Private Implementation) chỉ bằng cách nhìn vào chữ viết hoa/thường.
2. **Loại bỏ sự rườm rà**: Giúp mã nguồn ngắn gọn hơn nhờ lược bỏ các từ khóa lặp đi lặp lại như `export`, `public`, `private` ở đầu mỗi khai báo.
3. **Đồng nhất hóa phong cách lập trình**: Mọi lập trình viên Go trên thế giới đều tuân theo cùng một tiêu chuẩn viết hoa tên hàm công khai, giúp code dễ đọc và chia sẻ hơn.
