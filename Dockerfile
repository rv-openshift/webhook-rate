FROM        golang:1.21.3-alpine3.18 as build
WORKDIR     /app
COPY        main.go /app
COPY        internal /app/internal
RUN         cd /app \
            && go mod init webhook \
            && go mod tidy \
            && go build -o webhook .

FROM        alpine:3.18
ARG         pwd
ENV         emailpwd $pwd
COPY        --from=build /app/webhook /usr/local/bin/webhook
EXPOSE      8080
ENTRYPOINT  ["/usr/local/bin/webhook"]
