FILE=".env"
if [ -f "$FILE" ]; then
  set -o allexport;
  source $FILE;
  set +o allexport;
fi;
echo "Current version: "
migrate -source "file://migrations" -database $PG_ADDRESS version
migrate -source "file://migrations" -database $PG_ADDRESS up