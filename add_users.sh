USER_EMAIL="bill@example.com"
USER_PASS="Password1!"
USER_NAME="bill"
docker exec test_mm mattermost team create --name botsample --display_name "Sample Bot playground" --email "admin@example.com"
docker exec test_mm mattermost user create --email="bot@example.com" --password="Password1!" --username="samplebot"
docker exec test_mm mattermost user create --email="${USER_EMAIL}" --password="${USER_PASS}" --username="${USER_NAME}"
docker exec test_mm mattermost roles system_admin bill
docker exec test_mm mattermost team add botsample samplebot bill
docker exec test_mm mattermost user verify samplebot

echo "Default user credentials"
echo "Email: ${USER_EMAIL} - Password: ${USER_PASS}"
