# Mattermost Bot Sample

## Overview

This sample Bot shows how to use the Mattermost [Go driver](https://github.com/mattermost/mattermost-server/blob/master/model/client4.go) to interact with a Mattermost server, listen to events and respond to messages. Documentation for the Go driver can be found [here](https://godoc.org/github.com/mattermost/mattermost-server/model#Client).

Highlights of APIs used in this sample:
 - Log in to the Mattermost server
 - Create a channel
 - Modify user attributes 
 - Connect and listen to WebSocket events for real-time responses to messages
 - Post a message to a channel

This Bot Sample was tested with Mattermost server version 3.10.0.

## Setup Server Environment

### Via Docker And Docker-Compose
1 - Ensure [Docker](https://www.docker.com/get-started) and [Docker-Compose](https://docs.docker.com/compose/install/) are installed for your system

2 - Run `./add_users.sh`. The login information for the Mattermost client will be printed

3 - Run `docker-compose up -d --build` and the mattermost client will be built and will expose the port `8065` to your system's localhost

4 - Start the Bot.
```
make run
```
You can verify the Bot is running when 
  - `Server detected and is running version X.Y.Z` appears on the command line.
  - `Mattermost Bot Sample has started running` is posted in the `Debugging For Sample Bot` channel.

See "Test the Bot" for testing instructions
### Via Direct System Install/Setup
1 - [Install](http://docs.mattermost.com/install/requirements.html) or [upgrade](https://docs.mattermost.com/administration/upgrade.html) to Mattermost server version 3.10+, and verify that the Mattermost server is running on [http://localhost:8065](http://localhost:8065). 

On the commands below, if you are running Mattermost server version 5.0 or later, use `./bin/mmctl`. If you are running version 4.10 or earlier, use `./bin/platform`.

Learn more about the `mmctl` CLI tool in the [Mattermost documentation](https://docs.mattermost.com/administration/mmctl-cli-tool.html).

2 - Create a team for the Bot to run. If you have an existing team, you may skip this step and replace `team_name` with your existing team in subsequent steps.
```
./bin/mmctl team create --name botsample --display_name "Sample Bot playground" --email "admin@example.com"
```
3 - Create the user account the Bot will run as.
```
./bin/mmctl user create --email="bot@example.com" --password="Password1!" --username="samplebot"
```
4 - Create a second user, `bill`, which we will use to log in and interact with the Bot.
```
./bin/mmctl user create --email="bill@example.com" --password="Password1!" --username="bill"
```
5 - (Optional) Give `bill` `system_admin` permissions.
```
./bin/mmctl roles system_admin bill
```
6 - Add users to the team
```
./bin/mmctl team add botsample samplebot bill
```
7 - Verify the e-mail address
```
./bin/mmctl user verify samplebot
```
8 - Log in to [http://localhost:8065](http://localhost:8065) as `bill` and verify the account was created successfully. Then, navigate to the `botsample` team you created in step 2 to interact with the Bot.

## Setup Bot Development Environment

1 - Follow the [Developer Machine Setup](https://docs.mattermost.com/developer/dev-setup.html) instructions to setup the bot development environment.

2 - Clone the GitHub repository to run the sample.
```
git clone https://github.com/mattermost/mattermost-bot-sample-golang.git
cd mattermost-bot-sample-golang
```
3 - Start the Bot.
```
make run
```
You can verify the Bot is running when 
  - `Server detected and is running version X.Y.Z` appears on the command line.
  - `Mattermost Bot Sample has started running` is posted in the `Debugging For Sample Bot` channel.

## Test the Bot

1 - Log in to the Mattermost server as `bill@example.com` and `Password1!`.

2 - Join the `Debugging For Sample Bot` channel.

3 - Post a message in the channel such as `are you running?` to see if the Bot responds. You should see a response similar to `Yes I'm running` if the Bot is running.

## Stop the Bot

1 - In the terminal window, press `CTRL+C` to stop the bot. You should see `Mattermost Bot Sample has stopped running` posted in the `Debugging For Sample Bot` channel.
