# Creating user
You should go to your unifi web controller with your browser, note the IP of the controller somewhere we refer it in the doc below as `controllerIP`
Create a user in your unifi controller web interface and we will refer to as `username` and `password` for now on.

# Debug cli
For all the below command, just run `make` once, this will generate a `upc` binary.

## available commands
```
Usage of ./upc:
  -block string
    	mac address of device to block
  -c string
    	Unifi controller host
  -config string
    	config file to launch in cron mode
  -list string
    	list [client|network|all|device]
  -p string
    	Unifi controller username
  -poeoff string
    	DeviceID of switch to disable poe on, to be use in conjunction of -port
  -poeon string
    	DeviceID of switch to enable poe on, to be use in conjunction of -port
  -port int
    	Port to allow poe on, must be used in conjunction with either -poeon or -poweroff
  -u string
    	Unifi controller username
  -unblock string
    	mac address of device to unblock
```

## Listing clients
`upc -u username -p password -c controllerIP -list client`

## Listing Networks
`upc -u username -p password -c controllerIP -list network`

## Listing UnifiDevices
`upc -u username -p password -c controllerIP -list device`

## Blocking a client from accessing the network
`upc -u username -p password -c controllerIP -block 7c:2f:80:18:74:e5`

## UnBlocking a client from accessing the network
`upc -u username -p password -c controllerIP -unblock 7c:2f:80:18:74:e5`

## Power off POE on a given switch port
First identify the switch ID of the poe switch you want to power on a port by running 
`upc -u username -p password -c 10.0.80.20 -list device`

Let say you see your switch as the ID `5d61b90be30dfa0ddd69c990`
`upc -u username -p password -c controllerIP -poeoff 5d61b90be30dfa0ddd69c990 -port 7`

## Power on POE off a given switch port
First identify the switch ID of the poe switch you want to power off a port by running 
`upc -u username -p password -c 10.0.80.20 -list device`

Let say you see your switch as the ID `5d61b90be30dfa0ddd69c990`
`upc -u username -p password -c controllerIP -poeon 5d61b90be30dfa0ddd69c990 -port 7`

# Docker image

A [docker image](https://hub.docker.com/repository/docker/jeanlaurent/upc) is available you can run the above command like 
`docker run -ti jeanlaurent/upc -u username -p password -c controllerIP -list device`
