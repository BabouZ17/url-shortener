# url-shortener
A little url shortener project

## Setup using docker-compose
```
docker-compose up --build
```

## Setup using your own instance of postgresql
```
export POSTGRES_PASSWORD=your_password
export CONFIG_PATH=${PWD}/url-shortener/pkg/config/local.json

go run cmd/main.go
```

## Add a url
```
curl -X POST \
  'localhost:8080/urls' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "target": "https://www.google.com"
}'
```

## List all saved urls
```
curl -X GET \
  'localhost:8080/urls' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json'
}'
```

## Consume shortened url (aka "alias")
To use the shortened url, make a request using the "alias" returned when saving the full url.
```
curl -X GET \
  'localhost:8080/redirect/XVlBzg' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json'
```