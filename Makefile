all: fmt
	go test ./... -coverprofile=covprofile
	go tool cover -html=covprofile -o coverage.html


fmt:
	goimports -w .

