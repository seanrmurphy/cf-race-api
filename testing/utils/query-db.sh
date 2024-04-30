#! /usr/bin/env bash

echo
echo "Querying basic db info..."
wrangler d1 execute cf-race-api --local --command "SELECT name, sql FROM sqlite_master;"

echo
echo "Querying race info table..."
wrangler d1 execute cf-race-api --local --command "SELECT * FROM race_info;"

echo
echo "Querying race results table..."
wrangler d1 execute cf-race-api --local --command "SELECT * FROM race_results;"
