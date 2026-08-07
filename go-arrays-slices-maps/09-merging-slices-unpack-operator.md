# Bài 09: Toán Tử Unpack (Variadic Operator `...`) Khi Append Slice

> [!NOTE]
> *Học cú pháp đặc biệt `...` (Unpack / Variadic Expansion) để rã các phần tử của một Slice và nối (merge) 2 Slice lại với nhau bằng hàm `append()`.*

---

## 💥 1. Vấn Đề Khi Nối (Merge) 2 Slice

Hàm `append()` trong Go cho phép truyền nhiều phần tử đơn lẻ:
```go
prices = append(prices, 5.99, 12.99, 19.99) // Hợp lệ!
```

Tuy nhiên, nếu bạn có 2 Slice sẵn có và muốn nối Slice thứ 2 vào Slice thứ 1:

```go
prices := []float64{10.99, 8.99}
discountPrices := []float64{101.99, 80.99, 20.59}

// ❌ Bị báo lỗi biên dịch!
prices = append(prices, discountPrices) 
```

**Lý do:** 
- `prices` là một danh sách kiểu `[]float64` (mỗi phần tử là 1 số `float64`).
- `discountPrices` là cả một Slice (`[]float64`), không phải là một số `float64` đơn lẻ.
- Truyền `discountPrices` vào sẽ khiến Go hiểu bạn muốn chèn cả một Slice vào làm 1 phần tử của `prices` -> **Sai kiểu dữ liệu!**

---

## ✨ 2. Giải Pháp: Toán Tử Unpack (`...`)

Go cung cấp cú pháp **3 dấu chấm (`...`)** đặt ngay sau biến Slice thứ hai.

Cú pháp này hướng dẫn Go: **"Hãy mở bao bì (unpack) Slice này ra thành từng phần tử đơn lẻ phân tách bằng dấu phẩy và truyền lần lượt vào hàm `append()`."**

```go
package main

import "fmt"

func main() {
    prices := []float64{10.99, 8.99}
    discountPrices := []float64{101.99, 80.99, 20.59}

    // ✅ Thêm 3 dấu chấm (...) sau discountPrices
    prices = append(prices, discountPrices...)

    fmt.Println(prices) 
    // Output: [10.99 8.99 101.99 80.99 20.59]
}
```

---

## 💡 3. Tóm Tắt Quy Tắc

| Thao tác | Cú pháp | Giải thích |
| :--- | :--- | :--- |
| Thêm các giá trị đơn lẻ | `append(slice, val1, val2)` | Truyền trực tiếp các giá trị rời rạc. |
| Nối cả một Slice khác | `append(slice1, slice2...)` | **Phải có dấu `...`** để rã các phần tử của `slice2`. |
