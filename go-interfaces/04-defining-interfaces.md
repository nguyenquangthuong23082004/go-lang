# Bài 04: Cách Định Nghĩa Interface Trong Go & Quy Ước Đặt Tên

> [!NOTE]
> *Chúng ta sẽ tìm hiểu cách sử dụng từ khóa `type` kết hợp `interface` để định nghĩa một Giao diện (Interface) trong Go, đồng thời nắm vững quy ước đặt tên đuôi "-er" của Go.*

---

## 🛠️ 1. Cú Pháp Khai Báo Interface

Cú pháp khai báo Interface sử dụng kết hợp từ khóa `type` và `interface`:

```go
type saver interface {
	Save() error
}
```

Trong đó:
* **`saver`**: Tên của Interface (viết thường chữ cái đầu `s` vì chỉ sử dụng cục bộ trong file `main.go`).
* **`Save() error`**: Chữ ký phương thức (Method Signature). Nó định nghĩa một phương thức có tên `Save`, không nhận tham số đầu vào và trả về một kiểu `error`.

---

## 📜 2. Bản Chất "Hợp Đồng" Của Interface

Interface trong Go hoạt động tương tự như một bản hợp đồng cam kết về hành vi:

1. **Chỉ khai báo chữ ký, không có logic:** Bạn không thể viết thân hàm `{ ... }` cho các phương thức bên trong Interface. Nó chỉ quy định tên hàm, kiểu tham số đầu vào và kiểu trả về.
2. **Cam kết hành vi:** Bất kỳ kiểu dữ liệu nào (như các Struct `Note`, `Todo`) triển khai đầy đủ các phương thức được khai báo trong Interface thì được coi là **đáp ứng hợp đồng** của Interface đó.
3. **Đa phương thức:** Một Interface có thể chứa bao nhiêu chữ ký phương thức tùy ý:
   ```go
   type FileHandler interface {
       Save(string) error
       Delete() error
   }
   ```

---

## 🏷️ 3. Quy Ước Đặt Tên Đuôi "-er" Của Go (Idiomatic Go Naming)

Cộng đồng lập trình Go có một quy ước đặt tên rất phổ biến: **Nếu một Interface chỉ yêu cầu duy nhất một phương thức, tên của Interface đó sẽ bằng tên phương thức cộng thêm hậu tố "-er".**

Một số ví dụ tiêu biểu từ Thư viện chuẩn (Go Standard Library):
* Phương thức `Read()` -> Interface **`Reader`**
* Phương thức `Write()` -> Interface **`Writer`**
* Phương thức `Save()` -> Interface **`Saver`** (như ví dụ của chúng ta)
* Phương thức `String()` -> Interface **`Stringer`**

Việc tuân thủ quy ước này giúp code của bạn đồng bộ với phong cách thiết kế chung của hệ sinh thái Go.
