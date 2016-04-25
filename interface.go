package sandbox

import "net/http"

// Plugin represents the required interface implemented by plugins.
type Plugin interface {
	// Name is used to retrieve the plugin name identifier.
	Name() string
	// Description is used to retrieve a human friendly
	// description of what the plugin does.
	Description() string
	// Enable is used to enable the current plugin.
	// If the plugin has been already enabled, the call is no-op.
	Enable()
	// Disable is used to disable the current plugin.
	Disable()
	// Remove is used to disable and remove a plugin.
	Remove()
	// IsEnabled is used to check if a plugin is enabled or not.
	IsEnabled()
	// Run is used to run the plugin task.
	// Note: add erro reporting layer
	Run(http.Handler) http.Handler
}
