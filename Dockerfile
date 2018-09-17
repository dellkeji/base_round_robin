FROM golang:1.9.0

WORKDIR /usr/src/round_robin_with_weight

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY service-initial.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/service-initial.sh

ADD . .
RUN make release