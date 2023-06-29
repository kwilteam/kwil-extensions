package extension

import (
	"fmt"
	"github.com/kwilteam/kwil-extensions/types"
)

type ErrInvalidArgumentType struct {
	Expect types.ScalarType
	Got    types.ScalarType
	Pos    int
}

func (e ErrInvalidArgumentType) Error() string {
	return fmt.Sprintf("expect %s for arg #%d, got %s", e.Expect, e.Pos, e.Got)
}

type ErrInvalidArgumentNum struct {
	Expect int
	Got    int
}

func (e ErrInvalidArgumentNum) Error() string {
	return fmt.Sprintf("expect %d args, got %d", e.Expect, e.Got)
}

type ErrInvalidReturnType struct {
	Expect types.ScalarType
	Got    types.ScalarType
	Pos    int
}

func (e ErrInvalidReturnType) Error() string {
	return fmt.Sprintf("expect %s for returned #%d, got %s", e.Expect, e.Pos, e.Got)
}

type ErrInvalidReturnNum struct {
	Expect int
	Got    int
}

func (e ErrInvalidReturnNum) Error() string {
	return fmt.Sprintf("expect %d returned, got %d", e.Expect, e.Got)
}
