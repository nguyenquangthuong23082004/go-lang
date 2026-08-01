# Bài 21: Tạo Chuỗi Định Dạng Mà Không In Ra (Sprint & Sprintf) trong Go

> [!NOTE]
> *Đôi khi bạn cần định dạng một chuỗi văn bản nhưng không muốn hiển thị nó ra terminal ngay lập tức (ví dụ: cần lưu vào biến để ghi file hoặc gửi qua mạng). Go cung cấp nhóm hàm `Sprint` giúp bạn giải quyết yêu cầu này.*

---

## 🆚 1. Nhóm Hàm `Print` vs Nhóm Hàm `Sprint`

Trong package `fmt`, sự khác biệt lớn nhất giữa hai nhóm hàm này là:

* **Nhóm `Print` (gồm `Print`, `Printf`, `Println`):** Luôn xuất kết quả trực tiếp ra terminal ngay khi được gọi. Chúng không trả về chuỗi văn bản để lưu trữ.
* **Nhóm `Sprint` (gồm `Sprint`, `Sprintf`, `Sprintln`):** Không in gì ra màn hình terminal. Thay vào đó, chúng **trả về (return) một chuỗi văn bản (`string`)** đã được xử lý xong để bạn gán vào biến.

---

## 🛠️ 2. Cách Sử Dụng `fmt.Sprintf()`

Hàm `fmt.Sprintf` hoạt động giống hệt `Printf` về mặt sử dụng các ký tự giữ chỗ (placeholders như `%v`, `%.2f`), nhưng nó trả về kết quả dưới dạng chuỗi:

```go
formattedFV := fmt.Sprintf("Danh nghĩa: %.2f\n", giaTriTuongLai)
formattedRFV := fmt.Sprintf("Thực tế (trừ lạm phát): %.2f\n", giaTriThucTe)
```

### Phân tích cơ chế:
1. Hàm `fmt.Sprintf` nhận định dạng, làm tròn giá trị số thực về 2 chữ số thập phân, chèn ký tự ngắt dòng `\n` ở cuối.
2. Trả về kết quả chuỗi hoàn chỉnh (Ví dụ: `"Danh nghĩa: 5871.21\n"`).
3. Gán chuỗi đó vào biến `formattedFV` bằng toán tử `:=`.

---

## 📺 3. Hiển Thị Các Biến Chuỗi Đã Định Dạng

Sau khi đã lưu các chuỗi hoàn thiện vào biến, chúng ta có thể dùng hàm in thô đơn giản `fmt.Print` để hiển thị chúng:

```go
fmt.Print(formattedFV, formattedRFV)
```

*Vì các biến chuỗi đã tích hợp sẵn ký tự ngắt dòng `\n`, giao diện hiển thị vẫn tự động xuống dòng một cách hoàn hảo.*

---

## 💡 4. Tại Sao Phương Pháp Này Lại Hữu Ích?

Việc tách biệt việc định dạng chuỗi ra khỏi việc in ấn giúp mã nguồn của bạn linh hoạt hơn:
* **Tái sử dụng dữ liệu:** Một chuỗi được định dạng có thể vừa in ra màn hình, vừa được ghi vào file lịch sử log, vừa gửi về server web.
* **Phân tách trách nhiệm (Separation of Concerns):** Phân chia rõ ràng phần xử lý logic hiển thị (Format) và phần giao tiếp thiết bị ngoại vi (I/O).
