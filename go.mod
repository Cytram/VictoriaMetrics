module github.com/VictoriaMetrics/VictoriaMetrics

require (
	cloud.google.com/go v0.71.0 // indirect
	cloud.google.com/go/storage v1.12.0
	github.com/VictoriaMetrics/fastcache v1.5.7

	// Do not use the original github.com/valyala/fasthttp because of issues
	// like https://github.com/valyala/fasthttp/commit/996610f021ff45fdc98c2ce7884d5fa4e7f9199b
	github.com/VictoriaMetrics/fasthttp v1.0.7
	github.com/VictoriaMetrics/metrics v1.12.3
	github.com/VictoriaMetrics/metricsql v0.7.2
	github.com/aws/aws-sdk-go v1.35.20
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/go-kit/kit v0.10.0
	github.com/golang/snappy v0.0.2
	github.com/klauspost/compress v1.11.2
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/prometheus/prometheus v1.8.2-0.20200911110723-e83ef207b6c2
	github.com/valyala/fastjson v1.6.1
	github.com/valyala/fastrand v1.0.0
	github.com/valyala/fasttemplate v1.2.1
	github.com/valyala/gozstd v1.8.3
	github.com/valyala/histogram v1.1.2
	github.com/valyala/quicktemplate v1.6.3
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201102192140-3288bc1ea189 // indirect
	google.golang.org/api v0.34.0
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201102152239-715cce707fb0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

go 1.13
