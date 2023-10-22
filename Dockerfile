FROM        golang:1.20.10-alpine3.18
WORKDIR     /webhook
COPY        main.go /webhook/main.go 
COPY        start.sh /webhook/start.sh
RUN         apk add --update -t build-deps curl
RUN         cd /webhook \
            && go mod init webhook \
            && go get golang.org/x/time/rate \
            && go mod tidy \
            && chmod +x start.sh 
EXPOSE      9000
ENTRYPOINT  ["/webhook/start.sh"]
