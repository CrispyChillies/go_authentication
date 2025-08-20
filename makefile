APP_NAME=go_auth_app

run:
	go run ./...

build:
	go build -o $(APP_NAME) ./...

clean:
	rm -f $(APP_NAME)
