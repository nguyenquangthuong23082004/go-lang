package main

import (
	"fmt"
	"math"
)

// var khoanDauTu float64 = 5000
// var laiSuatHangNam float64 = 5.5
// var soNamDauTu float64 = 3
// var age int

// var giaTriTuongLai = khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)

func main() {
	const tyLeLamPhat float64 = 3.5
	name, age, isStudent := "Thuong", 22, true
	fmt.Println(age, name, isStudent)

	// khoanDauTu, laiSuatHangNam, soNamDauTu := 5000.0, 5.5, 3.0
	var khoanDauTu, laiSuatHangNam, soNamDauTu float64
	fmt.Print("Nhập số tiền đầu tư: ")
	fmt.Scan(&khoanDauTu)
	fmt.Print("Nhập lãi suất hàng năm (%): ")
	fmt.Scan(&laiSuatHangNam)
	fmt.Print("Nhập số năm đầu tư: ")
	fmt.Scan(&soNamDauTu)

	giaTriTuongLai := khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	giaTriThucTe := giaTriTuongLai / math.Pow(1+tyLeLamPhat/100, soNamDauTu)
	fmt.Println("Giá trị tương lai của khoản đầu tư", khoanDauTu, "là", giaTriTuongLai)
	fmt.Println("Giá trị thực tế của khoản đầu tư", khoanDauTu, "là", giaTriThucTe)
}
