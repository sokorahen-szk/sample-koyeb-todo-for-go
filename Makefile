APP_NAME := main
ENV := PORT=8080

run:
	$(ENV) go run main.go

build:
	$(ENV) go build -o $(APP_NAME) main.go

clean:
	rm -f $(APP_NAME)
