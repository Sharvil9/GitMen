package plugins

type Plugin interface {
    Name() string
    Execute(args []string) error
}

var RegisteredPlugins = []Plugin{}

func RegisterPlugin(p Plugin) {
    RegisteredPlugins = append(RegisteredPlugins, p)
}