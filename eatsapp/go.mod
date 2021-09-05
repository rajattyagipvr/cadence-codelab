module trying

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02

require (
	github.com/apache/thrift v0.13.0
	github.com/cristalhq/jwt/v3 v3.1.0
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/m3db/prometheus_client_golang v0.8.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v1.2.1
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.7.0
	github.com/uber-go/tally v3.3.17+incompatible
	github.com/uber/cadence v0.22.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/tchannel-go v1.16.0
	go.temporal.io/sdk v1.9.0
	go.uber.org/atomic v1.7.0
	go.uber.org/cadence v0.18.2
	go.uber.org/goleak v1.0.0
	go.uber.org/thriftrw v1.25.0
	go.uber.org/yarpc v1.53.2
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20210420210106-798c2154c571
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	gopkg.in/yaml.v2 v2.4.0
)
