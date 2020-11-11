# Creating user
You should go to your unifi web controller with your browser, note the IP of the controller somewhere we refer it in the doc below as `controllerIP`
Create a user in your unifi controller web interface and we will refer to as `username` and `password` for now on.

# Debug cli
For all the below command, just run `make` once, this will generate a `upc` binary.

## available commands
```
sage of ./upc:
  -block string
    	Mac address or group of device to block
  -c string
    	Unifi controller host
  -config string
    	Comfiguration file holding groups
  -list string
    	List [client|network|all|device] (default "client")
  -p string
    	Unifi controller username
  -poeoff string
    	DeviceID of switch to disable poe on, to be use in conjunction of -port
  -poeon string
    	DeviceID of switch to enable poe on, to be use in conjunction of -port
  -port int
    	Port to allow poe on, must be used in conjunction with either -poeon or -poweroff
  -time
    	Prefix output with datetime default is false
  -u string
    	Unifi controller username
  -unblock string
    	Mac address or group of device to unblock```

## Listing clients
```
> upc -u username -p password -c controllerIP -list client
          NAME          | MANUFACTURER | WIFI/WIRED |   LAST SEEN    |        MAC        | BLOCKED
-------------------------+--------------+------------+----------------+-------------------+----------
  Nintendo Switch        | Nintendo     | wifi       | 13 seconds ago | b8:bb:ec:cc:ab:e2 | false
  iPhone-kid1            |              | wifi       | 13 seconds ago | 9e:bb:f4:cc:bd:f9 | false
  Time Capsule           | Apple        | ethernet   | 24 seconds ago | 00:bb:36:cc:ff:ec | false
  Eve Extend             | DexatekT     | wifi       | 13 seconds ago | 3c:bb:9d:cc:dd:a6 | false
  DESKTOP-LQCEK6V        | Microsof     | wifi       | 30 seconds ago | 28:bb:a8:cc:aa:10 | false
```

## Listing Networks
```
> upc -u username -p password -c controllerIP -list network
     NAME    | SECURITY | ENABLED | GUEST
-------------+----------+---------+--------
  guestwifi  | wpapsk   | false   | true
  foobar     | wpapsk   | false   | false
  mywifi     | wpapsk   | true    | false
```

## Listing UnifiDevices
```
> upc -u username -p password -c controllerIP -list device
     NAME     | TYPE | MODEL  |            ID            |        MAC        |  PORT  | PORTHASPOE | PORTISUP |         PORTCONF
--------------+------+--------+--------------------------+-------------------+--------+------------+----------+---------------------------
  Gateway     | ugw  | UGW3   | 123445a09292921234567651 | 18:bb:29:aa:e2:dd |        |            |          |
              |      |        |                          |                   | wan    | false      | true     |
              |      |        |                          |                   | lan    | false      | true     |
              |      |        |                          |                   | lan2   | false      | false    |
  Office      | uap  | U7LT   | 123445a09292921234567652 | 74:bb:c2:aa:db:dd |        |            |          |
  Basement    | uap  | U7NHD  | 123445a09292921234567653 | b4:bb:e4:aa:f8:dd |        |            |          |
  1st Floor   | uap  | U7LT   | 123445a09292921234567655 | 74:bb:c2:aa:98:dd |        |            |          |
  Main Switch | usw  | US8P60 | 123445a09292921234567654 | 74:bb:c2:aa:0a:dd |        |            |          |
              |      |        |                          |                   | Port 1 | false      | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 2 | false      | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 3 | false      | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 4 | false      | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 5 | true       | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 6 | true       | true     | 5a5db6819348239423948233
              |      |        |                          |                   | Port 7 | true       | true     | 5a5db6819348239423948234
              |      |        |                          |                   | Port 8 | true       | true     | 5a5db6819348239423948233
```

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
