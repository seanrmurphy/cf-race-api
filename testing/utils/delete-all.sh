#! /usr/bin/env bash

echo
echo "Deleting records from race results table..."
wrangler d1 execute cf-race-api --local --command "DELETE FROM race_results;"

echo
echo "Deleting records from race info table..."
wrangler d1 execute cf-race-api --local --command "DELETE FROM race_info;"
