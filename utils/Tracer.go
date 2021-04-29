package utils

import (
	"Backend-Loans/domain/dto"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Trace(serviceName string, c *gin.Context, status int, response dto.Response) {
	tracer, closer := InitJaeger(serviceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	clientSpan := tracer.StartSpan("Request_Data")
	defer clientSpan.Finish()

	b, _ := json.Marshal(response)
	bytes := []byte(b)

	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, c.Request.URL.Path)
	ext.HTTPMethod.Set(clientSpan, c.Request.Method)
	ext.HTTPStatusCode.Set(clientSpan, uint16(status))
	ext.Error.Set(clientSpan, false)
	ext.Component.Set(clientSpan, BytesToString(bytes))

	tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
