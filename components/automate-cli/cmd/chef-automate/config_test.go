package main

import (
	"os"
	"testing"
	dc "github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/api/config/shared"
	w "github.com/chef/automate/api/config/shared/wrappers"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var fqdn = "a2.test.com"

func TestIsConfigChanged(t *testing.T) {

	t.Run("check if the values are changed For Postgresql", func(t *testing.T) {
		reqConfig := PostgresqlConfig{
			Host: fqdn,
			Port: 80,
		}

		existingConfig := PostgresqlConfig{
			Host: "a5.test.com",
			Port: 22,
		}

		isChanged := isConfigChanged(existingConfig, reqConfig)
		assert.True(t, isChanged)

	})
	t.Run("check if the some values are added For Postgresql", func(t *testing.T) {
		reqConfig := PostgresqlConfig{
			Host:              fqdn,
			Port:              80,
			CheckpointTimeout: "10c",
		}

		existingConfig := PostgresqlConfig{
			Host: "a5.test.com",
		}

		isChanged := isConfigChanged(existingConfig, reqConfig)
		assert.True(t, isChanged)

	})
	t.Run("check if the no values are added or changed", func(t *testing.T) {
		reqConfig := PostgresqlConfig{
			Host:              fqdn,
			Port:              80,
			CheckpointTimeout: "10c",
		}

		existingConfig := PostgresqlConfig{
			Host:              fqdn,
			Port:              80,
			CheckpointTimeout: "10c",
		}

		isChanged := isConfigChanged(existingConfig, reqConfig)
		assert.False(t, isChanged)

	})

}

func TestPostgresSqlDecodeFromInput(t *testing.T) {
	t.Run("If there are some parameters", func(t *testing.T) {
		req := `[pg_dump]
enable = true
path = "/mnt/automate_backups/postgresql/pg_dump"
[replication]
lag_health_threshold = 20480
max_replay_lag_before_restart_s = 180
name = "replication"
password = "replication"`

		config, _ := getDecodedConfig(req, "postgresql")

		decodedConfig := config.(PostgresqlConfig)

		assert.Equal(t, decodedConfig.PgDump.Enable, true)
	})

}

func TestTomlFileCreateFromReqConfigLog(t *testing.T) {
	fileName := "logtoml.toml"
	req := dc.AutomateConfig{
		Global: &shared.GlobalConfig{
			V1: &shared.V1{
				Log: &shared.Log{
					RedirectSysLog:      w.Bool(true),
					RedirectLogFilePath: w.String("/var/tmp/"),
				},
			},
		},
	}

	createTomlFileFromConfig(req, fileName)
	assert.FileExists(t, fileName)
	os.Remove(fileName)
}

func TestErrorOnSelfManaged(t *testing.T) {
	testCases := []struct {
		isPostgresql bool
		isOpenSearch bool
		errorWant   error
	}{
		{
			true,
			false,
			errors.Errorf(ERROR_SELF_MANAGED_CONFIG_SHOW, "Postgresql"),
		},
		{
			false,
			true,
			errors.Errorf(ERROR_SELF_MANAGED_CONFIG_SHOW, "OpenSearch"),
		},
		{
			false,
			false,
			nil,
		},
	}

	for _, testCase := range testCases {
		errGot := errorOnSelfManaged(testCase.isPostgresql,testCase.isOpenSearch)
		if errGot == nil {
			assert.Equal(t, testCase.errorWant, errGot)
		} else {
			assert.EqualError(t, testCase.errorWant, errGot.Error())
		}
	}
}
