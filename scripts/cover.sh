#!/usr/bin/env bash

export MONGO_URL=$(docker inspect -f '{{.NetworkSettings.Networks.workshopgo_default.IPAddress}}' workshopgo_mongodb_1)
echo "mode: set" > full_coverage.out
for pkg in $(go list ./... | grep -v /vendor/); do
    go test -v -cpu 1 -coverprofile=coverage.out -covermode=set $pkg
    if [ $? -ne 0 ]; then
        exit 1
    fi
    grep -h -v "^mode: set" coverage.out >> full_coverage.out
done

go tool cover -func full_coverage.out
rm -f coverage.out full_coverage.out

exit 0
