#!/bin/sh
if [ -n "PG_ADDRESS" ]; then
  echo "Current version: "
  migrate -source "file://migrations" -database $PG_ADDRESS version
  migrate -source "file://migrations" -database $PG_ADDRESS up
fi;
app