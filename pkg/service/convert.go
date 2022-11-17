package service

import (
	"log"
	"os"
)

const (
	msgDelete string = "INFO: deleting working directory %s"
	msgCreate string = "INFO: creating workig directory %s"
	dir       string = "/manifests"
)

// Convert function that reads BuildConfigs from a directory and converts them to
// shipwright build manifests
func Convert(config string) error {

	cfg, err := readConfigFile(config)
	if err != nil {
		return err
	}
	log.Printf(msgDelete, cfg.Spec.WorkingDirectory)
	os.RemoveAll(cfg.Spec.WorkingDirectory)
	log.Printf(msgCreate, cfg.Spec.WorkingDirectory)
	os.MkdirAll(cfg.Spec.WorkingDirectory+dir, os.ModePerm)

	bcs, err := readAllBuildConfigs(cfg.Spec.BuildConfigPath)
	if err != nil {
		return err
	}

	err = convertBuildConfigs(cfg, bcs)
	if err != nil {
		return err
	}
	return nil
}
