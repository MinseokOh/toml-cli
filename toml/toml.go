package toml

import (
	"fmt"
	"io/ioutil"

	lib "github.com/pelletier/go-toml"
)

// Toml is a struct
type Toml struct {
	path string
	dest string
	raw  []byte
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

func (t *Toml) readFile() error {
	var err error
	t.raw, err = ioutil.ReadFile(t.path)
	if err != nil {
		return err
	}

	return nil
}

func (t *Toml) load() error {
	var err error

	t.tree, err = lib.Load(string(t.raw))
	if err != nil {
		return err
	}

	return nil
}

// Dest set output given path
func (t *Toml) Dest(path string) {
	t.dest = path
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

// Write edited toml tree given path.
// if dest is not setted, overwrite it.
func (t *Toml) Write() error {
	var err error
	var toml string

	path := t.dest
	if path == "" {
		path = t.path
	}

	toml, err = t.tree.ToTomlString()

	err = ioutil.WriteFile(path, []byte(toml), 0644)
	if err != nil {
		return err
	}

	return nil
}
