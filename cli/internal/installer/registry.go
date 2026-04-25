package installer

// Registry resolves an installer by canonical tool name.
type Registry struct {
	byName map[string]Installer
}

// Get returns the installer for name, or false if no such installer is
// registered.
func (r *Registry) Get(name string) (Installer, bool) {
	inst, ok := r.byName[name]
	return inst, ok
}

// Names returns all registered tool names.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.byName))
	for n := range r.byName {
		names = append(names, n)
	}
	return names
}

// DefaultRegistry returns a Registry with every supported tool registered.
func DefaultRegistry() *Registry {
	installers := []Installer{
		&Claude{},
		&Codex{},
		&Copilot{},
		&Gemini{},
	}
	r := &Registry{byName: make(map[string]Installer, len(installers))}
	for _, inst := range installers {
		r.byName[inst.Name()] = inst
	}
	return r
}
