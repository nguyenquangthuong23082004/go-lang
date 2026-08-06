# Bài 04: Giới Thiệu Về Slices (Lát Cắt/Mảng Con) Trong Go

> [!NOTE]
> *Trong bài học này, chúng ta sẽ làm quen với khái niệm **Slice (Lát cắt)** - một trong những tính năng cốt lõi và được sử dụng nhiều nhất trong Go. Slice cho phép chúng ta tham chiếu đến một phần (mảng con) của một mảng có sẵn.*

---

## 🔍 1. Slice Là Gì?

Như đã biết, **Array** có kích thước cố định. Tuy nhiên, trong thực tế, nhiều lúc chúng ta chỉ cần làm việc với một **tập hợp con (subset)** của mảng, hoặc cần các cấu trúc dữ liệu có kích thước linh hoạt hơn.

**Slice (Lát cắt)** là một góc nhìn (view) trỏ đến một đoạn liên tục của một mảng nằm bên dưới. Nó không thực sự lưu trữ bất kỳ dữ liệu nào; nó chỉ mô tả một phần của một mảng có sẵn.

---

## 🛠️ 2. Cú Pháp Slicing (Cắt Lát) Mảng

Để tạo một Slice từ một mảng hiện có, chúng ta sử dụng cú pháp:

```go
slice_name := mang_goc[start_index:end_index]
```

### ⚠️ Quy tắc quan trọng về chỉ mục (Index rule):
* **`start_index` (Inclusive - Bao gồm)**: Phần tử tại vị trí bắt đầu sẽ nằm trong Slice.
* **`end_index` (Exclusive - Ngoại trừ)**: Phần tử tại vị trí kết thúc sẽ **KHÔNG** nằm trong Slice (chương trình sẽ dừng lại ngay trước index này).

Ví dụ:
```go
prices := [4]float64{10.99, 9.99, 45.99, 20.0}
featuredPrices := prices[1:3]
```

Trong ví dụ này:
* Mảng gốc `prices` có chỉ mục từ `0` đến `3`:
  * `prices[0] = 10.99`
  * `prices[1] = 9.99`
  * `prices[2] = 45.99`
  * `prices[3] = 20.0`
* Lệnh cắt lát `prices[1:3]` sẽ bắt đầu từ index `1` (giá trị `9.99`) và đi đến sát index `3` (chỉ lấy đến index `2` là `45.99`, bỏ qua phần tử tại index `3`).
* Kết quả là Slice `featuredPrices` sẽ chứa hai giá trị ở giữa: `[9.99 45.99]`.

---

## 💻 3. Mã Nguồn Minh Họa (`lists.go`)

```go
package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	fmt.Println(prices[2])

	var productNames [4]string
	productNames = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	// Tạo một Slice chứa các phần tử từ index 1 đến trước index 3
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices) // Output: [9.99 45.99]
}
```

---

## 💡 4. Tại Sao Slices Lại Quan Trọng?

Slices cực kỳ hữu ích và linh hoạt:
1. **Tiết kiệm bộ nhớ**: Khi tạo một Slice từ một mảng lớn, Go không sao chép dữ liệu của mảng đó ra vùng nhớ khác. Slice chỉ đơn giản là chứa một con trỏ trỏ đến vị trí của mảng gốc, giúp tối ưu hiệu năng.
2. **Nền tảng của Mảng Động**: Slice chính là cách Go triển khai mảng động (dynamic arrays) mà chúng ta có thể thêm/bớt phần tử một cách linh hoạt (chúng ta sẽ tìm hiểu thêm trong các bài tiếp theo).
