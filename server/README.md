GOOS=linux GOARCH=amd64 go build -o propin-serv -trimpath -ldflags "-s -w -buildid=" main.go
GOOS=linux GOARCH=arm64 go build -o propin-serv -trimpath -ldflags "-s -w -buildid=" main.go

./propin-serv -token [YOUR_BOT_TOKEN] -id [YOUR_TG_ID] 2>&1 > ./log/propin-serv.log &