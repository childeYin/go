#/bin/bash

if [ -z "$1" ]
then
	echo "please enter start or stop"
	exit
fi

case "$1" in
	"start")  
		 `server.go messages.go    `;;
    "stop")  
		 `ps -ef | grep 'server.go' | awk '{print $2}' | xargs  kill -9`;;
     * ) 
		echo " please enter start or stop"
esac
