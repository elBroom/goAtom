FROM golang:1.8.3-alpine
RUN apk add --update alpine-sdk
WORKDIR /go/src/github.com/elBroom/goAtom
ADD . .
RUN go build -a -o app_ .
EXPOSE 3030
CMD ["./app_"]

# docker build -t elbroom/goatom .
# docker push elbroom/goatom