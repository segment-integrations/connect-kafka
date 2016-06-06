# connect-kafka

Connect [Segment](https://segment.com/) to [Kafka](http://kafka.apache.org/) in 5 minutes. Pipe Segment's data sources into your Kafka cluster.

Release: `ALPHA`

<img src="http://hortonworks.com/wp-content/uploads/2016/03/kafka-logo-wide.png" data-canonical-src="http://hortonworks.com/wp-content/uploads/2016/03/kafka-logo-wide.png" width="200" height="105" />

## Features
`connect-kafka` is a simple server that you deploy in your infrastructure. It listens for Segment events and forwards them to the Kafka topic of your choice.

- Easily forward web, mobile, server analytics events to your Kafka instance
- Deploys in your infrastructure
- Supports any Kafka cluster 
- Built with [Heroku Kafka](https://www.heroku.com/kafka) support in mind (with public/private space support)
- Deploys in 5 minutes
- Allows SSL (or not) connections
- Supports all Segment standard methods (`identify`, `track`, `page`, `screen`, `group`)

## Quickstart

1. *Connect to Kafka* - connect the `connect-kafka` to your Kafka instance.
2. *Setup Webbook* - Enter connect-kafka's listen address into your Segment webhook menu.

### Connect to Kafka

You can connect to any internal Kafka deployment. 

#### Heroku Kafka

Below is an example to connect to a Heroku Kafka in a public space (via SSL):

```bash
go get -u github.com/segment-integrations/connect-kafka
heroku config:get KAFKA_URL -a kafka-integration-demo  # copy the kafka broker urls into command below
heroku config:get KAFKA_TRUSTED_CERT -a kafka-integration-demo > kafka_trusted_cert.cer
heroku config:get KAFKA_CLIENT_CERT -a kafka-integration-demo > kafka_client_cert.cer
heroku config:get KAFKA_CLIENT_CERT_KEY -a kafka-integration-demo > kafka_client_key_cert.cer
connect-kafka \
 --debug \
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
