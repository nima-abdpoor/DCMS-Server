package helpers

import (
	"log"
	"strings"
)

func CheckUserPass(username, password string) bool {
	userPass := make(map[string]string)
	userPass["hello"] = "itsme"
	userPass["john"] = "doe"
	userPass["nima"] = "abdpoor"

	log.Println("checkUserPass", username, password, userPass)

	if val, ok := userPass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
