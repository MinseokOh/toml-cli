package toml

import (
	"fmt"

	lib "github.com/pelletier/go-toml"
)

// Toml is a struct
type Toml struct {
	path string
	out string

	raw    []byte

	tree *lib.Tree
}

// NewToml returns the Toml
func NewToml(path string) (Toml, error) {
	toml := Toml{
		path: path,
	}

	err := toml.readFile()
	if err != nil {
		return toml, err
	}

	err = toml.load()
	if err != nil {
		return toml, err
	}

	return toml, nil
}

func (t *Toml) load() error {
	var err error

	t.tree, err = lib.LoadBytes(t.raw)

	if err != nil {
		return err
	}

	return nil
}

// Dest set output given path
func (t *Toml) Out(path string) {
	t.out = path
}

// Get the value at key in the Tree.
// [Wrapped function go-toml.]
func (t *Toml) Get(query string) interface{} {
	return t.tree.Get(query)
}

// Set the value at key in the Tree.
// [Wrapped function go-toml.]
func (t *Toml) Set(query string, data interface{}) error {
	if !t.tree.Has(query) {
		return fmt.Errorf("not have key")
	}

	t.tree.Set(query, data)
	return nil
}
