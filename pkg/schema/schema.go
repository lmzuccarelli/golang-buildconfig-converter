package schema

// BuildConfig struct for the oopenshift buildconfig manifest.
// This is a naive implementation, as in reality the
// actual api from the buildconfig project should be referenced here
// As this was a POC I'm not hassled :) - please feel free to change
type BuildConfig struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		RunPolicy string `yaml:"runPolicy"`
		Triggers  []struct {
			Type   string `yaml:"type"`
			Github struct {
				Secret string `yaml:"secret"`
			} `yaml:"github,omitempty"`
			Generic struct {
				Secret string `yaml:"secret"`
			} `yaml:"generic,omitempty"`
		} `yaml:"triggers"`
		Source struct {
			Git struct {
				URI string `yaml:"uri"`
			} `yaml:"git"`
			ContextDir string `yaml:"contextDir"`
			Images     []struct {
				From struct {
					Kind string `yaml:"kind"`
					Name string `yaml:"name"`
				} `yaml:"from"`
				Paths []struct {
					SourcePath     string `yaml:"sourcePath"`
					DestinationDir string `yaml:"destinationDir"`
				} `yaml:"paths"`
			} `yaml:"images"`
		} `yaml:"source"`
		Strategy struct {
			Type           string `yaml:"type"`
			SourceStrategy struct {
				From struct {
					Kind string `yaml:"kind"`
					Name string `yaml:"name"`
				} `yaml:"from"`
			} `yaml:"sourceStrategy"`
			DockerStrategy struct {
				ImageOptimizationPolicy string `yaml:"imageOptimizationPolicy"`
				DockerfilePath          string `yaml:"dockerfilePath"`
				From                    struct {
					Kind string `yaml:"kind"`
					Name string `yaml:"name"`
				} `yaml:"from"`
			} `yaml:"dockerStrategy"`
		} `yaml:"strategy"`
		Output struct {
			To struct {
				Kind string `yaml:"kind"`
				Name string `yaml:"name"`
			} `yaml:"to"`
		} `yaml:"output"`
		PostCommit struct {
			Script string `yaml:"script"`
		} `yaml:"postCommit"`
	} `yaml:"spec"`
}

// Build struct for the shipwight build manifest.
// This is a naive implementation, as in reality the
// actual api from the build project should be referenced here
// As this was a POC I'm not hassled :) - please feel free to change
type Build struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Source struct {
			URL        string `yaml:"url"`
			ContextDir string `yaml:"contextDir"`
		} `yaml:"source"`
		Strategy struct {
			Name string `yaml:"name"`
			Kind string `yaml:"kind"`
		} `yaml:"strategy"`
		Dockerfile string `yaml:"dockerfile"`
		Output     struct {
			Image       string `yaml:"image"`
			Credentials struct {
				Name string `yaml:"name"`
			} `yaml:"credentials"`
		} `yaml:"output"`
	} `yaml:"spec"`
}

type MetaDataSchema struct {
	Name string `yaml:"name"`
}

// ConfigSpecSchema specific to config
type ConfigSpecSchema struct {
	WorkingDirectory string `yaml:"workingDirectory"`
	BuildConfigPath  string `yaml:"buildConfigPath"`
	LogLevel         string `yaml:"logLevel"`
}

// Config struct to simplify the use of complex cli input parameters
type Config struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   MetaDataSchema
	Spec       ConfigSpecSchema
}
