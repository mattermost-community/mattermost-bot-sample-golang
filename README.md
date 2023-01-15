# Mattermost Bot Sample

## Overview

This sample Bot shows how to use the Mattermost [Go driver](https://github.com/mattermost/mattermost-server/blob/master/model/client4.go) to interact with a Mattermost server, listen to events and respond to messages. Documentation for the Go driver can be found [here](https://pkg.go.dev/github.com/mattermost/mattermost-server/v6/model).

Highlights of APIs used in this sample:

- Log in to the Mattermost server
- Create a channel
- Modify user attributes
- Connect and listen to WebSocket events for real-time responses to messages
- Post a message to a channel

This Bot Sample was tested with Mattermost server version 7.5.2.

## Setup Server Environment

### Via Docker And Docker-Compose

1 - Ensure [Docker](https://www.docker.com/get-started) and [Docker-Compose](https://docs.docker.com/compose/install/) are installed on your system.

2 - Run `docker-compose up -d --build` and the mattermost client will be built and will expose the port `8065` to your system's localhost.

3 - Run `./add_users.sh`. The login information for the Mattermost client will be printed.

4 - Log into your mattermost instance using these credentials and create a bot account following [these](https://developers.mattermost.com/integrate/reference/bot-accounts/) instructions.

5 - Copy `example.env` to `.env` and fill in the bot token (obtained from the previous step), team name, etc. Alternatively, just provide your credentials as environment variables.

5 - Start the Bot.

```
make run
```

You can verify the Bot is running when

- `Logged in to mattermost` appears on the command line.
- `Hi! I am a bot.` is posted in your specified channel.

See "Test the Bot" for testing instructions

### Via Direct System Install/Setup

1 - [Install](http://docs.mattermost.com/install/requirements.html) or [upgrade](https://docs.mattermost.com/administration/upgrade.html) to Mattermost server version 3.10+, and verify that the Mattermost server is running on [http://localhost:8065](http://localhost:8065).

On the commands below, if you are running Mattermost server version 5.0 or later, use `./bin/mmctl`. If you are running version 4.10 or earlier, use `./bin/platform`.

Learn more about the `mmctl` CLI tool in the [Mattermost documentation](https://docs.mattermost.com/administration/mmctl-cli-tool.html).

2 - Create a team for the Bot to run. If you have an existing team, you may skip this step and replace `botsample` with your existing team in subsequent steps.

```
./bin/mmctl team create --name botsample --display_name "Sample Bot playground" --email "admin@example.com"
```

3 - Create a user, `bill`, which we will use to log in and interact with the Bot.

```
./bin/mmctl user create --email="bill@example.com" --password="Password1!" --username="bill"
```

4 - (Optional) Give `bill` `system_admin` permissions.

```
./bin/mmctl roles system_admin bill
```

5 - Log in to [http://localhost:8065](http://localhost:8065) as `bill` and verify the account was created successfully.

6 - Create a bot account following [these](https://developers.mattermost.com/integrate/reference/bot-accounts/) instructions.

7 - Copy `example.env` to `.env` and fill in the bot token (obtained from the previous step), team name, etc. Alternatively, just provide your credentials as environment variables.

## Setup Bot Development Environment

1 - Follow the [Developer Machine Setup](https://docs.mattermost.com/developer/dev-setup.html) instructions to setup the bot development environment.

2 - Clone the GitHub repository to run the sample.

```
git clone https://github.com/mattermost/mattermost-bot-sample-golang.git
cd mattermost-bot-sample-golang
```

3 - Log into your mattermost instance using these credentials and create a bot account following [these](https://developers.mattermost.com/integrate/reference/bot-accounts/) instructions.

4 - Copy `example.env` to `.env` and fill in the bot token (obtained from the previous step), team name, etc. Alternatively, just provide your credentials as environment variables.

5 - Run with

```
make run
```

You can verify the Bot is running when

- `Logged in to mattermost` appears on the command line.
- `Hi! I am a bot.` is posted in your specified channel.

## Test the Bot

1 - Log in to the Mattermost server as `bill@example.com` and `Password1!`.

2 - Join your specified channel.

3 - Post a message in the channel such as `hello?` to see if the Bot responds. You should see a response if the Bot is running.

## Stop the Bot

1 - In the terminal window, press `CTRL+C` to stop the bot.
