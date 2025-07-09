package toml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewToml(t *testing.T) {

	_, err := NewToml("./config.toml")
	require.NotNil(t, err)

	toml, err := NewToml("../sample/get-set/app.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)
}

func TestGet(t *testing.T) {
	toml, err := NewToml("../sample/get-set/app.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)

	query := "app.name"
	appName := toml.Get(query)

	require.Equal(t, appName, "demo-app")

	query = "server.port"
	port := toml.Get(query)

	require.Equal(t, port, int64(8080))

	query = "nonexistent.key"
	res := toml.Get(query)
	require.Nil(t, res)
}

func TestSet(t *testing.T) {
	toml, err := NewToml("../sample/get-set/app.toml")
	require.Nil(t, err)
	require.NotNil(t, toml)

	toml.Out("../sample/get-set/app_test_output.toml")

	query := "app.name"
	value := "test-app"

	err = toml.Set(query, value)
	require.Nil(t, err)

	toml.Write()

	// Verify the change
	result := toml.Get(query)
	require.Equal(t, result, "test-app")
}

func TestMerge(t *testing.T) {
	base, err := NewToml("../sample/merge/base.toml")
	require.Nil(t, err)
	require.NotNil(t, base)

	override, err := NewToml("../sample/merge/override.toml")
	require.Nil(t, err)
	require.NotNil(t, override)

	base.Out("../sample/merge/base_test_output.toml")

	// Merge override into base
	err = base.Merge(&override)
	require.Nil(t, err)

	// Verify merged values
	// Override should win
	host := base.Get("server.host")
	require.Equal(t, host, "0.0.0.0")

	port := base.Get("server.port")
	require.Equal(t, port, int64(3000))

	// Base values should remain if not overridden
	timeout := base.Get("server.timeout")
	require.Equal(t, timeout, int64(30))

	// New sections should be added
	logLevel := base.Get("logging.level")
	require.Equal(t, logLevel, "info")

	base.Write()
}
