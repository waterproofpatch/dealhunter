#!/bin/bash
if [ -z "$LOCAL_DEV" ]
then
  curl -X POST -H "Authorization: Bearer $(echo $SHUTDOWN_ACCESS_TOKEN)" http://localhost:$PORT/shutdown
else
  curl -k -X POST -H "Authorization: Bearer $(echo $SHUTDOWN_ACCESS_TOKEN)" https://localhost:$PORT/shutdown
fi

