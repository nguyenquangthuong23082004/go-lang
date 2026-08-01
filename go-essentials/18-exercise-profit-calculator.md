# Bài 17: Bài Tập Thực Hành: Xây Dựng Bộ Tính Toán Lợi Nhuận (Profit Calculator)

> [!IMPORTANT]
> *Để làm chủ các kiến thức vừa học về biến, hằng số, kiểu dữ liệu, toán tử, ép kiểu, và hàm nhập/xuất dữ liệu, bạn sẽ tự thực hành xây dựng một ứng dụng tính toán thực tế.*

---

## 📋 1. Yêu Cầu Đề Bài

Hãy viết một chương trình dòng lệnh bằng Go đóng vai trò là một **Bộ tính toán lợi nhuận (Profit Calculator)**. Chương trình cần thực hiện các công việc sau:

### 📥 Bước 1: Nhận dữ liệu nhập từ người dùng
Yêu cầu người dùng nhập lần lượt 3 thông số (kiểu số thực `float64`):
1. Doanh thu (**Revenue**)
2. Chi phí (**Expenses**)
3. Tỷ lệ thuế (**Tax Rate**) - tính theo phần trăm (Ví dụ nhập `20` tương đương `20%`).

### 🧮 Bước 2: Thực hiện các phép tính tài chính
* **Lợi nhuận trước thuế (EBT - Earnings Before Tax):**
  $$\text{EBT} = \text{Revenue} - \text{Expenses}$$
* **Lợi nhuận sau thuế / Lợi nhuận ròng (Profit / EAT - Earnings After Tax):**
  $$\text{Profit} = \text{EBT} \times \left(1 - \frac{\text{Tax Rate}}{100}\right)$$
* **Tỷ số giữa Lợi nhuận trước thuế và sau thuế (Ratio):**
  $$\text{Ratio} = \frac{\text{EBT}}{\text{Profit}}$$

### 📺 Bước 3: In các kết quả ra terminal
In cả 3 giá trị vừa tính được (`EBT`, `Profit`, và `Ratio`) ra màn hình terminal kèm theo lời mô tả rõ ràng.

---

## 🛠️ 2. Hướng Dẫn Các Bước Thực Hiện Locally

1. **Tạo thư mục dự án mới:** Tạo thư mục `03-profit-calculator` nằm trong `/home/thuong/Desktop/go-learning/learning/`.
2. **Khởi tạo Module:** Di chuyển vào thư mục đó trên terminal và khởi tạo Go Module:
   ```bash
   go mod init profit-calculator
   ```
3. **Tạo file code:** Tạo file `main.go` và bắt đầu lập trình.
4. **Kiểm tra chạy thử:** Sử dụng lệnh `go run .` để chạy và nhập thử các số liệu thực tế.

> [!TIP]
> Hãy tự mình suy nghĩ cấu trúc biến, kiểu dữ liệu (tường minh hay tự suy luận), và cách in dữ liệu thân thiện nhất. Lời giải chi tiết và phân tích sâu sẽ được cung cấp ở bài học kế tiếp!
