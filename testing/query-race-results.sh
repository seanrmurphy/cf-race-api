#! /usr/bin/env bash

API_ENDPOINT="${APIENDPOINT:=https://cf-race-api.NAME.workers.dev}"

echo
echo "Querying first race results..."
curl -s -X GET \
	"${APIENDPOINT}/race/1/results" | jq

echo
echo
echo "Querying second race results..."
curl -s -X GET \
	"${APIENDPOINT}/race/2/results" | jq

echo
echo
echo "Querying third race results..."
curl -s -X GET \
	"${APIENDPOINT}/race/3/results" | jq -r
