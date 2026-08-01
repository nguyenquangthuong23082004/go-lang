package main

import (
	"fmt"
	"math"
)

var khoanDauTu = 5000
var laiSuatHangNam = 5.5
var soNamDauTu = 3
var age int

var giaTriTuongLai = float64(khoanDauTu) * math.Pow(1+laiSuatHangNam/100, float64(soNamDauTu))

func main() {
	fmt.Println(age)
	fmt.Println("Giá trị tương lai của khoản đầu tư", khoanDauTu, "là", giaTriTuongLai)
}
