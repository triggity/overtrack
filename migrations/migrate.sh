#! /usr/bin/env bash

curl -H 'Content-Type: application/x-ndjson' -XPOST -u elastic:changeme 'localhost:9200/ow/maps/_bulk?pretty' --data-binary @gamemaps.json
curl -H 'Content-Type: application/x-ndjson' -XPOST -u elastic:changeme 'localhost:9200/ow/users/_bulk?pretty' --data-binary @users.json
curl -H 'Content-Type: application/x-ndjson' -XPOST -u elastic:changeme 'localhost:9200/ow/games/_bulk?pretty' --data-binary @games.json