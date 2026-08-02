# Bài 01: Giới Thiệu Về Con Trỏ (Pointers) Trong Go

> [!NOTE]
> *Chào mừng bạn đến với chương học về **Con Trỏ (Pointers)**. Đây là một tính năng đặc biệt mạnh mẽ trong Go, giúp quản lý ô nhớ hiệu quả và kiểm soát khả năng thay đổi dữ liệu của chương trình.*

---

## 📍 1. Con Trỏ (Pointer) Là Gì?

Thông thường, khi bạn khai báo một biến, ví dụ: `age := 22`, máy tính sẽ dành riêng một ô nhớ trong bộ nhớ RAM để lưu trữ giá trị `22` này.

* **Biến thông thường**: Lưu trữ trực tiếp **giá trị** dữ liệu (như `22`, `"hello"`, `true`).
* **Con trỏ (Pointer)**: Là một biến đặc biệt **không lưu giá trị trực tiếp**, mà nó lưu **địa chỉ ô nhớ** của một biến khác.

*Nói cách khác, con trỏ hoạt động giống như một chiếc thẻ ghi địa chỉ nhà, trỏ trực tiếp tới vị trí ngôi nhà thực tế trong RAM.*

---

## ❓ 2. Tại Sao Chúng Ta Cần Con Trỏ?

Có hai lý do cốt lõi khiến con trỏ trở thành một công cụ không thể thiếu:

### A. Hiệu năng bộ nhớ (Memory Efficiency)
Khi bạn truyền một tham số vào một hàm, mặc định Go sẽ tạo ra một **bản sao** của dữ liệu đó trong bộ nhớ.
* Nếu biến đó là một kiểu dữ liệu lớn (ví dụ: struct chứa thông tin hàng nghìn khách hàng), việc sao chép liên tục khi truyền qua lại giữa các hàm sẽ tốn rất nhiều RAM và làm chậm ứng dụng.
* Thay vào đó, nếu ta chỉ truyền một **con trỏ** (địa chỉ ô nhớ), Go chỉ truyền đi một địa chỉ có dung lượng cực kỳ nhỏ (64-bit trên máy tính hiện đại), giúp ứng dụng chạy nhanh vượt trội.

### B. Cho phép thay đổi giá trị gốc (Mutability)
Vì Go truyền tham số theo dạng bản sao (pass by value), mọi sự thay đổi giá trị của tham số bên trong hàm sẽ **không ảnh hưởng** đến biến gốc ở bên ngoài.
* Bằng cách truyền con trỏ (địa chỉ ô nhớ), hàm có thể truy cập trực tiếp vào ô nhớ gốc và chỉnh sửa giá trị của biến gốc.

---

## 🎯 3. Nội Dung Sẽ Học Trong Phần Này

Chúng ta sẽ khám phá con trỏ từng bước thông qua các chủ đề:
1. **Toán tử địa chỉ `&`**: Lấy địa chỉ ô nhớ của một biến.
2. **Kiểu dữ liệu con trỏ `*T`**: Khai báo biến con trỏ trỏ đến kiểu dữ liệu tương ứng.
3. **Toán tử giải tham chiếu (Dereferencing) `*`**: Lấy hoặc ghi đè giá trị tại địa chỉ ô nhớ mà con trỏ đang nắm giữ.
4. **Quy tắc sử dụng**: Khi nào thực sự nên sử dụng con trỏ và khi nào không nên dùng.
