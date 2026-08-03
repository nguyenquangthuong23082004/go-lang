package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Sử dụng constructor function để tạo đối tượng
	var appUser *User = newUser(userFirstName, userLastName, userBirthdate)

	// Gọi method xuất dữ liệu lần đầu
	appUser.outputUserDetails()

	// Gọi method xóa tên người dùng (Pointer Receiver)
	appUser.clearUserName()

	// Xuất dữ liệu lần hai để kiểm tra
	appUser.outputUserDetails()
}

func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func (u User) outputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Created At:", u.createdAt)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
