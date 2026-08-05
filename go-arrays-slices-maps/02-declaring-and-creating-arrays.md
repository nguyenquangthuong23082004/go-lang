# Bài 02: Khai Báo Và Khởi Tạo Array (Mảng) Trong Go

> [!NOTE]
> *Trong bài học này, chúng ta sẽ bắt đầu làm quen với cách lưu trữ danh sách dữ liệu có cùng kiểu thông qua **Array (Mảng)**. Chúng ta sẽ tìm hiểu lý do tại sao Struct không phù hợp cho trường hợp này, cú pháp khai báo mảng tĩnh trong Go, và cách in mảng ra màn hình.*

---

## 📍 1. Tại Sao Struct Không Phù Hợp Để Lưu Danh Sách?

Như đã biết, `Struct` rất hữu dụng để gom các trường thông tin khác nhau của một đối tượng (ví dụ: Product gồm `Title`, `ID`, `Price`). 

Tuy nhiên, nếu ta cần lưu một danh sách các giá trị mô tả cùng một đối tượng (ví dụ: nhiệt độ hàng ngày trong tuần, danh sách giá sản phẩm để phân tích), việc dùng `Struct` sẽ rất cồng kềnh:

```go
type TemperatureData struct {
    Day1 float64
    Day2 float64
    Day3 float64
    // Phải khai báo thủ công từng ngày... rất dài dòng và không linh hoạt!
}
```

Để giải quyết vấn đề này, Go cung cấp **Array (Mảng)** - cấu trúc dữ liệu cho phép lưu trữ nhiều phần tử có cùng kiểu dữ liệu.

---

## 🛠️ 2. Cú Pháp Khai Báo Và Khởi Tạo Array

Để tạo một mảng trong Go, chúng ta sử dụng cú pháp với cặp dấu ngoặc vuông `[...]` đặt trước kiểu dữ liệu của các phần tử.

### Cú pháp khai báo và khởi tạo trực tiếp:
```go
biến := [số_phần_tử]kiểu_dữ_liệu{các_giá_trị_cách_nhau_bởi_dấu_phẩy}
```

Ví dụ cụ thể với mảng chứa 4 phần tử kiểu `float64` lưu giá sản phẩm:
```go
prices := [4]float64{10.99, 9.99, 45.99, 20.0}
```

* **`[4]`**: Xác định số lượng phần tử cố định của mảng.
* **`float64`**: Kiểu dữ liệu của tất cả phần tử trong mảng.
* **`{10.99, 9.99, 45.99, 20.0}`**: Khởi tạo giá trị trực tiếp cho từng phần tử.

---

## 💻 3. Minh Họa Code Thực Tế

Tập tin `lists.go` được thiết lập như sau:

```go
package main

import "fmt"

func main() {
    // Khai báo một mảng tĩnh chứa 4 phần tử float64
    prices := [4]float64{10.99, 9.99, 45.99, 20.0}

    // In toàn bộ mảng ra console
    fmt.Println(prices) // Output: [10.99 9.99 45.99 20]
}
```

---

## ⚠️ 4. Một Số Quy Tắc Và Giới Hạn Của Array Trong Go

* **Đồng nhất kiểu dữ liệu**: Tất cả phần tử trong mảng bắt buộc phải có cùng kiểu (ví dụ: cùng là `float64`, `int`, hoặc `string`). Bạn không thể trộn lẫn các kiểu dữ liệu khác nhau trong một mảng.
* **Kích thước cố định (Fixed Size)**: Độ dài của mảng là một phần cấu thành nên kiểu dữ liệu của nó (ví dụ: `[4]float64` là một kiểu dữ liệu hoàn toàn khác với `[5]float64`). Khi đã khai báo, bạn **không thể** thay đổi kích thước của mảng (không thể thêm hoặc bớt phần tử). Đây là một hạn chế lớn của Array trong Go, và chúng ta sẽ tìm hiểu cách khắc phục thông qua **Slices** ở các bài tiếp theo.
