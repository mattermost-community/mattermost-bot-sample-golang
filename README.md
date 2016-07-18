# Mattermost Bot Sample

## Overview

This sample Bot shows how to use the Mattermost [Go driver](https://github.com/mattermost/platform/blob/master/model/client.go) interact with a Mattermost server, listen to events and respond to messages.  Documentation for the Go driver can be found [here](https://godoc.org/github.com/mattermost/platform/model#Client)

Highlights of APIs used in this sample:
 - Login to the Mattermost server
 - Create a channel
 - Modify a users attributes 
 - Connect and listen to WebSocket events for real time responses to messages
 - Post a message to a channel

## Setup Server Environment

1.  Install Mattermost server version 3.3+.  This sample is known it work with Mattermost server version 3.3.  Make sure the Mattermost server is running on [http://localhost:8065](http://localhost:8065).  If you need help installing Mattermost please refer to our installations guides [here](http://docs.mattermost.com/install/requirements.html)
2.  Create the team needed for the bot to run
    - If you have an existing team then skip this step and replace team_name with your existing team.
```
./bin/platform -create_team -team_name="botsample" -email="admin@example.com" -password="password1" -username="samplebot"
```
3.  Create the user account the Bot will run as
```
./bin/platform -create_user -team_name="botsample" -email="bot@example.com" -password="password1" -username="samplebot"
```
4.  Setup a 2nd user which we will use to login and interact with the Bot.
```
./bin/platform -create_user -team_name="botsample" -email="bill@example.com" -password="password1" -username="bill"
```
5.  Optional:  You can give the 2nd user `system_admin` permissions
```
./bin/platform -assign_role -email="bill@example.com" -role="system_admin"
```
6.  Login as bill to make sure everything is OK with the sample account.  Login into at [http://localhost:8065](http://localhost:8065) with the `bill` account and navigate to the `botsample` team to interact with the bot.

## Setup Bot Development Environment

1.  Follow the [Developer Machine Setup](http://docs.mattermost.com/developer/developer-setup.html) instructions to setup the bot development environment.
2.  Clone the Github repository to run the sample
```
git clone https://github.com/mattermost/mattermost-bot-sample-golang.git
cd mattermost-bot-sample-golang
```
3.  Start the bot
```
make run
```
  - You can verify the bot is running when you see `Server detected and is running version 3.2.0` on the command line
  - You will also see `Mattermost Bot Sample has started running` posted in the `Debugging For Sample Bot` channel

## Test the Bot

1.  Login to the Mattermost server as `bill@example.com` and `password1`
2.  Join the `Debugging For Sample Bot` channel
3.  Post a message in the channel like `are you running?` to see if the Bot responds.  You should see a response like `Yes I'm running` if the Bot is running.


## Stop the Bot

1.  In the terminal window press `CTRL+C` to stop the bot.
  -  You will see `Mattermost Bot Sample has stopped running` posted in the `Debugging For Sample Bot` channel
