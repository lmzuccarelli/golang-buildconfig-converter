package service

import (
	"fmt"
	"io/ioutil"

	"github.com/luigizuccarelli/golang-buildconfig-converter/pkg/schema"
	"gopkg.in/yaml.v2"
)

const (
	build           string = "Build"
	buildConfig     string = "BuildConfig"
	manifests       string = "/manifests/"
	dotYml          string = ".yaml"
	errMsgYaml      string = "converting struct to yaml for %s"
	errMsgUnmarshal string = "unmarshaling build config to struct %v for %s"
)

// mapToBuild function maps the BuildConfig to Build structure
// this is the first pass and needs to be verified
func mapToBuild(bc schema.BuildConfig) schema.Build {
	b := schema.Build{}
	b.Kind = build
	b.APIVersion = bc.APIVersion
	b.Metadata.Name = bc.Metadata.Name
	b.Spec.Source.URL = bc.Spec.Source.Git.URI
	b.Spec.Source.ContextDir = bc.Spec.Source.ContextDir
	b.Spec.Strategy.Kind = bc.Spec.Strategy.Type
	b.Spec.Dockerfile = bc.Spec.Strategy.DockerStrategy.DockerfilePath
	b.Spec.Output.Image = bc.Spec.Output.To.Kind
	return b
}

// readAllBuildConfigs - reads all the BuildConfigs from a given directory
func readAllBuildConfigs(dir string) ([]schema.BuildConfig, error) {
	var bcs []schema.BuildConfig
	var bc *schema.BuildConfig
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []schema.BuildConfig{}, err
	}
	for _, f := range files {
		if !f.IsDir() {
			yfile, err := ioutil.ReadFile(dir + "/" + f.Name())
			if err != nil {
				return []schema.BuildConfig{}, err
			}
			err = yaml.Unmarshal(yfile, &bc)
			if err != nil {
				return []schema.BuildConfig{}, fmt.Errorf(errMsgUnmarshal, err, f.Name())
			}
			// we are only interested in BuildConfigs
			if bc.Kind == buildConfig {
				bcs = append(bcs, *bc)
			}
		}
	}
	return bcs, nil
}

// readConfigFile - reads a config file to eliminate the need for complex cli parameters
func readConfigFile(file string) (*schema.Config, error) {
	var cfg *schema.Config
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		return &schema.Config{}, err
	}
	err = yaml.Unmarshal(yfile, &cfg)
	if err != nil {
		return &schema.Config{}, err
	}
	return cfg, nil
}

// convertBuildConfigs - takes the list of BuildConfigs, maps them
// to ashipwright Build struct, marshals to yaml and writes to disk
func convertBuildConfigs(cfg *schema.Config, bcs []schema.BuildConfig) error {
	for _, bc := range bcs {
		b := mapToBuild(bc)
		yml, err := yaml.Marshal(b)
		if err != nil {
			return fmt.Errorf(errMsgYaml, bc.Metadata.Name)
		}
		err = ioutil.WriteFile(cfg.Spec.WorkingDirectory+manifests+bc.Metadata.Name+dotYml, yml, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
