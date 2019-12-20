
# Listing clients
go run *.go -u username -p password -c 10.0.80.20 -block 7c:2f:80:18:74:e5

# Blocking a client
go run *.go -u username -p password -c 10.0.80.20 -block 7c:2f:80:18:74:e5

# UnBlocking a client
go run *.go -u username -p password -c 10.0.80.20 -unblock 7c:2f:80:18:74:e5

# Running in server mode

go run *.go -u username -p password -c 10.0.80.20 -m server