package server

// defaultExtension is a struct that contains the default values for
// an extension
var defaultExtension = &extensionBuilder{
	config: &ExtensionConfig{
		requiredMetadata: make(map[string]string),
		methods:          make(map[string]MethodFunc),
	},
	logFunc: func(l string) {},
}

// Builder creates a new ExtensionBuilder object.
func Builder() ExtensionBuilder {
	return defaultExtension
}

type extensionBuilder struct {
	config  *ExtensionConfig
	logFunc LoggerFunc
}

// ExtensionBuilder is the interface for creating an extension server
type ExtensionBuilder interface {
	// WithMethods specifies the methods that should be provided
	// by the extension
	WithMethods(map[string]MethodFunc) ExtensionBuilder
	// WithRequiredMetadata specifies the metadata that should be sent
	// with each request
	// TODO: this will be replaced by a metadata validation function
	WithRequiredMetadata(map[string]string) ExtensionBuilder
	// Named specifies the name of the extensions.
	Named(string) ExtensionBuilder
	// WithLoggerFunc specifies what should occur when a log is emitted.
	// By default, logs will not be emitted.
	WithLoggerFunc(LoggerFunc) ExtensionBuilder
	// Build creates the extensions
	Build() (*ExtensionServer, error)
}

func (b *extensionBuilder) Named(name string) ExtensionBuilder {
	b.config.name = name
	return b
}

func (b *extensionBuilder) WithRequiredMetadata(requiredMetadata map[string]string) ExtensionBuilder {
	b.config.requiredMetadata = requiredMetadata
	return b
}

func (b *extensionBuilder) WithMethods(methods map[string]MethodFunc) ExtensionBuilder {
	b.config.methods = methods
	return b
}

func (b *extensionBuilder) WithLoggerFunc(fn LoggerFunc) ExtensionBuilder {
	b.logFunc = fn
	return b
}

func (b *extensionBuilder) Build() (*ExtensionServer, error) {
	return &ExtensionServer{
		logFn: b.logFunc,
		extension: &Extension{
			conf: b.config,
		},
	}, nil
}
