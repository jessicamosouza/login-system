package sendSms

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func send(phone, message string, w http.ResponseWriter) {
	accountSid := "AC4e63ab2914ed192dd291d04557482cc1"
	authToken := "b34423d0ad5d19a551141b0d2855ad6e"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+55" + phone)
	params.SetFrom("+16802083585")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		log.Println("Response: " + string(response))
	}

	w.Write([]byte("Mensagem enviada"))

}
