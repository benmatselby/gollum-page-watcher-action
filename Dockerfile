FROM golang:1.18-alpine
LABEL maintainer="Ben Selby <benmatselby@gmail.com>"

RUN apk add --no-cache \
	bash \
	ca-certificates \
	curl \
	git

WORKDIR /src
COPY . /src

RUN go build -o /bin/action

ENTRYPOINT ["/bin/action"]
