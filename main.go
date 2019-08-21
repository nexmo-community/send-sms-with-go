package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nexmo-community/nexmo-go"
)

var apiKey = os.Getenv("API_KEY")
var apiSecret = os.Getenv("API_SECRET")
var fromNumber = os.Getenv("FROM_NUMBER")
var toNumber = os.Getenv("TO_NUMBER")

func main() {

	// Auth
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret(apiKey, apiSecret)

	// Init Nexmo
	client := nexmo.NewClient(http.DefaultClient, auth)

	// SMS
	smsContent := nexmo.SendSMSRequest{
		From: fromNumber,
		To:   toNumber,
		Text: "This is a message sent from Go!",
	}

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
}
