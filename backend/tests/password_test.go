package tests

import (
	"backend/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const password = "golang@Cn@123"

func Test_generate_pwd(t *testing.T) {
	passwd := "golang@Cn@123"
	cryptPwd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	fmt.Printf("%s:%s", "密码加密：", cryptPwd)
	fmt.Println()
	fmt.Print(utils.TimeStamp())
	fmt.Println()
}

func Test_compare_pwd(t *testing.T) {
	cryptPwd := "$2a$10$sq9OLRpPH./NPVNgwF/bMO6KKgUDU2kkc1dDwJ/HvHQhcN9ARPzZG"
	if err := bcrypt.CompareHashAndPassword([]byte(cryptPwd), []byte(password)); err != nil {
		fmt.Println("密码不匹配")
	} else {
		fmt.Println("密码匹配")
	}
}
