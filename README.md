# DealsApp

## Generate Keys

```bash
cd backend
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

cd ../
openssl genrsa -out client.key 2048
openssl req -new -x509 -sha256 -key client.key -out client.crt -days 3650
```

## Development

### Backend

```bash
cd backend
go mod tidy
go run main.go
```

### Frontend

```bash
ng serve --configuration development --port 4300
```

## Deployment

## Backend

Deployed to alwaysdata.

```bash
bash backend/scripts/deploy.sh
```

## Frontend

Deployed to Amazon Amplify.

```bash
git checkout master
git push
```
