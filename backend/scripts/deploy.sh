# Run from the dir containing main.go
GOOS=linux GOARCH=386 go build .
scp deals dealhunter@ssh-dealhunter.alwaysdata.net:~/deals
#curl -X POST --basic --user "APIKEY: $API_KEY" https://api.alwaysdata.com/v1/site/waterproofpatch/restart/
#curl -X POST --basic --user "$API_KEY:" https://api.alwaysdata.com/v1/site/waterproofpatch.alwaysdata.net/restart/