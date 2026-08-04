# Bài 01: Giới Thiệu Về Interfaces (Giao Diện) Trong Go

> [!NOTE]
> *Sau khi tìm hiểu về Structs và Custom Types, chúng ta sẽ bước sang một khái niệm nâng cao và cực kỳ quan trọng khác trong Go giúp thiết kế hệ thống linh hoạt và lỏng lẻo (loose coupling): **Interfaces**.*

---

## 🔍 1. Interface Trong Go Là Gì?

Trong lập trình nói chung và Go nói riêng, một **Interface (Giao diện)** là một kiểu dữ liệu đặc biệt dùng để định nghĩa **hành vi (behaviors)** của đối tượng thay vì cấu trúc dữ liệu của nó:

* Một Interface chỉ chứa các **chữ ký phương thức (method signatures)** (tên phương thức, tham số đầu vào và kiểu trả về) mà không có bất kỳ phần mã thực thi (implementation) nào bên trong.
* Các kiểu dữ liệu tùy chỉnh khác (như Structs) sẽ hiện thực hóa (implement) các phương thức này bằng cách định nghĩa các hàm có receiver tương ứng.

---

## 🎯 2. Các Nội Dung Sẽ Học Trong Phần Này

Trong chương này, chúng ta sẽ đi sâu vào tìm hiểu cách Go giải quyết các bài toán thiết kế hướng đối tượng thông qua Interfaces:

### A. Khái niệm & Cơ chế hoạt động
* Tìm hiểu xem Interfaces là gì, cách hoạt động ngầm dưới nền của Go.
* Cách thiết kế các Struct tương thích tự động với Interface (**Implicit Implementation**).

### B. Tạo & Sử dụng Interface
* Cách định nghĩa một Interface mới và cách truyền đối tượng Interface làm đối số cho hàm.
* Cách áp dụng Interface để tạo ra tính **Đa hình (Polymorphism)**, giúp một hàm có thể chấp nhận nhiều loại Struct khác nhau miễn là chúng có cùng hành vi.

### C. Tìm hiểu về Empty Interface (`any`)
* Khám phá kiểu dữ liệu đặc biệt có thể chứa bất kỳ giá trị nào trong Go và các kỹ thuật ép kiểu (**Type Assertions**, **Type Switches**).
