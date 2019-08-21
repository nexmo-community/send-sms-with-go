# Send an SMS with Go

How to use the [nexmo-go](https://github.com/nexmo-community/nexmo-go) client library to send an SMS using Go.

## Prerequisites

You'll need to have Golang installed on your development machine. Installation instructions can be found on the [official Golang website](https://golang.org/).

Alternatively, if you're new to Go, or you don't want to go through the installation process, you can work directly in the [Golang Playground](https://play.golang.org/) instead.

## Using the Nexmo API Client for Go

Your SMS and Go journey begins by installing the API client itself. You can do this by running:

```bash
go get https://github.com/nexmo-community/nexmo-go
```

Next, fire up your editor and create a new file called `main.go`. Then scaffold the basics of a Go application by typing (or copying) the following code:

```go
package main

import (
  "net/http"
  "github.com/nexmo-community/nexmo-go"
)

func main() {

}
```

> Note: If you save and the files in the import statement disappear, don't worry, they'll come back once you use them inside the `main()` function.

Now it's time to put some meat on those bones and instatiate the Nexmo client so you can actually make it do things.

Inside the `main()` function add the following:

```go
// Auth
auth := nexmo.NewAuthSet()
auth.SetAPISecret("API_KEY", "API_SECRET")

// Init Nexmo
client := nexmo.NewClient(http.DefaultClient, auth)
```

There are two things happening here.

First, you create an `auth` object that combines your API key and secret together using a helper function that will ensure everything is formatted correctly.

> Note: Your API key and secret can be found by logging into your [Nexmo Dashboard](https://dashboard.nexmo.com/sign-in). If you don't have an account yet, you can [sign up here](https://dashboard.nexmo.com/sign-up) and get a starter free credit to run this code!

Second, you instantiate a new `client` that will hold all the functionality the nexmo-go library provides. Your `auth` object is passed into this.

With this in place you can now perform actions on the Nexmo API, such as sending an SMS.

## Send SMS Messages with Go

With the Nexmo API client ready to go, your code will now look like this:

```go
package main

import (
  "net/http"
  "github.com/nexmo-community/nexmo-go"
)

func main() {
  // Auth
  auth := nexmo.NewAuthSet()
  auth.SetAPISecret(apiKey, apiSecret)

  // Init Nexmo
  client := nexmo.NewClient(http.DefaultClient, auth)
}
```

In order to send an SMS with Go you first need to create a `SendSMSRequest` object that contains all of the information an SMS needs to make it to its destination such as the number it should be send to, the number it's sendiing from and the text that should be displayed.

It can be added to `main.go` below where you instatiated the `client` and is formatted like this:

```go
smsContent := nexmo.SendSMSRequest{
  From: "447700900004",
  To:   "14155550105",
  Text: "This is a message sent from Go!",
}
```

Of couse, you'll want to replace these numbers with real ones for testing. The `To` number can be your own number but the `From` number must be an SMS capable number purchased via your [Nexmo Dashboard](https://dashboard.nexmo.com).

Now the only thing left to do is to actually _send_ the SMS. This is done using the `SendSMS` function provided by the API client.

Add a new line with the following code:

```go
smsResponse, _, err := client.SMS.SendSMS(smsContent)
```

Finally, add a quick bit of error checking and response output:

```go
// If there are errors, log them out
if err != nil {
  log.Fatal(err)
}
// All is well, print the message status returned
fmt.Println("Status:", smsResponse.Messages[0].Status)
```

The stage is set to send that SMS! Head to your terminal and from inside the folder you're working in run:

```bash
go run main.go
```

If everything worked you'll see `Status: 0` returned to the screen just before the familiar sound of your SMS notification rings out signalling your success.

## All The Code

For your reference, all the code for sending an SMS with Go comes together like so:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/nexmo-community/nexmo-go"
)

func main() {

	// Auth
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("API_KEY", "API_SECRET")

	// Init Nexmo
	client := nexmo.NewClient(http.DefaultClient, auth)

	// SMS
	smsContent := nexmo.SendSMSRequest{
    From: "447700900004",
    To:   "14155550105",
    Text: "This is a message sent from Go!",
  }

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
}
```

## Where To Go From Here?

The next change you can make to this is to make the file a little more secure by removing the hardcoded API key, API secret, and the phone numbers.

A good way to do this is to move them to environment variables that are stored in a `.env` file.

Try implementing this using the [godotenv](https://github.com/joho/godotenv) package and quickly shore up your security.

## Further Reading

If sending an SMS with Go has go you exicted about what other communication elements you could be adding to your application then take a look at the examples on the [nexmo-go GitHub repository](https://github.com/nexmo-community/nexmo-go).

There you'll find code for using many other aspects of the Nexmo API such as making phone calls, receiving SMS messages and verifying phone numbers.
