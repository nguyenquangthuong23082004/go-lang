package main

import (
	"fmt"
	"math"
)

var khoanDauTu float64 = 5000
var laiSuatHangNam float64 = 5.5
var soNamDauTu float64 = 3
var age int

var giaTriTuongLai = khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)

func main() {
	fmt.Println(age)
	fmt.Println("Giá trị tương lai của khoản đầu tư", khoanDauTu, "là", giaTriTuongLai)
}
