#!/bin/bash

# Welcome Information.
username=`whoami`
echo "Hello $username!"
echo "Welcome to use Redis."


# Check the :6379 Port.
# p=6379
# i1=`/usr/bin/nmap -sS 127.0.0.1 -p $p | grep $p | awk '{printf $2}'`
# if [[ "$i1"=="open" ]]; then
# 	# Redis Server Running.
# 	echo ":6379 Port has been running. -- $i1"
# 	exit
# fi


# running Redis Server Process.
redis-server &

