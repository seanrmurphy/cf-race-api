#! /usr/bin/env bash

API_ENDPOINT="${APIENDPOINT:=https://cf-race-api.NAME.workers.dev}"

echo
echo "Querying filtered race results..."
curl -s -X POST \
	-H "Content-Type: application/json" \
	-d @filter.json \
	"${APIENDPOINT}/results" | jq
