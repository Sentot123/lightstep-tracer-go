package lightstep

import (
	"golang.org/x/net/context"

	opentracing "github.com/opentracing/opentracing-go"
)

var logError = NewLogOnError()

// Flush forces a synchronous Flush.
func Flush(ctx context.Context, tracer opentracing.Tracer) {
	lsTracer, ok := tracer.(Tracer)
	if !ok {
		logError(newErrNotLisghtepTracer(tracer))
		return
	}
	lsTracer.Flush(ctx)
}

// CloseTracer synchronously flushes the tracer, then terminates it.
func Close(ctx context.Context, tracer opentracing.Tracer) {
	lsTracer, ok := tracer.(Tracer)
	if !ok {
		logError(newErrNotLisghtepTracer(tracer))
		return
	}
	lsTracer.Close(ctx)
}

// GetLightStepAccessToken returns the currently configured AccessToken.
func GetLightStepAccessToken(tracer opentracing.Tracer) (string, error) {
	lsTracer, ok := tracer.(Tracer)
	if !ok {
		return "", newErrNotLisghtepTracer(tracer)
	}

	return lsTracer.Options().AccessToken, nil
}

// DEPRECATED: use Flush instead.
func FlushLightStepTracer(lsTracer opentracing.Tracer) error {
	tracer, ok := lsTracer.(Tracer)
	if !ok {
		return newErrNotLisghtepTracer(lsTracer)
	}

	tracer.Flush(context.Background())
	return nil
}

// DEPRECATED: use Close instead.
func CloseTracer(tracer opentracing.Tracer) error {
	lsTracer, ok := tracer.(Tracer)
	if !ok {
		return newErrNotLisghtepTracer(tracer)
	}

	lsTracer.Close(context.Background())
	return nil
}
