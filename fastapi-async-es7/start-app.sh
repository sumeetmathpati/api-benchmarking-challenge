#!/bin/sh
# start-app.sh

while ! eval "[ $(curl --write-out %{http_code} --silent --output /dev/null http://elastic:9200/_cat/health?h=st) = 200 ]"; do
  echo "ES not ready yet..."
  sleep 1
done

>&2 echo "ES is up!"
>&2 curl --location --request PUT 'http://elastic:9200/events' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "settings": {
      "number_of_shards": 1,
      "number_of_replicas": 1
    }
  }'
exec "$@"