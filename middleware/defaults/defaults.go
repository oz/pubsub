package defaults

import (
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/audit"
	"github.com/lileio/pubsub/middleware/logrus"
	"github.com/lileio/pubsub/middleware/opentracing"
	"github.com/lileio/pubsub/middleware/prometheus"
	"github.com/lileio/pubsub/middleware/recover"
)

// Middleware is a helper to import the default middleware for pubsub
var Middleware = []pubsub.Middleware{
	logrus.Middleware{},
	prometheus.Middleware{},
	opentracing.Middleware{},
	audit.Middleware{},
	recover.Middleware{},
}

// MiddlewareWithRecovery returns the default middleware but allows
// you to inject a function for dealing with panics
func MiddlewareWithRecovery(fn recover.RecoveryHandlerFunc) []pubsub.Middleware {
	return []pubsub.Middleware{
		logrus.Middleware{},
		prometheus.Middleware{},
		opentracing.Middleware{},
		audit.Middleware{},
		recover.Middleware{RecoveryHandlerFunc: fn},
	}
}
