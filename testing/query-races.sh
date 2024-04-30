#! /usr/bin/env bash

API_ENDPOINT="${APIENDPOINT:=https://cf-race-api.NAME.workers.dev}"

echo
echo "Querying first race..."
curl -X GET \
	"${APIENDPOINT}/race/1"

echo
echo
echo "Querying second race..."
curl -X GET \
	"${APIENDPOINT}/race/2"

echo
echo
echo "Querying third race..."
curl -X GET \
	"${APIENDPOINT}/race/3"
