# alertigo
Golang alerting api for Telegram messaging, tho goal is to have a chat where one of the guests will be a bot.
Hence, this bot will be able to send alert messages to the group and also be able to be listening the conversation 
in order to answer some technical questions. (ie: /status, and anything starting by a slash).

Compiled with runtime with: 
+ GOOS=windows GOARCH=386 go build -o alertigo.exe *.go
+ GOOS=linux GOARCH=386 go build -o alertigo.linux *.go
+ GOOS=darwin GOARCH=386 go build -o alertigo *.go

Build Docker image with
+ cp /source_cfg_files/*env* .
+ docker build -f docker/Dockerfile . -tag alertigo
+ docker run --publish 8000:8000 --name alertigo --rm alertigo --restart=always alertigo
Kubernetes
docker build --file=docker/Dockerfile -t docker-registry.oneboxtickets.com/oneboxtm/dynamic-pricing .
docker push docker-registry.oneboxtickets.com/oneboxtm/dynamic-pricing
then stop the pod within kubernetes to let the cluster instantiate a new one from the latest in docker registry.


## TODO list
+ set alert with threashold, time window, elk term filter, time frame, min num events, elk index, alert name, elk hosrt, elk port
+ oauth onebox
+ Get environment and general configurations from environment vars
+ unit testing
+ Handle requests by gorutines pool and control them by channels (marcio.io)
+ dockerize for kubernetes


## DONE list
+ load configuration from yml file
+ register to eureka
+ send message to chat
+ make bot to be listenig the chat and answer basic questions (/status, /who is down, ...)
+ load alert rules from yaml config file