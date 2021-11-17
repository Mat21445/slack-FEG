package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
)

// struct that will hold data from Epic Games
type tmpData struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					Title         string `json:"title"`
					Description   string `json:"description"`
					EffectiveDate string `json:"effectiveDate"`
				} `json:"elements"`
			} `json:"searchStore"`
		} `json:"Catalog"`
	} `json:"data"`
}

func main() {
	// read the url from CLI flag "-url="
	url := flag.String("url", SlackURL, "Slack webhook URL")
	flag.Parse()
	if *url == "YOUR_SLACK_URL" {
		fmt.Println("Please add correct url of your Slack webhook! Try 'slack-FEG -url=<your-Slack-webhook-URL>'.")
		return
	}

	// time at the start of application (will be used to check whether the game was announced yet or not)
	TmpTime := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	p := &TmpTime //pointer at the TmpTime variable

	// time loop
	for _ = range time.Tick(time.Hour) {
		CheckAndSend(url, TmpTime)

		// Change the TmpTime variable
		*p = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	}
}

func CheckAndSend(url *string, time1 string) {
	// Download and unmarshal JSON information
	var web_data tmpData = fetch()

	// loop that checks if the game is currently free and not posted yet
	for i := 0; i < len(web_data.Data.Catalog.SearchStore.Elements); i++ {
		if web_data.Data.Catalog.SearchStore.Elements[i].EffectiveDate < time.Now().UTC().Format("2006-01-02T15:04:05.000Z") && web_data.Data.Catalog.SearchStore.Elements[i].EffectiveDate > time1 {

			// send the message using Slack webhook
			attachment1 := CreateFEGAttachment(web_data.Data.Catalog.SearchStore.Elements[i].Title, web_data.Data.Catalog.SearchStore.Elements[i].Description)
			payload := slack.Payload{
				Attachments: []slack.Attachment{attachment1},
			}
			err := slack.Send(*url, "", payload)
			if len(err) > 0 {
				fmt.Printf("error: %s\n", err)
			}
		}
	}
}

// function for getting the data from Epic Games Website
func fetch() tmpData {

	response, err := http.Get("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=PL&allowCountries=PL")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var p1 tmpData
	err = json.Unmarshal(body, &p1)
	if err != nil {
		log.Fatal(err)
	}

	return p1
}

// function that creates the attachment
func CreateFEGAttachment(title string, description string) (attachment slack.Attachment) {

	attachment.AddField(slack.Field{Title: "New Free Epic Game Available!:", Value: title}).AddField(slack.Field{Title: "Description:", Value: description})
	attachment.AddAction(slack.Action{Type: "button", Text: "Check it!", Url: "https://www.epicgames.com/store/en-US/free-games", Style: "primary"})

	return attachment
}
