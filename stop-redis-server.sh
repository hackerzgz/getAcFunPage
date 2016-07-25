#!/bin/bash

# Stop redis-cli
redis-cli shutdown

# Stop redis-server Service.
service redis-server stop