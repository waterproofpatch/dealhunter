# Run from the dir containing main.go
GOOS=linux GOARCH=386 go build .
scp deals waterproofpatch@ssh-waterproofpatch.alwaysdata.net:~/
