package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Sử dụng constructor function từ package user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Tạo đối tượng admin mới bằng NewAdmin
	admin := user.NewAdmin("admin@example.com", "supersecret")

	// Gọi method xuất dữ liệu lần đầu
	appUser.OutputUserDetails()

	// Gọi method xóa tên người dùng (Pointer Receiver)
	appUser.ClearUserName()

	// Xuất dữ liệu lần hai để kiểm tra
	appUser.OutputUserDetails()

	fmt.Println("\n--- ADMIN DETAILS ---")
	// Gọi trực tiếp các method của User trên đối tượng Admin (Anonymous Embedding)
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
