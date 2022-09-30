set pathprefix=%1

swag init -g %pathprefix%/main.go -o %pathprefix%/api/v1 --pd
