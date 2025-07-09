package toml

import (
	"os"
)

func (t *Toml) readFile() error {
	var err error
	t.raw, err = os.ReadFile(t.path)
	return err
}

// Write edited toml tree given path.
// if dest is not setted, overwrite it.
func (t *Toml) Write() error {
	path := t.out
	if path == "" {
		path = t.path
	}

	toml, err := t.tree.ToTomlString()
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(toml), 0644)
}
