package toml

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func (t *Toml) String() string {
	fmt.Println(t.tokens)
	var sb strings.Builder

	if len(t.tokens) == 1 {
		return ""
	}

	for i := 0; i < len(t.tokens); i++ {
		cur := t.tokens[i]
		next := t.tokens[i+1]

		if cur.typ == tokenString {
			sb.WriteString("\"")
			sb.WriteString(cur.val)
			sb.WriteString("\"")
		} else {
			sb.WriteString(cur.val)
		}

		if next.typ == tokenEOF {
			return sb.String()
		}

		line := next.Line - cur.Line
		for j := 0; j < line; j++ {
			sb.WriteString("\n")
		}

		// for new line with space
		if line > 0 && next.Col > 1 {
			for j := 1; j < next.Col; j++ {
				sb.WriteString(" ")
			}
		}

		cs := cur.Col
		ns := next.Col

		// do for token string
		if cur.typ == tokenString {
			cs++
		}

		if next.typ == tokenString {
			ns--
		}

		col := ns - cs - len(cur.val)

		for j := 0; j < col; j++ {
			sb.WriteString(" ")
		}
	}

	return ""
}

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
