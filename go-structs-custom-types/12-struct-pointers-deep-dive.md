# Bài 12: Con Trỏ Struct & Sự Khác Biệt Giữa Giá Trị vs Con Trỏ

> [!NOTE]
> *Tóm tắt ngắn gọn và trực quan về Con trỏ Struct, phân biệt giữa dạng Giá trị (Value) vs dạng Con trỏ (Pointer) trong thực tế.*

---

## 🏡 1. Ẩn Dụ Thực Tế: Ngôi Nhà (Value) vs Địa Chỉ Nhà (Pointer)

Hãy tưởng tượng Struct `User` giống như một **Ngôi nhà**:

* **Kiểu Giá Trị (`User`):** Tương đương với **Ngôi nhà thực tế**. Khi bạn truyền nó vào hàm hoặc gán cho biến khác, Go sẽ **sao chép xây thêm một ngôi nhà bản sao giống hệt**.
  * Sửa ngôi nhà bản sao (trong hàm/method) thì **ngôi nhà gốc không đổi**.
  * Việc sao chép liên tục gây tốn tài nguyên bộ nhớ.
* **Kiểu Con Trỏ (`*User`):** Tương đương với **Mảnh giấy ghi địa chỉ** của ngôi nhà đó (ví dụ: *"123 Go Street"*). Mảnh giấy này cực kỳ nhẹ (8 bytes).
  * Truyền địa chỉ cho hàm giúp hàm tìm đến và **sửa đổi trực tiếp trên ngôi nhà gốc**.

---

## ⚙️ 2. Khởi Tạo Con Trỏ Struct: `&User{...}`

Khi bạn viết `&User{...}`: Go sẽ tạo một thực thể mới và trả về **địa chỉ ô nhớ** (`*User`) của nó.

```go
func newUser(firstName, lastName, birthdate string) *User {
	return &User{ // Tạo và trả về con trỏ (địa chỉ)
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```
* **Escape Analysis (An toàn tuyệt đối):** Trong Go, việc trả về địa chỉ của struct cục bộ từ hàm là hoàn toàn an toàn. Trình biên dịch sẽ tự động đưa struct này lên bộ nhớ **Heap** để nó tiếp tục tồn tại kể cả sau khi hàm kết thúc.

---

## 👥 3. Khởi Tạo Nhiều Đối Tượng Độc Lập

Khi bạn tạo nhiều đối tượng từ cùng một struct blueprint:
```go
appUser1 := newUser("John", "Doe", "01/01/1990")
appUser2 := newUser("Jane", "Smith", "02/02/1995")
```
* Chúng là **hai ngôi nhà hoàn toàn riêng biệt** nằm ở hai địa chỉ ô nhớ khác nhau.
* Thay đổi dữ liệu của `appUser1` không bao giờ ảnh hưởng tới `appUser2`.

---

## ⚖️ 4. Khi Nào Dùng Con Trỏ (`*User`) vs Giá Trị (`User`)?

| Dùng Con Trỏ (`*User`) | Dùng Giá Trị (`User`) |
| :--- | :--- |
| **1. Cần sửa đổi dữ liệu gốc:** Method cần thay đổi các trường dữ liệu bên trong Struct (Pointer Receivers). | **1. Struct nhỏ & Chỉ đọc (Read-only):** Không cần thay đổi dữ liệu bên trong Struct. |
| **2. Tối ưu bộ nhớ:** Tránh sao chép Struct lớn có nhiều thuộc tính khi truyền qua lại giữa các hàm. | **2. Không chia sẻ trạng thái:** Dữ liệu mang tính chất tạm thời, độc lập và bất biến. |
| **3. Cần giá trị `nil` (Null):** Biểu diễn trạng thái "Không có dữ liệu" (Ví dụ: Không tìm thấy user -> trả về `nil`). | |
| **4. Chia sẻ tài nguyên dùng chung:** Các dịch vụ dùng chung (như kết nối Database, Logger) cần tất cả mọi nơi trỏ chung về 1 đối tượng gốc. | |
