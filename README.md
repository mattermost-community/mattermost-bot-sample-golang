# mattermost-bot-sample-golang

If you have an existing team then skip this step and replace team_name with your existing team.
```
./bin/platform -create_team -team_name="botsample" -email="admin@example.com" -password="password1" -username="samplebot"
```

```
./bin/platform -create_user -team_name="botsample" -email="bot@example.com" -password="password1" -username="samplebot"
```


Lets setup a 2nd user named bill which we will use to login to Mattermost server.
```
./bin/platform -create_user -team_name="botsample" -email="bill@example.com" -password="password1" -username="bill"
```

Optional:  You can give bill `system_admin permissions`
```
./bin/platform -assign_role -email="bill@example.com" -role="system_admin"
```

Feel free to login into the Mattermost server at http://localhost:8065 with the bill account and navigate to the `botsample` team to interact with the bot.