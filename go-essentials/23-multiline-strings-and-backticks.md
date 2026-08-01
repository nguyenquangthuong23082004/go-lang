# Bài 22: Chuỗi Nhiều Dòng (Multiline Strings) & Ký Tự Backtick trong Go

> [!NOTE]
> *Khi cần xử lý các chuỗi văn bản dài gồm nhiều dòng (ví dụ: một bức thư, mã HTML, hoặc hướng dẫn sử dụng), Go cung cấp ký tự backtick (`` ` ``) để định nghĩa chuỗi thô nhiều dòng.*

---

## 🚫 1. Hạn Chế Của Chuỗi Nháy Kép (`""`)

Dấu nháy kép chỉ hỗ trợ khai báo chuỗi trên một dòng đơn lẻ. Nếu bạn nhấn Enter xuống dòng giữa chuỗi nháy kép:
```go
// LỖI: newline in string
fmt.Println("Xin chào
mọi người")
```
Trình biên dịch Go sẽ báo lỗi cú pháp ngay lập tức. Để xuống dòng trong chuỗi nháy kép, bạn bắt buộc phải dùng ký tự đặc biệt `\n` trên cùng một dòng.

---

## 💡 2. Giải Pháp: Chuỗi Ký Tự Thô Bằng Dấu Backtick (`` ` ``)

Dấu backtick (nằm ở góc trên bên trái bàn phím, dưới phím Esc) định nghĩa một **Raw String Literal** (chuỗi ký tự thô).

```go
fmt.Printf(`Danh nghĩa: %.2f
Thực tế (trừ lạm phát): %.2f
`, giaTriTuongLai, giaTriThucTe)
```

### ⚙️ Các đặc tính quan trọng của chuỗi dùng Backtick:

1. **Cho phép xuống dòng trực tiếp:** Bất kỳ thao tác xuống dòng nào bạn gõ trong VS Code sẽ được bảo toàn nguyên vẹn khi hiển thị ra terminal.
2. **Không thông dịch ký tự escape (`\n`):** Do là chuỗi "thô", ký tự `\n` sẽ hiển thị nguyên văn là chữ `\n` trên màn hình chứ không làm nhiệm vụ xuống dòng nữa. Khi dùng backtick, hãy nhấn Enter trực tiếp để xuống dòng thay vì viết `\n`.
   * *Ví dụ:* `` `Dòng 1\nDòng 2` `` sẽ in ra là `Dòng 1\nDòng 2`.

---

## 📝 3. Khi Nào Nên Sử Dụng?

* **Dấu nháy kép (`""`):** Thích hợp cho các dòng văn bản ngắn trên một dòng đơn lẻ, hoặc khi bạn cần sử dụng các ký tự escape đặc biệt như `\n`, `\t`.
* **Dấu backtick (`` ` ``):** Thích hợp khi viết các đoạn văn bản dài trải dài nhiều dòng, hoặc khi bạn viết các đoạn mã mẫu (như HTML, JSON) trực tiếp trong Go mà không muốn bị lỗi định dạng.
