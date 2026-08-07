# Bài 07: Dynamic Arrays (Mảng Động) Với Slices & Hàm `append()`

> [!NOTE]
> *Học cách tạo mảng động (Dynamic Arrays) trong Go thông qua Slice, hiểu cơ chế hoạt động của hàm `append()` để thêm phần tử, và phương pháp xóa phần tử khỏi Slice.*

---

## 🧩 1. Tại Sao Lại Cần Dynamic Arrays (Mảng Động)?

- **Hạn chế của Array (Mảng cố định):** Khi khai báo mảng cố định (ví dụ `[4]float64`), bạn phải xác định chính xác số lượng phần tử ngay lúc viết code. Không thể thêm phần tử thứ 5 vào mảng 4 phần tử.
- **Thực tế lập trình:** Nhiều trường hợp chúng ta không thể biết trước số lượng phần tử (ví dụ: lấy danh sách sản phẩm/giá từ cơ sở dữ liệu, nhận dữ liệu người dùng nhập,...).
- **Giải pháp của Go:** Sử dụng **Slice** làm **Dynamic Array** (Mảng động có thể co giãn kích thước).

---

## 🛠️ 2. Khai Báo Slice Trực Tiếp (Tạo Mảng Động)

Cú pháp khai báo Slice trực tiếp (không cần ghi kích thước trong cặp ngoặc vuông `[]`):

```go
package main

import "fmt"

func main() {
    // Không ghi số lượng trong [] -> Go sẽ tự tạo Slice và Array ẩn bên dưới
    prices := []float64{10.99, 8.99}

    // Truy cập phần tử như mảng bình thường
    fmt.Println(prices[1]) // Output: 8.99

    // Thay đổi phần tử đã có
    prices[1] = 9.99
    fmt.Println(prices) // Output: [10.99 9.99]
}
```

> [!WARNING]
> **Vẫn không thể gán trực tiếp vào chỉ mục chưa tồn tại!**  
> Việc gán `prices[2] = 5.99` khi Slice chỉ có 2 phần tử sẽ gây lỗi crash: `panic: runtime error: index out of range [2] with length 2`.

---

## ➕ 3. Thêm Phần Tử Mới Với Hàm Built-in `append()`

Để thêm phần tử vào Slice, Go cung cấp hàm tích hợp sẵn **`append()`**.

### Cú pháp:
```go
newSlice := append(originalSlice, value1, value2, ...)
```

### Cơ chế hoạt động của `append()`:
1. `append()` **KHÔNG** làm thay đổi trực tiếp Slice ban đầu.
2. Nó trả về một **Slice mới** chứa tất cả phần tử cũ cộng thêm phần tử mới vừa thêm vào.

> [!IMPORTANT]
> ### 🧠 Khi nào Go mới cấp phát vùng nhớ MỚI cho mảng?
> Khi thực hiện `append()`, Go xử lý theo 2 trường hợp:
> 
> - **Trường hợp 1: Sức chứa (`cap`) còn ĐỦ chỗ:**  
>   Go **KHÔNG tạo mảng mới**. Go sẽ ghi phần tử mới vào ô nhớ trống tiếp theo của mảng gốc hiện tại. Slice cũ và Slice mới vẫn **dùng chung mảng gốc**.
> 
> - **Trường hợp 2: Sức chứa (`cap`) đã HẾT chỗ:**  
>   Go sẽ **cấp phát một vùng nhớ MỚI chứa mảng MỚI** có sức chứa lớn hơn (thường gấp đôi `cap` cũ), tự động **sao chép (copy)** toàn bộ dữ liệu từ mảng cũ sang mảng mới, rồi chèn phần tử mới vào. Mảng cũ không còn dùng tới sẽ được Garbage Collector dọn dẹp.


### Ví dụ 1: Tạo Slice mới với `append()`
```go
prices := []float64{10.99, 9.99}

// Thêm 5.99 vào prices, kết quả trả về một slice mới
updatedPrices := append(prices, 5.99)

fmt.Println(prices)        // Output: [10.99 9.99] (Slice cũ giữ nguyên)
fmt.Println(updatedPrices) // Output: [10.99 9.99 5.99] (Slice mới)
```

### Ví dụ 2: Gán đè lại chính Slice ban đầu (Cách dùng phổ biến nhất)
Nếu bạn muốn Slice cũ cập nhật thêm phần tử mới, hãy gán ngược lại kết quả của `append` cho chính biến đó:

```go
prices := []float64{10.99, 9.99}

// Gán đè lại biến prices
prices = append(prices, 5.99)

fmt.Println(prices) // Output: [10.99 9.99 5.99]
```

---

## 🗑️ 4. Xóa Phần Tử Khỏi Slice (Removing Elements)

Go **không cung cấp** sẵn hàm `remove()` riêng biệt vì tính năng trích xuất (Slicing) đã quá mạnh mẽ để làm điều này.

### Cách xóa phần tử đầu tiên:
Muốn xóa phần tử đầu tiên (chỉ mục `0`), ta chỉ cần trích xuất slice từ chỉ mục `1` cho đến hết:

```go
prices := []float64{10.99, 9.99, 5.99}

// Lấy từ phần tử có chỉ mục 1 đến hết -> Loại bỏ phần tử đầu tiên (10.99)
discountedPrices := prices[1:]

fmt.Println(discountedPrices) // Output: [9.99 5.99]
```

### Cách xóa phần tử bất kỳ ở giữa:
Có thể kết hợp cú pháp Slicing `[ :index]` và `[index+1: ]` cùng với hàm `append()`:

```go
prices := []float64{10.99, 9.99, 45.99, 5.99}
// Giả sử muốn xóa phần tử 45.99 (chỉ mục 2)

prices = append(prices[:2], prices[3:]...)
fmt.Println(prices) // Output: [10.99 9.99 5.99]
```

---

## 📝 Tóm Tắt Bài Học

1. **Slice = Dynamic Array:** Không điền độ dài trong `[]` khi khai báo (`[]float64{...}`).
2. **Hàm `append()`:** Dùng để thêm phần tử vào Slice.
   - Thường gán đè lại biến gốc: `slice = append(slice, item)`.
   - Go tự động quản lý việc tạo mảng gốc mới đằng sau hậu trường khi vượt quá sức chứa.
3. **Thực tế lập trình:** Slice được sử dụng phổ biến hơn Array rất nhiều nhờ tính linh hoạt.
4. **Xóa phần tử:** Kết hợp cú pháp Slicing `[start:end]` để loại bỏ phần tử không mong muốn.
