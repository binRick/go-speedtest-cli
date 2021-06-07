FROM alpine:latest as build
ARG GO_VERSION=1.16.4

RUN apk update
RUN apk upgrade
RUN apk add git tcpdump ngrep bash zsh rsync sqlite dnsmasq nftables iperf curl wget tmux file gcompat \
            librrd strace openssh npm gcc automake make

WORKDIR /usr/src
COPY docker/files/go1.16.4.linux-amd64.tar.gz /usr/src/go1.16.4.linux-amd64.tar.gz
RUN ls -altr /usr/src && \
    file  /usr/src/go1.16.4.linux-amd64.tar.gz && \
    tar zxf go1.16.4.linux-amd64.tar.gz && mv go /opt && \
    ln -sf /opt/go/bin/go /usr/bin/go && \
    ln -sf /opt/go/bin/gofmt /usr/bin/gofmt  && \
    command -v go && \
    go version


RUN apk --update add \
      musl-dev \
      util-linux-dev python3 && \
    pip3 install speedtest-cli

    
WORKDIR /go-speedtest-cli
COPY ./Makefile ./*.go ./go.mod ./go.sum /go-speedtest-cli/.
RUN go mod tidy
RUN go get

RUN rm -rf /go-speedtest-cli/bin

RUN make binary
#RUN /go-speedtest-cli/bin/go-speedtest-cli 

RUN make static
#RUN /go-speedtest-cli/bin/go-speedtest-cli-static

RUN make test
