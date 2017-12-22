build: main.go
	@go build -o bin/blockchain main.go

blockchainClient: blockchainClient.go
	@go build -o bin/blockchainClient blockchainClient.go
	@chmod +x bin/blockchainClient

powTestClient: powClient.go
	@go build -o bin/powClient powClient.go
	@chmod +x bin/powClient
