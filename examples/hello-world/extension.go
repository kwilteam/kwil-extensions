package helloworld

import (
	"github.com/kwilteam/kwil-extensions/server"
	"github.com/kwilteam/kwil-extensions/types"
)

var helloWorldExtension = server.ExtensionConfig{}

var config = map[string]string{
	"greeting":    "Hello", // this is a default value
	"punctuation": "",      // this is a required value
}

// no required metadata for this extension
var requiredMetadata = map[string]string{}

var helloMethod = server.Method{
	Name: "say_hello",
	RequiredInputs: []types.ScalarType{
		types.ScalarType_STRING, // the name of the person to greet
	},
	Function: func(inputs []*types.ScalarValue, metadata map[string]string) ([]*types.ScalarValue, error) {},
}
