#! /usr/bin/env bash

API_ENDPOINT="${APIENDPOINT:=https://cf-race-api.NAME.workers.dev}"

echo
echo "Creating first race..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race1.json \
	"${APIENDPOINT}/race"

echo
echo
echo "Creating second race..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race2.json \
	"${APIENDPOINT}/race"

echo
echo
echo "Creating third race..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race3.json \
	"${APIENDPOINT}/race"
