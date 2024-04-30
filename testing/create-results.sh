#! /usr/bin/env bash

API_ENDPOINT="${APIENDPOINT:=https://cf-race-api.NAME.workers.dev}"

echo
echo "Creating first race results..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race-results1.json \
	"${APIENDPOINT}/race/1/results"

echo
echo
echo "Creating second race results..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race-results2.json \
	"${APIENDPOINT}/race/2/results"

echo
echo
echo "Creating third race results..."
curl -X POST \
	-H "Content-Type: application/json" \
	-d @race-results3.json \
	"${APIENDPOINT}/race/3/results"
