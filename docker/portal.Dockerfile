##
## Build 
##
FROM golang:1.16-alpine AS ohurlshortener_builder 
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.com.cn,direct
ADD . /app
WORKDIR /app
RUN go mod download && go build -o ohurlshortener .

##
## Deploy
##
FROM alpine:latest
WORKDIR /app
COPY --from=ohurlshortener_builder /app/ohurlshortener .
EXPOSE 9091
ENTRYPOINT ["/app/ohurlshortener","-s","portal","-c","config.ini"] 

