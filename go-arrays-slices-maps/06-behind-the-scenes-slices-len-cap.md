# Bài 06: Cơ Chế Hoạt Động Của Slice: Bộ Nhớ, Len & Cap (Length & Capacity)

> [!NOTE]
> *Đi sâu vào bên trong hậu trường của Go để hiểu rõ bản chất mối quan hệ giữa **Slice** và **Array**: Slice là cửa sổ (window)/tham chiếu (reference) đến mảng bên dưới, cơ chế chia sẻ bộ nhớ, cùng hai khái niệm cực kỳ quan trọng: **Độ dài (`len`)** và **Sức chứa (`cap`)**.*

---

## 🪟 1. Slice Là "Cửa Sổ" Chiếu Vào Mảng Gốc (Reference Window)

Bản chất trong Go:
- Khi tạo một mảng (Array), dữ liệu thực sự được lưu trữ liền kề trong vùng nhớ.
- Khi tạo một Slice dựa trên mảng đó, Go **KHÔNG** sao chép dữ liệu sang vùng nhớ mới. Slice chỉ đóng vai trò là một **tham chiếu** (hay một "cửa sổ" xem dữ liệu) chiếu tới vùng nhớ mảng gốc.

### ⚠️ Hệ quả: Sửa dữ liệu Slice làm đổi dữ liệu Mảng gốc!

Do dùng chung vùng nhớ, nếu bạn thay đổi giá trị của một phần tử thông qua Slice, giá trị trong mảng gốc cũng bị thay đổi tương ứng.

```go
package main

import "fmt"

func main() {
    prices := [4]float64{10.99, 9.99, 45.99, 20.0}

    // featuredPrices lấy từ chỉ mục 1 đến hết: [9.99, 45.99, 20.0]
    featuredPrices := prices[1:]

    // Thay đổi phần tử đầu tiên của featuredPrices (chính là prices[1])
    featuredPrices[0] = 199.99

    fmt.Println(prices) // Kết quả: [10.99, 199.99, 45.99, 20.0]
}
```

### 💡 Lợi ích về mặt hiệu năng
Việc không sao chép lại mảng giúp Slice hoạt động cực kỳ nhẹ và tối ưu bộ nhớ (Memory Efficient), vì thao tác tạo Slice không tiêu tốn bộ nhớ bổ sung cho dữ liệu phần tử.

---

## 📏 2. Tìm Hiểu Về Độ Dài `len()` và Sức Chứa `cap()`

Mỗi Slice trong Go quản lý 3 thành phần metadata:
1. **Con trỏ (Pointer):** Trỏ đến phần tử đầu tiên của mảng gốc mà Slice truy cập.
2. **Độ dài (`len` - Length):** Số lượng phần tử hiện có trong Slice.
3. **Sức chứa (`cap` - Capacity):** Số lượng phần tử tối đa mà Slice có thể mở rộng tính từ vị trí bắt đầu của Slice tới cuối mảng gốc.

Hai hàm có sẵn (built-in) không cần import thư viện:
- `len(slice)`: Trả về độ dài (Length).
- `cap(slice)`: Trả về sức chứa (Capacity).

---

## 🔍 3. Ví Dụ Phân Tích Sự Khác Biệt Giữa `len` và `cap`

Xét đoạn mã ví dụ sau:

```go
package main

import "fmt"

func main() {
    prices := [4]float64{10.99, 9.99, 45.99, 20.0}

    featuredPrices := prices[1:] // [9.99, 45.99, 20.0]
    highlightedPrices := featuredPrices[:1] // [9.99]

    fmt.Println(len(featuredPrices), cap(featuredPrices))       // Output: 3 3
    fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // Output: 1 3
}
```

### Phân tích chi tiết:
1. **`prices` (Array gốc):** Có 4 phần tử `[10.99, 9.99, 45.99, 20.0]`.
2. **`featuredPrices` (`prices[1:]`):**
   - Bắt đầu từ chỉ mục `1` (`9.99`) đến hết mảng.
   - `len` = 3 (chứa `9.99`, `45.99`, `20.0`).
   - `cap` = 3 (từ vị trí `9.99` đến phần tử cuối cùng của `prices` có 3 vị trí).
