# Run it
go run *.go -u username -p password -c 10.0.80.20 -config config.yaml

with config.yaml

```
mac: 00:22:d7:52:8b:74
blockTime: 18:00
unblockTime: 19:00
```

should result with device with mac address of `00:22:d7:52:8b:74` being blocked between 18:00 and 19:00 and unblocked otherwise

# Debug cli

## Listing clients
go run *.go -u username -p password -c 10.0.80.20 -list client

## Listing Networks
go run *.go -u username -p password -c 10.0.80.20 -list network

## Blocking a client
go run *.go -u username -p password -c 10.0.80.20 -block 7c:2f:80:18:74:e5

## UnBlocking a client
go run *.go -u username -p password -c 10.0.80.20 -unblock 7c:2f:80:18:74:e5

