FROM ubuntu

COPY target/connect-kafka-linux-amd64 /connect-kafka

ENTRYPOINT ["/connect-kafka"]
