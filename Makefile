APP_NAME := main
ENV := APP_PORT=8000

run:
	$(ENV) go run *.go

build:
	$(ENV) go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)
