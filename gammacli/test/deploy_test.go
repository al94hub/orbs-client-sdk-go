package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeployCounter(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("deploy", "./_counter/contract.go", "-name", "CounterExample")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = cli.Run("run-query", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "0"`))

	out, err = cli.Run("send-tx", "counter-add.json")
	t.Log(out)
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "previous count is 0"`))

	out, err = cli.Run("run-query", "counter-get.json")
	t.Log(out)
	require.NoError(t, err, "get should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))
	require.True(t, strings.Contains(out, `"Value": "25"`))
}

func TestDeployCorruptContract(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("deploy", "./_corrupt/corrupt.go", "-name", "CounterExample")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`))
	require.True(t, strings.Contains(out, `compilation of deployable contract 'CounterExample' failed`))
}

func TestDeployOfAlreadyDeployed(t *testing.T) {
	cli := GammaCli().DownloadLatestGammaServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("deploy", "./_counter/contract.go", "-name", "CounterExample")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "SUCCESS"`))

	out, err = cli.Run("deploy", "./_counter/contract.go", "-name", "CounterExample")
	t.Log(out)
	require.NoError(t, err, "deploy should succeed")
	require.True(t, strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`))
	require.True(t, strings.Contains(out, `contract already deployed`))
}

func TestRunMethodWithoutDeploy(t *testing.T) {
	cli := GammaCli().WithExperimentalServer().StartGammaServer()
	defer cli.StopGammaServer()

	out, err := cli.Run("send-tx", "counter-add.json")
	t.Log(out)
	require.NoError(t, err, "add should succeed")
	require.True(t, strings.Contains(out, `"RequestStatus": "SYSTEM_ERROR"`))
	require.True(t, strings.Contains(out, `"ExecutionResult": "ERROR_UNEXPECTED"`))
}
