build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/mailer mailer/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/validator validator/main.go