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

const tyLeLamPhat = 3.5

func main() {
	name, age, isStudent := "Thuong", 22, true
	fmt.Println(age, name, isStudent)

	// khoanDauTu, laiSuatHangNam, soNamDauTu := 5000.0, 5.5, 3.0
	var khoanDauTu, laiSuatHangNam, soNamDauTu float64
	outputText("Nhập số tiền đầu tư: ")
	fmt.Scan(&khoanDauTu)
	outputText("Nhập lãi suất hàng năm (%): ")
	fmt.Scan(&laiSuatHangNam)
	outputText("Nhập số năm đầu tư: ")
	fmt.Scan(&soNamDauTu)

	giaTriTuongLai, giaTriThucTe := calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu)
	// formattedFV := fmt.Sprintf("Danh nghĩa: %.2f\n", giaTriTuongLai)
	// formattedRT := fmt.Sprintf("Thực tế (trừ lạm phát): %.2f\n", giaTriThucTe)
	// fmt.Println("Giá trị tương lai của khoản đầu tư", khoanDauTu, "là", giaTriTuongLai)
	// fmt.Println("Giá trị thực tế của khoản đầu tư", khoanDauTu, "là", giaTriThucTe)
	fmt.Printf("Danh nghĩa: %.2f\nThực tế (trừ lạm phát): %.2f\n", giaTriTuongLai, giaTriThucTe)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(khoanDauTu, laiSuatHangNam, soNamDauTu float64) (float64, float64) {
	fv := khoanDauTu * math.Pow(1+laiSuatHangNam/100, soNamDauTu)
	rfv := fv / math.Pow(1+tyLeLamPhat/100, soNamDauTu)
	return fv, rfv
}
