package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func checkMail(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"user_info.user_mail": newValues}, person)
	if person.UserInfos.UserMail != "" {
		return false
	}
	return true
}

func checkPhone(newValues string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"contact_info.user_phone": newValues}, person)
	if person.Contacts.UserPhone != "" {
		return false
	}
	return true
}
func checkBeaconType(beaconType int) string {
	if beaconType == 0 {
		return "Tasma"
	}
	if beaconType == 1 {
		return "Bileklik"
	}
	if beaconType == 5 {
		return "Kalemlik"
	}
	return ""
}
func checkObjID(id string) (string, bool) {
	var s = bson.IsObjectIdHex(id)
	if s == true {
		return id, true
	}
	return "", false
}
func writeResponse(w http.ResponseWriter, jsonValue string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonValue))
}
func addError(byteJSON []byte) []byte {
	var m map[string]interface{}
	json.Unmarshal(byteJSON, &m)
	m["error"] = false
	newData, _ := json.Marshal(m)
	return newData
}
func checkPermission(token string) bool {
	person := &Person{}
	connection.Collection("users").FindOne(bson.M{"user_info.user_token": token}, person)
	if person.UserInfos.RoleLvl == 5 {
		return true
	}
	return false
}
func checkPhoneNumber(number string) bool {
	regex := regexp.MustCompile("^[+]?(?:[0-9]{2})?[0-9]{10}$")
	match := regex.MatchString(number)
	if match == true {
		return true
	}
	return false
}
func checkEmailValidity(email string) bool {
	regex := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-zA-Z0-9-.]+$")

	match := regex.MatchString(email)
	if match == true {
		return true
	}
	return false
}
func sendRegisterMail(token string, email string) bool {
	url := "http://localhost:8090/registercontrol?token="
	from := mail.NewEmail("Example User", "abdurrahman262@hotmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>" + url + token + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(response.StatusCode)
		return false
	}
	return true

}
