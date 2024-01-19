# Run from the dir containing main.go
GOOS=linux GOARCH=386 go build .

# io error when overwriting in-use file, need smarter staging system
scp scripts/shutdown.sh dealhunter@ssh-dealhunter.alwaysdata.net:~/shutdown.sh
echo "Running shutdown..."
ssh dealhunter@ssh-dealhunter.alwaysdata.net 'bash ~/shutdown.sh'
echo "Ran shutdown..."
scp deals dealhunter@ssh-dealhunter.alwaysdata.net:~/deals-2

curl -X POST --basic --user "$APIKEY account=dealhunter:" https://api.alwaysdata.com/v1/site/dealhunter/restart/