3. **`highlightedPrices` (`featuredPrices[:1]`):**
   - Bắt đầu từ vị trí `9.99` và chỉ chọn lấy 1 phần tử.
   - `len` = 1 (vì hiện tại slice chỉ lấy hiển thị đúng 1 phần tử: `[9.99]`).
   - `cap` = 3 (vì từ điểm bắt đầu `9.99` kéo dài **cho tới hết mảng gốc** vẫn còn đúng 3 ô nhớ `[9.99, 45.99, 20.0]`).

> [!TIP]
> ### 💡 Sơ đồ giải thích trực quan tại sao `cap = 3`:
> 
> ```text
> Chỉ mục mảng gốc:   [0]      [1]      [2]      [3]
> Dữ liệu mảng gốc:  10.99  | 9.99  |  45.99    20.00
>                    (bỏ)   └──┬──┘   └──────┬──────┘
>                              │             │
>                     len = 1 (chỉ lấy ô này) │
>                              │             │
>                              └─────────────┴───────► cap = 3 (vẫn còn 3 ô tính tới hết mảng gốc)
> ```
> 
> **Công thức tính:** $\text{cap} = \text{Số phần tử từ BẮT ĐẦU Slice} \rightarrow \text{CUỐI MẢNG GỐC}$.
> - `len`: Số phần tử bạn đang **muốn hiển thị/truy cập** tại thời điểm này.
> - `cap`: Số phần tử **khả dụng tối đa** đằng sau vị trí bắt đầu mà slice có thể mở rộng tới được.

---


## ➡️ 4. Quy Tắc Vàng: Chỉ Mở Rộng Về Phía Phải (To The Right)

Go có một quy tắc đặc biệt về việc Reslice (trích xuất lại):
> **Bạn có thể trích xuất mở rộng thêm phần tử về phía PHẢI (phía cuối mảng gốc), nhưng KHÔNG THỂ mở rộng lùi về phía TRÁI (phần tử đã bị bỏ qua ở đầu).**

### Ví dụ về Mở rộng lại Slice (Reslicing):

Dù `highlightedPrices` chỉ có độ dài là `1` (`[9.99]`), nhưng nhờ `cap` = 3, bạn hoàn toàn có thể reslice chính `highlightedPrices` để lấy thêm phần tử về bên phải:

```go
// Ban đầu: highlightedPrices có len=1, cap=3 ([9.99])
highlightedPrices = highlightedPrices[:3] 

fmt.Println(highlightedPrices)                 // Output: [9.99, 45.99, 20.0]
fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // Output: 3 3
```

### ❌ Những gì đã bỏ qua ở bên trái sẽ mất hẳn:
Do `featuredPrices` bắt đầu từ chỉ mục `1` của `prices` (đã bỏ qua phần tử `10.99` ở chỉ mục `0`), bất kỳ Slice nào tạo từ `featuredPrices` hoặc `highlightedPrices` cũng **không bao giờ** truy cập ngược lại phần tử `10.99` được nữa. Phần tử bên trái đó coi như đã vô hình đối với Slice này.

---

## 📝 Tóm Tắt Bài Học

1. **Slice là Reference:** Sửa phần tử của Slice sẽ làm thay đổi mảng gốc.
2. **Không tốn bộ nhớ:** Slice không copy mảng, rất tối ưu hiệu năng.
3. **`len()` vs `cap()`:** 
   - `len`: Số phần tử hiện có trong Slice.
   - `cap`: Số phần tử tối đa có thể với tới tính từ đầu Slice đến cuối mảng gốc.
4. **Reslice vươn về bên phải:** Nhờ `cap`, một Slice có `len` ngắn vẫn có thể được trích xuất lại (reslice) để lấy thêm các phần tử ở bên phải trong phạm vi `cap`.
