package sandbox

import (
	"net/http"

	"gopkg.in/vinxi/layer.v0"
)

type Plugin struct {
	Name        string
	Description string
}

// PluginLayer represents a plugins layer designed to intrument
// proxies providing plugin based dynamic configuration
// capabilities, such as register/unregister or
// enable/disable plugins at runtime satefy.
type PluginLayer struct {
	pool []Plugin
}

// NewPluginLayer creates a new plugins layer.
func NewPluginLayer() *PluginLayer {
	return &PluginLayer{}
}

// Register implements the middleware Register method.
func (l *PluginLayer) Register(mw *layer.Middleware) {
	mw.Use("error", l.run)
	mw.Use("request", l.run)
}

func (l *PluginLayer) run(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// no-op for now
		h.ServeHTTP(w, r)
	})
}
