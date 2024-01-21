# Run from the dir containing main.go
GOOS=linux GOARCH=386 go build .

SITE_ID=834669

echo "Uploading manager script..."
# don't override running file, causes IO errors - just did his once
#scp scripts/manager.sh dealhunter@ssh-dealhunter.alwaysdata.net:~/manager2.sh

# uploading binary to staging
echo "Uploading binary..."
scp deals dealhunter@ssh-dealhunter.alwaysdata.net:~/deals-staging

# restart the site
curl -X POST --basic --user "$API_KEY account=dealhunter:" https://api.alwaysdata.com/v1/site/$SITE_ID/restart
