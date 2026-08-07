# Bài Tập Thực Hành: Arrays & Slices (Practice Tasks)

### Task 7 (Bonus): Tạo Mảng Động Chứa Struct & Thêm Phần Tử Mới
- **Yêu cầu:** Định nghĩa `Product` struct (gồm `id`, `title`, `price`), tạo mảng động chứa 2 sản phẩm và `append()` thêm sản phẩm thứ 3.
- **Code:**
```go
type Product struct {
    id    string
    title string
    price float64
}

// Tạo Slice chứa các Struct Product
products := []Product{
    {id: "p1", title: "Go Programming Book", price: 29.99},
    {id: "p2", title: "Ergonomic Keyboard", price: 99.99},
}

// Thêm sản phẩm thứ 3
products = append(products, Product{id: "p3", title: "4K Monitor", price: 299.99})
```
