module github.com/lamassuiot/lamassu-cloud-proxy

go 1.16

replace github.com/lamassuiot/lamassu-aws-connector => /home/ikerlan/lamassu/lamassu-aws-2/aws-connector

require (
	github.com/cloudevents/sdk-go/v2 v2.6.0
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.10.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lamassuiot/lamassu-aws-connector v0.0.1
	github.com/lamassuiot/lamassu-ca v1.0.10
	github.com/lib/pq v1.10.4
	github.com/opentracing/opentracing-go v1.2.0
	github.com/streadway/amqp v1.0.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
)
