# Bài 03: Các Cách Khai Báo Array Và Truy Cập Phần Tử

> [!NOTE]
> *Trong bài học này, chúng ta sẽ mở rộng kiến thức về Array bằng cách tìm hiểu cách khai báo mảng không khởi tạo ngay (sử dụng từ khóa `var`), khám phá khái niệm giá trị mặc định (Zero value) của mảng, và cách truy xuất, cập nhật giá trị của từng phần tử đơn lẻ qua chỉ mục (Index).*

---

## 📍 1. Khai Báo Mảng Không Khởi Tạo Ngay (Dùng `var`)

Ngoài cách sử dụng toán tử gán nhanh `:=`, bạn có thể khai báo một biến mảng trước bằng từ khóa `var` và định nghĩa kiểu mảng mà không cần gán giá trị ngay lập tức.

### Cú pháp:
```go
var tên_biến [số_phần_tử]kiểu_dữ_liệu
```

Ví dụ:
```go
var productNames [4]string
```

### 💡 Khái niệm Zero Value của Array:
Khi khai báo như trên mà chưa gán giá trị, Go sẽ tự động khởi tạo mảng đó với các **giá trị mặc định (Zero values)** của kiểu dữ liệu tương ứng:
* Đối với `string`: Mảng chứa các chuỗi rỗng `""`.
* Đối với số (`int`, `float64`): Mảng chứa các số `0` hoặc `0.0`.
* Đối với `bool`: Mảng chứa các giá trị `false`.

Khi in mảng `productNames` vừa khai báo ở trên, chương trình sẽ không bị lỗi mà in ra một mảng chứa 4 khoảng trống (các chuỗi rỗng): `[   ]`.

---

## 🛠️ 2. Gán Giá Trị Cho Mảng Và Gán Một Phần

Chúng ta có thể gán toàn bộ mảng hoặc chỉ khởi tạo một vài phần tử đầu tiên:
```go
productNames = [4]string{"A Book"}
```
Trong ví dụ này:
* Phần tử đầu tiên (chỉ mục `0`) sẽ có giá trị là `"A Book"`.
* Các phần tử còn lại (chỉ mục `1`, `2`, `3`) vẫn giữ nguyên giá trị mặc định là chuỗi rỗng `""`.

---

## 🎯 3. Truy Cập Và Cập Nhật Phần Tử Qua Chỉ Mục (Index)

Trong Go, mảng sử dụng cơ chế **Zero-indexed** (chỉ mục bắt đầu từ `0`).

### A. Đọc giá trị của một phần tử (Read Access)
Sử dụng cú pháp ngoặc vuông `[index]` ngay sau tên biến mảng.
* Để lấy phần tử thứ **nhất**: dùng `mảng[0]`.
* Để lấy phần tử thứ **ba**: dùng `mảng[2]`.

Ví dụ:
```go
fmt.Println(prices[2]) // Truy xuất phần tử thứ 3 của mảng prices
```

### B. Thay đổi giá trị của một phần tử (Write Access)
Ta cũng dùng chỉ mục để gán giá trị mới cho một vị trí cụ thể trong mảng:
```go
productNames[2] = "A Carpet" // Gán giá trị "A Carpet" vào vị trí thứ 3 (index 2)
```

---

## 💻 4. Mã Nguồn Minh Họa (`lists.go`)

Dưới đây là mã nguồn hoàn chỉnh của bài học:

```go
package main

import "fmt"

func main() {
	// 1. Khai báo và khởi tạo trực tiếp
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices) // Output: [10.99 9.99 45.99 20]

	// 2. Truy xuất phần tử thứ 3 (chỉ mục 2)
	fmt.Println(prices[2]) // Output: 45.99

	// 3. Khai báo mảng bằng var (giá trị mặc định)
	var productNames [4]string
	
	// Khởi tạo mảng có 1 phần tử, các phần tử còn lại là chuỗi rỗng
	productNames = [4]string{"A Book"}
	
	// 4. Cập nhật phần tử ở chỉ mục 2 (vị trí thứ 3)
	productNames[2] = "A Carpet"
	
	// In toàn bộ mảng productNames ra màn hình
	fmt.Println(productNames) // Output: [A Book  A Carpet ]
}
```
