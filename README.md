# slack-FEG
This is a simple Go application that gets the information about latest free games on Epic Games Store and then sends you a message about them via Slack.
It frequently (once per hour) gets information from Epic Games Store (response from https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=PL&allowCountries=PL), checks if anything new appeared and sends the message to Slack (using Webhook).<br><br>

Possible upgrades and additional features:
1. More information can be taken from JSON response (especially original price, but also images (url) - unluckily, large sized, probably hard to use as a part [or background] of a Slack attachement). The app could have shown the promotion end time to hurry the user.
2. Time loop inside application could be replaced by using cron (Linux), sheduled task (Win), etc. This was omitted as it would complicate the process of using the application. Basicly, the application could be set to check new free games only on Thursday after 5 P.M., since they are announced by this time.<br><br>

Layout of the message sent:<br><br>
![image](https://user-images.githubusercontent.com/92634025/140386422-ffb55a68-8e19-466b-be10-6bdfb801ee72.png)

#Build and Installation

**Build with a default Slack Webhook URL (no user input needed while running the App):**<br>
Build the application:
```
$ go build -ldflags "-X main.SlackURL="<your-Slack-Webhook-url>""
```
And then install (if needed):
```
$ go install
```
You can also instantly build and install with:
```
$ go install -ldflags "-X main.SlackURL="<your-Slack-Webhook-url>""
```
<br>

**Build without default URL (user input necessary on each startup):**<br>

Build the application:
```
$ go build
```
Install:
```
$ go install
```
Correct startup command:
```
$ slack-FEG -url=<your-Slack-Webhook-url>
```
