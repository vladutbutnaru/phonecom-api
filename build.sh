#!/usr/bin/env bash

export GOPATH=`pwd`

go get -u github.com/urfave/cli
go get -u github.com/go-resty/resty
go get -u github.com/yukithm/json2csv/cmd/json2csv
go get -u github.com/stretchr/testify

rm -r src/phonecom-go-sdk/
cp -r SDKs/go-client-generated src/phonecom-go-sdk
sed -i '/\/\/ to determine the Content-Type header/c\\tclearEmptyParams(localVarQueryParams)\n\n\t\/\/ to determine the Content-Type header' src/phonecom-go-sdk/*_api.go
sed -i '/case "ssv":/c\\tcase "ssv", "multi":' src/phonecom-go-sdk/api_client.go

echo 'package swagger


func clearEmptyParams(paramMap map[string][]string) {

	for key, value := range paramMap {
		if (len(value) == 1 && value[0] == "") {
			delete(paramMap, key)
		}
	}
}' > src/phonecom-go-sdk/util.go

GOARCH=amd64 GOOS=darwin go build -o phonecom-mac phonecom
GOARCH=amd64 GOOS=windows go build -o phonecom-windows phonecom
GOARCH=amd64 GOOS=linux go build -o phonecom-linux phonecom
