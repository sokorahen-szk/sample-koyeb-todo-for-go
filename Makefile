APP_NAME := main

run:
	go build -o $(APP_NAME) main.go
	APP_PORT=8000 \
		./$(APP_NAME)

clean:
	rm -f $(APP_NAME)
