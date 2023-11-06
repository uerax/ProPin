GOOS=linux GOARCH=amd64 go build -o propin-cli -trimpath -ldflags "-s -w -buildid=" main.go
GOOS=linux GOARCH=arm64 go build -o propin-cli -trimpath -ldflags "-s -w -buildid=" main.go

./propin-cli -s [YOUR_SERVER_ADDR:PORT]/report -n [YOUR_SERVER_NAME] 2>&1 > ./log/propin-cli.log &