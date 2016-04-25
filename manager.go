package sandbox

import (
	"net/http"

	"gopkg.in/vinxi/layer.v0"
	"gopkg.in/vinxi/vinxi.v0"
)

type Layer struct {
	layer *layer.Layer
}

type Options struct {
	Optional bool
}

type Rule interface {
	Name() string
	Description() string
	Options() Options
	JSONConfig() string
	Match(*http.Request) bool
}

type Scope struct {
	disabled bool
	rules    []Rule
	plugins  *PluginLayer

	Name        string
	Description string
}

func NewScope() *Scope {
	return &Scope{Plugins: NewPluginLayer()}
}

func (s *Scope) AddRule(rules ...Rule) {
	s.rules = append(s.rules, rules...)
}

func (s *Scope) Rules() []Rule {
	return s.rules
}

func (s *Scope) Disable() {
	s.disabled = true
}

func (s *Scope) Enable() {
	s.disabled = false
}

func (s *Scope) HandleHTTP(h http.Handler) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if s.disabled {
			h.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(502)
		w.Write([]byte("bad request"))
	}
}

func NewInstance(vinxi *vinxi.Vinxi) *Instance {
	plugins := NewPluginLayer()
	plugins.Register(vinxi)
	return &Instance{plugins: plugins, vinxi: vinxi}
}

type Manager struct {
	Server   *http.Server
	instance *vinxi.Vinxi
	scopes   []*Scope
}

func Manage(instance *vinxi.Vinxi) *Manager {
	m := &Manager{instance: instance}
	instance.Layer.UsePriority("request", layer.Tail, m)
	return m
}

func (a *Manager) HandleHTTP(w http.ResponseWriter, req *http.Request, h http.Handler) {
	next := h

	for _, scope := range scopes {
		next = http.HandlerFunc(scope.HandleHTTP(next))
	}

	next.ServeHTTP(w, r)
}

func (a *Manager) Listen(opts ServerOptions) error {
	a.Server = NewServer(opts)
	return server.ListenAndServe()
}
