#/bin/bash

if [ -z "$1" ]
then
	echo "please enter dbstart or filestart or stop"
	exit
fi

case "$1" in
	"dbstart")  
		 `go run login.go database.go client.go  messages.go`;;
    "filestart") 
    	 `go run login.go client.go  messages.go config.go`;;
    "stop")  
		 `ps -ef | grep 'login.go' | awk '{print $2}' | xargs  kill -9`;;
     * ) 
		echo " please enter dbstart or filestart or stop"
esac
