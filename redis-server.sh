#!/bin/bash

# Welcome Information.
username=`whoami`
echo "Hello $username!"
echo "Welcome to use Redis."

# Check the Running Process.



# Find redis-conf
redis_path=`whereis redis`
redis_path=${redis_path#redis: }
redis_path+="/redis.conf"
printf 'Your redis-conf -> %s\n' $redis_path

# running Redis Server Process.
redis-server $redis_path &

