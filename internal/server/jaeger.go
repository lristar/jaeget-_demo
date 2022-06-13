package demo

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func NewJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.195.223:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return tracer, closer, err
	}
	opentracing.SetGlobalTracer(tracer) // StartspanFromContext创建新span时会用到
	return tracer, closer, nil
}
