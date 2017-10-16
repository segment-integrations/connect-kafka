# connect-kafka

This program is an example implementation of a [Segment](https://segment.com/) [Webhook](https://segment.com/docs/integrations/webhooks) consumer that publishes events to [Kafka](http://kafka.apache.org/).

This is not an officially supported Segment product, but is meant to demonstrate a simple server that you can fork or emulate to route Segment data to your internal systems. It may even suit your needs as is! 

<img src="http://hortonworks.com/wp-content/uploads/2016/03/kafka-logo-wide.png" data-canonical-src="http://hortonworks.com/wp-content/uploads/2016/03/kafka-logo-wide.png" width="200" height="105" />

## Features
`connect-kafka` is a simple server that you deploy in your infrastructure and expose to the internet. It listens for Segment events and forwards them to the Kafka topic of your choice.

- Easily forward web, mobile, server analytics events to your Kafka instance
- Deploys in your infrastructure
- Supports any Kafka cluster
- Built with [Heroku Kafka](https://www.heroku.com/kafka) support in mind (with public/private space support)
- Supports SSL (or not) connections to your cluster
- Supports all Segment standard methods (`identify`, `track`, `page`, `screen`, `group`)

## Quickstart

1. *Connect to Kafka* - connect the `connect-kafka` to your Kafka instance.
2. *Setup Webbook* - Enter connect-kafka's listen address into your Segment webhook menu.

## FAQ

#### Does this support shared secret authentication? 

Not yet, though we'd love a contribution that adds it! 

#### How do Segment Webhooks behave if my server goes down?

We will retry the requests to the server 5 times over an hour if your server becomes unavailable.

#### Will the events arrive in order?

Because we're dealing with unbounded streaming data, we can't guarantee that your events arrive in the absolute order that they were collected in your client devices. As such, we recommend using the `timestamp` fields on each message with event-time windowing approaches in your destinations and streaming data applications.

### Connect to Kafka

Download `connect-kafka` using curl:

```bash
curl -s http://connect.segment.com/install-connect-kafka.sh | sh
```

If you just want the binary and install it yourself:

```bash
http://connect.segment.com/connect-kafka-darwin-amd64
```

You can also use Docker:

```bash
make docker
docker run segment/connect-kafka [...]
```

You can connect to any internal Kafka deployment.

```
$ connect-kafka -h

Usage:
  connect-kafka
    --topic=<topic>
    --broker=<url>...
    [--trusted-cert=<path> --client-cert=<path> --client-cert-key=<path>]
  connect-kafka -h | --help
  connect-kafka --version

Options:
  -h --help                   Show this screen
  --version                   Show version
  --topic=<topic>             Kafka topic name
  --broker=<url>              Kafka broker URL
```

#### Heroku Kafka

Below is an example to connect to a Heroku Kafka in a public space (via SSL):

```bash
go get -u github.com/segment-integrations/connect-kafka
heroku config:get KAFKA_URL -a kafka-integration-demo  # copy the kafka broker urls into command below
heroku config:get KAFKA_TRUSTED_CERT -a kafka-integration-demo > kafka_trusted_cert.cer
heroku config:get KAFKA_CLIENT_CERT -a kafka-integration-demo > kafka_client_cert.cer
heroku config:get KAFKA_CLIENT_CERT_KEY -a kafka-integration-demo > kafka_client_key_cert.cer
connect-kafka \
 --topic=segment \
 --broker=kafka+ssl://ec2-51-16-10-109.compute-1.amazonaws.com:9096 \
 --broker=kafka+ssl://ec2-62-7-61-181.compute-1.amazonaws.com:9096 \
 --broker=kafka+ssl://ec2-33-20-240-35.compute-1.amazonaws.com:9096 \
 --trusted-cert=kafka_trusted_cert.cer \
 --client-cert=kafka_client_cert.cer \
 --client-cert-key=kafka_client_key_cert.cer
 ```

### Setup Webhook

1. Go to the Segment.com and select the source you want to connect to Kafka
2. Add your `connect-kafka` server's address to the webhook integration's settings.

![](http://g.recordit.co/XcyIz2fqJv.gif)


## Testing

### via localtunnel

You can open up a localtunnel on your local machine while you're testing:

```
npm install -g localtunnel
lt --port 3000
```

Enter the resulting localtunnel url as the Segment webhook with `/listen` appended, like: `https://aqjujyhnck.localtunnel.me/listen`

## License

 MIT
