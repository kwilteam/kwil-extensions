package server

import (
	"fmt"

	"github.com/kwilteam/kwil-extensions/types"
)

type ExtConf struct {
	ConfigFunc       func(map[string]string) error
	RequiredMetadata map[string]string
	Methods          []func([]*types.ScalarValue, map[string]string) ([]*types.ScalarValue, error)
}

// ExtensionConfig is an application that is meant to extend the functionality
// of Kwil.  It is a containerized service that is meant to be able to
// execute arbitrary code.
type ExtensionConfig struct {
	// Configs is a map of configuration values that are passed to the
	// extension when it is started. It maps the name of the configuration
	// to the default value of the configuration.
	// If a default value is not provided, the Kwil node will be required to
	// provide a value before the extension can be started.
	Configs map[string]string

	// Methods is a list of methods that are provided by the extension.
	Methods []*Method

	// RequiredMetadata is a list of metadata keys that are required to be
	// included with each execution request.
	// These are typically values that would be defined when a Kwil application
	// is created.
	// It maps the name of the metadata to the default value of the metadata.
	// If a default value is not provided, the Kwil node will be required to
	// provide a value for each execution request.
	// e.g., for an extension that pulls ERC20 token balances, the required
	// metadata would be the address of the token.
	RequiredMetadata map[string]string
}

// A Method is a function that is provided by an extension.
type Method struct {
	// Name is the unique name of the method.
	// It will automatically be lowercased.
	Name string

	// Function is the function that is executed when the method is called.
	// It is provided the inputs to the method as a list of ScalarValues.
	// It is expected to return an array of ScalarValues.
	Function func(inputs []*types.ScalarValue, metadata map[string]string) ([]*types.ScalarValue, error)
}

// canExecute checks if the provided inputs are valid for the method.
func (m *Method) canExecute(inputs []*types.ScalarValue) error {
	if len(inputs) != len(m.RequiredInputs) {
		return fmt.Errorf("invalid number of inputs provided")
	}

	for i, requiredType := range m.RequiredInputs {
		if inputs[i].Type != requiredType {
			return fmt.Errorf("invalid input type: argument position %d received value of type %s, expected %s", i, inputs[i].Type, requiredType)
		}
	}

	return nil
}

type UserExtension interface {
	Configure(map[string]string) error
}
