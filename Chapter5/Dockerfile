FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/AdminTurnedDevOps/Quality-Code-With-Go/Chapter5
RUN cd /build && git clone https://github.com/AdminTurnedDevOps/Quality-Code-With-Go.git

RUN cd /build/Quality-Code-With-Go/Chapter5 && go build cloudstatuscheck.go
CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"

ENTRYPOINT [ "/build/Quality-Code-With-Go/Chapter5/cloudstatuscheck", "--azure" ]