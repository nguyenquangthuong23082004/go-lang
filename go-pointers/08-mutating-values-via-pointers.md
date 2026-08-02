# Bài 08: Thay Đổi Trực Tiếp Giá Trị Bằng Con Trỏ (Mutation Via Pointers)

> [!IMPORTANT]
> *Một trong những công dụng lớn nhất của con trỏ là khả năng thay đổi trực tiếp giá trị của biến gốc nằm tại địa chỉ ô nhớ truyền vào, loại bỏ sự cần thiết của câu lệnh `return`.*

---

## 🛠️ 1. Cơ Chế Thay Thế Trực Tiếp Trên Ô Nhớ

Khi hàm nhận đối số con trỏ, chúng ta có thể ghi đè dữ liệu mới vào đúng ô nhớ đó bằng cú pháp gán giải tham chiếu:

```go
func editAgeToAdultYears(age *int) {
	*age = *age - 18 // Ghi đè giá trị mới vào ô nhớ mà age đang trỏ tới
}
```

### Phân tích hai vế của phép gán `=`:
* **`*age` ở vế phải**: Biểu thị hành động **đọc dữ liệu** từ địa chỉ ô nhớ (Lấy ra giá trị `32`).
* **`*age` ở vế trái**: Biểu thị hành động **ghi dữ liệu** vào địa chỉ ô nhớ (Ghi đè giá trị mới là `14`).

---

## 💻 2. Ví Dụ Minh Họa Trực Quan

Hãy xem luồng chạy trong file `pointers.go`:

```go
func main() {
	age := 32 // Biến gốc chứa 32

	agePointer := &age

	// Truyền con trỏ vào hàm thay đổi
	editAgeToAdultYears(agePointer)

	// In biến gốc age
	fmt.Println("Age (After Mutation):", age) // Kết quả in ra: 14!
}
```

* **Hiện tượng**: Mặc dù chúng ta in biến gốc `age` (chứ không phải `*agePointer`), giá trị của nó vẫn bị thay đổi từ `32` thành `14`.
* **Giải thích**: Hàm `editAgeToAdultYears` đã tìm đến địa chỉ của `age` và sửa trực tiếp dữ liệu tại đó. Do biến `age` dùng chung địa chỉ ô nhớ này, nó lập tức phản ánh giá trị mới.

---

## ⚠️ 3. Quy Tắc Đặt Tên Hàm Để Tránh Hiểu Lầm (Side Effects)

Việc tự ý thay đổi dữ liệu gốc trong hàm đôi khi là một bất ngờ không mong muốn (side effect) đối với các nhà phát triển khác cùng dự án.

* **Tên cũ (`getAdultYears`)**: Có nghĩa là "lấy ra số tuổi trưởng thành". Người dùng hàm sẽ kỳ vọng biến gốc của họ được giữ nguyên và hàm trả về kết quả mới.
* **Tên mới (`editAgeToAdultYears` / `updateAgeToAdultYears`)**: Rõ ràng và tường minh hơn. Từ `edit` hoặc `update` cảnh báo rằng hàm này sẽ **chỉnh sửa trực tiếp** biến truyền vào và không cần giá trị trả về.
