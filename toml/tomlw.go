package toml

import (
	"io/ioutil"
)

func (t *Toml) readFile() error {
	var err error
	t.raw, err = ioutil.ReadFile(t.path)
	if err != nil {
		return err
	}

	return nil
}

// Write edited toml tree given path.
// if dest is not setted, overwrite it.
func (t *Toml) Write() error {
	var err error
	var toml string

	path := t.out
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
