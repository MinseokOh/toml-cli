package toml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewToml(t *testing.T) {

	_, err := NewToml("./config.toml")
	require.NotNil(t, err)

	toml, err := NewToml("../sample/example.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)
}

func TestGet(t *testing.T) {
	toml, err := NewToml("../sample/example.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)

	query := "title"
	title := toml.Get(query)

	require.Equal(t, title, "TOML Example")

	query = "database.server"
	server := toml.Get(query)

	require.Equal(t, server, "192.168.1.1")

	query = "aaaaaaaa"
	res := toml.Get(query)
	require.Nil(t, res)
}

func TestSet(t *testing.T) {
	toml, err := NewToml("../sample/example.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)

	toml.Dest("../sample/example_out.toml")

	query := "title"
	value := "my output"

	err = toml.Set(query, value)
	require.Nil(t, err)

	toml.Write()
}
