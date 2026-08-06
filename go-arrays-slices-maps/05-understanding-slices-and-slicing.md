# Bài 01: Tìm Hiểu Về Slices (Lát Cắt) & Cú Pháp Trích Xuất (Slicing)

> [!NOTE]
> *Chúng ta sẽ bắt đầu học về một trong những cấu trúc dữ liệu quan trọng nhất của Go: **Slice (Lát cắt)**. Tìm hiểu cách trích xuất (slicing) các phần của mảng, cách lược bỏ chỉ mục bắt đầu/kết thúc, và cách tạo slice từ một slice khác.*

---

## ✂️ 1. Khái Niệm Slicing (Trích Xuất)

**Slicing** là thao tác lấy ra một phần dữ liệu liên tiếp từ một mảng (Array) hoặc một Slice khác để tạo thành một Slice mới.

Cú pháp cơ bản:
```go
slice := array[start:end]
```
Trong đó:
* **`start`**: Chỉ mục bắt đầu trích xuất (được **bao gồm** trong kết quả).
* **`end`**: Chỉ mục kết thúc trích xuất (bị **loại trừ** khỏi kết quả).

---

## 💡 2. Cú Pháp Lược Bỏ Chỉ Mục (Omitting Indices)

Go cung cấp cú pháp viết tắt tiện lợi bằng cách bỏ qua một trong hai chỉ mục:

### A. Bỏ qua chỉ mục bắt đầu (`[:end]`)
Go sẽ tự hiểu là bắt đầu trích xuất từ phần tử đầu tiên (chỉ mục `0`):
```go
prices := [4]float64{10.99, 9.99, 45.99, 20.0}
fmt.Println(prices[:3]) // In ra: [10.99, 9.99, 45.99]
```

### B. Bỏ qua chỉ mục kết thúc (`[start:]`)
Go sẽ trích xuất từ chỉ mục `start` chạy dài đến hết mảng:
```go
fmt.Println(prices[2:]) // In ra: [45.99, 20.0]
```

---

## 🔀 3. Tạo Slice Từ Một Slice Khác (Slice of Slice)

Một tính năng cực kỳ linh hoạt trong Go là bạn hoàn toàn có thể trích xuất lát cắt trên một Slice đã tồn tại trước đó:

```go
// 1. Trích xuất từ mảng gốc tạo thành Slice featuredPrices
featuredPrices := prices[1:] // [9.99, 45.99, 20.0]

// 2. Tiếp tục trích xuất trên featuredPrices tạo thành highlightedPrices
highlightedPrices := featuredPrices[:1] // [9.99]
```

*Giải thích:* `highlightedPrices` lấy từ đầu của `featuredPrices` (là phần tử `9.99`) đến trước phần tử có chỉ mục `1` (loại trừ `45.99`).

---

## ⚠️ 4. Các Giới Hạn Cần Lưu Ý

1. **Không hỗ trợ chỉ mục âm (No negative indexing):** Go không cho phép bạn lấy phần tử cuối bằng chỉ mục âm (ví dụ: `prices[-1]`) giống như Python hay JavaScript. Viết như vậy sẽ bị báo lỗi cú pháp ngay lập tức.
2. **Lỗi vượt quá phạm vi (Out of bounds):** Bạn không thể đặt chỉ mục kết thúc vượt quá độ dài của mảng gốc.
   * Với mảng kích thước `4`, chỉ mục kết thúc lớn nhất được phép viết là `4` (vì phần tử cuối cùng có chỉ mục `3` được bao gồm, và `4` bị loại trừ).
   * Viết `prices[:5]` sẽ gây lỗi crash chương trình lúc chạy: `panic: runtime error: slice bounds out of range`.
