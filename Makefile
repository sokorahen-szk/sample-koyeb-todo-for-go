APP_NAME := main
ENV := PORT=8080

run:
	$(ENV) go run *.go

build:
	$(ENV) go build -o $(APP_NAME) *.go

clean:
	rm -f $(APP_NAME)
