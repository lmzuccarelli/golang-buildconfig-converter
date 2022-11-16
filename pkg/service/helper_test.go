package service

import (
	"errors"
	"testing"

	"github.com/luigizuccarelli/golang-buildconfig-converter/pkg/schema"
	"github.com/stretchr/testify/assert"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Inject (force) readAll test error")
}

func TestReadConfigFile(t *testing.T) {
	t.Run("Testing readConfigFile : should pass", func(t *testing.T) {
		testCfg := &schema.Config{
			APIVersion: "0.0.1",
			Kind:       "Config",
			Metadata: schema.MetaDataSchema{
				Name: "converter-config",
			},
			Spec: schema.ConfigSpecSchema{
				WorkingDirectory: "../../tests/working-dir",
				BuildConfigPath:  "../../tests/good-bc",
				LogLevel:         "info",
			},
		}
		cfg, err := readConfigFile("../../tests/config.yaml")
		if err != nil {
			t.Fatalf("Should not fail : found error %v", err)
		}
		assert.Equal(t, cfg, testCfg, "Config structs should be the same")
	})

	t.Run("Testing readConfigFile (wrong file name): should fail", func(t *testing.T) {
		_, err := readConfigFile("../../tests/conf.yaml")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})

	t.Run("Testing readConfigFile (bad yaml): should fail", func(t *testing.T) {
		_, err := readConfigFile("../../tests/config-bad.yaml")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})
}

func TestReadAllBuildConfigs(t *testing.T) {
	t.Run("Testing readAllBuildConfigs : should pass", func(t *testing.T) {
		bcs, err := readAllBuildConfigs("../../tests/good-bc")
		if err != nil {
			t.Fatalf("Should not fail : found error %v", err)
		}
		assert.Equal(t, 1, len(bcs), "Should have only one buildconfig")
	})

	t.Run("Testing readAllBuildConfigs (wrong file name): should fail", func(t *testing.T) {
		_, err := readAllBuildConfigs("../../tests/none")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})

	t.Run("Testing readConfigFile (bad yaml): should fail", func(t *testing.T) {
		_, err := readAllBuildConfigs("../../tests/bad-bc")
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})
}

func TestConverBuildConfigs(t *testing.T) {

	t.Run("Testing convertBuildConfigs : should pass", func(t *testing.T) {
		cfg, _ := readConfigFile("../../tests/config.yaml")
		bcs, _ := readAllBuildConfigs("../../tests/good-bc")
		err := convertBuildConfigs(cfg, bcs)
		if err != nil {
			t.Fatalf("Should not fail : found error %v", err)
		}
		assert.Equal(t, 1, len(bcs), "Should have only one buildconfig")
	})

	t.Run("Testing convertBuildConfigs : should fail", func(t *testing.T) {
		cfg, _ := readConfigFile("../../tests/config-wrong-dirs.yaml")
		bcs, _ := readAllBuildConfigs("../../tests/good-bc")
		err := convertBuildConfigs(cfg, bcs)
		if err == nil {
			t.Fatalf("Should fail : found error %v", err)
		}
	})
}
