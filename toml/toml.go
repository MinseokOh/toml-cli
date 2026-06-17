package toml

import (
	"fmt"
	"strings"

	lib "github.com/pelletier/go-toml/v2"
)

// Toml is a struct
type Toml struct {
	path string
	out  string

	raw []byte

	tree map[string]any
}

// NewToml returns the Toml
func NewToml(path string) (Toml, error) {
	toml := Toml{path: path}

	if err := toml.readFile(); err != nil {
		return toml, err
	}

	if err := toml.load(); err != nil {
		return toml, err
	}

	return toml, nil
}

func (t *Toml) load() error {
	t.tree = map[string]any{}
	return lib.Unmarshal(t.raw, &t.tree)
}

// Dest set output given path
func (t *Toml) Out(path string) {
	t.out = path
}

// Get the value at key using dot notation.
func (t *Toml) Get(query string) any {
	keys := strings.Split(query, ".")
	var cur any = t.tree
	for _, k := range keys {
		m, ok := cur.(map[string]any)
		if !ok {
			return nil
		}
		cur, ok = m[k]
		if !ok {
			return nil
		}
	}
	return cur
}

// Has reports whether a value exists at key using dot notation.
func (t *Toml) Has(query string) bool {
	return t.Get(query) != nil
}

// Set the value at key using dot notation.
func (t *Toml) Set(query string, data any) error {
	if !t.Has(query) {
		return fmt.Errorf("not have key")
	}

	keys := strings.Split(query, ".")
	cur := t.tree
	for _, k := range keys[:len(keys)-1] {
		next, ok := cur[k].(map[string]any)
		if !ok {
			next = map[string]any{}
			cur[k] = next
		}
		cur = next
	}
	cur[keys[len(keys)-1]] = data
	return nil
}

// Merge merges another TOML file into this one
func (t *Toml) Merge(other *Toml) error {
	return t.mergeTree(t.tree, other.tree)
}

// mergeTree recursively merges source tree into target tree
func (t *Toml) mergeTree(target, source map[string]any) error {
	for key, sourceValue := range source {
		if targetValue, ok := target[key]; ok {
			// If both are tables (nested objects), merge recursively
			if sourceTable, ok := sourceValue.(map[string]any); ok {
				if targetTable, ok := targetValue.(map[string]any); ok {
					if err := t.mergeTree(targetTable, sourceTable); err != nil {
						return err
					}
					continue
				}
			}
		}

		// For all other cases (primitives, arrays, or new keys), overwrite
		target[key] = sourceValue
	}
	return nil
}
