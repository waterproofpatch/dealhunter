# Run from the dir containing main.go
GOOS=linux GOARCH=386 go build .

# io error when overwriting in-use file, need smarter staging system
scp deals dealhunter@ssh-dealhunter.alwaysdata.net:~/deals
scp server.crt dealhunter@ssh-dealhunter.alwaysdata.net:~/server.crt
scp server.key dealhunter@ssh-dealhunter.alwaysdata.net:~/server.key

#curl -X POST --basic --user "APIKEY: $API_KEY" https://api.alwaysdata.com/v1/site/waterproofpatch/restart/
#curl -X POST --basic --user "$API_KEY:" https://api.alwaysdata.com/v1/site/waterproofpatch.alwaysdata.net/restart/
