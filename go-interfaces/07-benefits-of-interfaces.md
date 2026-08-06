# Bài 07: Điểm Lợi Của Interface So Với Cách Code Cũ

> [!NOTE]
> *Tóm tắt nhanh gọn và trực quan sự khác biệt, cùng các lợi ích rõ rệt khi viết code bằng Interface so với cách lập trình thông thường.*

---

## ⚖️ 1. Bảng So Sánh Trực Quan

| Tiêu chí so sánh | Cách Code Cũ (Không dùng Interface) | Cách Code Mới (Dùng Interface) |
| :--- | :--- | :--- |
| **Trùng lặp code (DRY)** | Phải viết lặp lại các khối lệnh check lỗi (`if err != nil`) và các câu lệnh in log tương tự cho từng đối tượng riêng biệt (`Note`, `Todo`). | Gom toàn bộ logic check lỗi, in log thành một hàm dùng chung duy nhất (`saveData`). |
| **Khả năng mở rộng (Scalability)** | Khi thêm Struct mới (như `Memo`, `Reminder`), bắt buộc phải copy-paste thêm code xử lý hoặc tạo thêm hàm lưu riêng cho kiểu dữ liệu đó. | Chỉ cần Struct mới thỏa mãn phương thức của Interface là tự động tương thích với hàm dùng chung, hoàn toàn **không cần sửa hay viết thêm code xử lý lỗi**. |
| **Độ liên kết (Coupling)** | **Liên kết chặt (Tight Coupling):** Các hàm xử lý bị ràng buộc cứng vào một kiểu dữ liệu cụ thể (như `note.Note`). | **Liên kết lỏng (Loose Coupling):** Hàm chỉ quan tâm đến hành vi (đối tượng có phương thức `Save()` hay không), không quan tâm đối tượng đó cụ thể là gì. |
| **Bảo trì (Maintenance)** | Khi muốn thay đổi logic (ví dụ: đổi câu thông báo, ghi log lỗi vào file thay vì in màn hình), phải đi tìm và sửa ở nhiều vị trí trong file `main.go`. | Chỉ cần chỉnh sửa tại **một nơi duy nhất** bên trong hàm dùng chung (`saveData`). |

---

## 🎯 2. Phân Tích Hai Điểm Lợi Cốt Lõi

### A. Triệt tiêu sự trùng lặp (DRY - Don't Repeat Yourself)
Khi không dùng Interface, logic kiểm tra lỗi ghi file được lặp lại:
```go
// Lặp lần 1 cho Todo
err = todo.Save()
if err != nil { ... }

// Lặp lần 2 cho Note
err = userNote.Save()
if err != nil { ... }
```
Nếu gom vào một hàm nhận tham số Interface `saver`, ta chỉ cần viết logic kiểm tra lỗi và in thông báo **một lần duy nhất** trong hàm `saveData`. Khi cần nâng cấp logic lưu trữ (ví dụ: thêm ghi log, gửi cảnh báo khi lưu lỗi), ta chỉ việc sửa đúng hàm `saveData`.

### B. Nguyên lý Đóng/Mở (Open-Closed Principle)
Đây là lợi ích lớn nhất trong kiến trúc phần mềm chuyên nghiệp. Code của bạn nên **mở rộng để viết thêm tính năng mới**, nhưng **đóng đối với việc sửa đổi những đoạn code cũ đang chạy ổn định**:
* Giả sử dự án hoạt động ổn định và cần bổ sung thêm 5 tính năng tương tự (ví dụ: `Reminder`, `Journal`, `Task`, `Memo`, `Post`).
* Nếu dùng Interface: Bạn chỉ cần viết định nghĩa cho 5 Struct mới đó và các phương thức `Save()` đi kèm. Hàm `main` và hàm `saveData` cũ **không cần sửa bất cứ dòng code nào**, tránh được hoàn toàn rủi ro phát sinh lỗi (bug) ngoài ý muốn ở hệ thống hiện tại.
