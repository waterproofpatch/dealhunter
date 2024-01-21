#!/bin/bash

STAGING_FILE="deals-staging"
PROD_FILE="deals"

# Function to handle SIGHUP
handle_sighup() {
    echo "SIGHUP signal received. Performing an action..."
    curl -X POST -H "Authorization: Bearer $(echo $SHUTDOWN_ACCESS_TOKEN)" http://dealhunter.alwaysdata.net/shutdown
    echo "Issued shutdown, sleeping..."
    sleep 3
    echo "Done sleeping, moving $STAGING to $PROD_FILE"
    mv $STAGING_FILE $PROD_FILE
    echo "Moved file, starting it..."
    ./$PROD_FILE &
    echo "Server started..."
}

# Trap SIGHUP and call handle_sighup function
trap handle_sighup SIGHUP

# Keep the script running
while true
do
    sleep 1
done
