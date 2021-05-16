package toml

import (
	"fmt"
	"io/ioutil"

	lib "github.com/pelletier/go-toml"
)

type Toml struct {
	path string
	dest string
	raw []byte
	tree *lib.Tree
}

func NewToml(path string) (Toml, error) {
	toml := Toml{
		path : path,
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

func (t *Toml) Dest(path string) {
	t.dest = path
}

func (t *Toml) Get(query string) interface{} {
	return t.tree.Get(query)
}

func (t *Toml) Set(query string, data interface{}) error {
	if !t.tree.Has(query) {
		return fmt.Errorf("not have key")
	}

	t.tree.Set(query, data)
	return nil
}

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