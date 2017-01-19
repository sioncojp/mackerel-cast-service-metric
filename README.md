# mackerel-cast-service-metric
Some results of mackerel-agent-plugin casts ServiceMetric.

# Usage

## Setting toml file.
```toml
api_key = "xxxxxxxxxxxxx"

[[rule]]
cmd = "mackerel-plugin-linux"
service_name = "web"
metric_name = ""

[[rule]]
cmd = "mackerel-plugin-uptime"
service_name = "web"
metric_name = "test"
```

## Run mackerel-cast-service-metric.
```shell
$ go get github.com/sioncojp/mackerel-cast-service-metric
$ glide install
$ go run /mackerel-cast-service-metric -c config.toml
```
