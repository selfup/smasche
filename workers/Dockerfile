FROM golang:alpine AS build

ENV CGO_ENABLED=0 APP=smasche

RUN apk add git --no-cache

COPY ./workers $GOPATH/src/github.com/selfup/$APP

WORKDIR $GOPATH/src/github.com/selfup/$APP

RUN go get github.com/selfup/gdsm

RUN go build -o /go/bin/$APP

FROM scratch

EXPOSE 9000

COPY --from=build /go/bin/smasche /go/bin/smasche

CMD ["/go/bin/smasche"]
