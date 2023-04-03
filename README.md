# Slack-Bot in Go

This is a simple Slack-Bot written in Go that responds to a "hello" message with a greeting. 

## Requirements

* [Go](https://golang.org/dl/) (version 1.16 or higher)
* A Slack account and a [Slack Bot Token](https://api.slack.com/tutorials/tracks/getting-a-token) and a [Slack App Token](https://api.slack.com/apis/connections/socket#:~:text=Under%20Basic%20Information%2C%20scroll%20to%20the%20App-level%20tokens,with%20setting%20up%20your%20app%20for%20Socket%20Mode)

## Installation

1. Clone this repository:
      
         git clone https://github.com/pratu098/slack-bot-go.git

2. Install the dependencies:
         
         go mod download

3. Set your Slack Bot Token and Slack App Token as an environment variable:
         
         export SLACK_APP_TOKEN=your_slack_app_token
         export SLACK_BOT_TOKEN=your_slack_bot_token

4. Start the bot:
         
         go run main.go


## Usage

To use the bot, simply send a message containing the word "hello" to the channel where the bot is located. The bot will respond with a greeting.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
