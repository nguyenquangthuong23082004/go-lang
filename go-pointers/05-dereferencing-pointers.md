# Bài 05: Khái Niệm Giải Tham Chiếu (Dereferencing Pointers)

> [!IMPORTANT]
> *Để lấy giá trị thực sự nằm bên dưới địa chỉ ô nhớ mà con trỏ đang lưu trữ, chúng ta sử dụng toán tử giải tham chiếu (Dereference Operator) bằng cách thêm dấu sao `*` vào trước tên biến con trỏ.*

---

## 🔍 1. Hành Vi Khi In Con Trỏ Trực Tiếp

Nếu bạn in trực tiếp biến con trỏ ra màn hình:
```go
fmt.Println("Age (Pointer Address):", agePointer)
```
Kết quả hiển thị trên màn hình sẽ là một địa chỉ dạng số Hexadecimal của bộ nhớ RAM:
```text
Age (Pointer Address): 0xc000012088
```

---

## 🛠️ 2. Cách Lấy Giá Trị Bằng Giải Tham Chiếu (Dereferencing)

Để bảo Go tìm đến địa chỉ `0xc000012088` kia và lấy ra giá trị thực sự chứa trong đó (là số `32`), ta đặt dấu sao `*` trước tên con trỏ:
```go
fmt.Println("Age (Value from Pointer):", *agePointer)
```
Kết quả in ra:
```text
Age (Value from Pointer): 32
```

---

## 🔄 3. Tính Chất Đối Nghịch Giữa `&` và `*`

Toán tử `&` và `*` hoạt động như hai chiều ngược nhau:

```text
       Lấy địa chỉ (&)
Biến thường ----------> Con trỏ (Địa chỉ ô nhớ)

       Giải tham chiếu (*)
Con trỏ --------------> Giá trị gốc
```

* **Dùng `&`**: Đi từ **giá trị** sang **địa chỉ**.
* **Dùng `*`**: Đi từ **địa chỉ** sang **giá trị**.

---

## 💻 4. Mã Nguồn Cập Nhật Trong `pointers.go`

```go
package main

import "fmt"

func main() {
	age := 32 // Biến thông thường

	agePointer := &age // Lấy địa chỉ của age

	fmt.Println("Age (Pointer Address):", agePointer) // In địa chỉ: 0xc00...
	fmt.Println("Age (Value from Pointer):", *agePointer) // In giá trị: 32
}
```
