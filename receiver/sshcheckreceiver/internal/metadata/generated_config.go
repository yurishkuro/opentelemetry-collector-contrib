// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for sshcheck metrics.
type MetricsConfig struct {
	SshcheckDuration     MetricConfig `mapstructure:"sshcheck.duration"`
	SshcheckError        MetricConfig `mapstructure:"sshcheck.error"`
	SshcheckSftpDuration MetricConfig `mapstructure:"sshcheck.sftp_duration"`
	SshcheckSftpError    MetricConfig `mapstructure:"sshcheck.sftp_error"`
	SshcheckSftpStatus   MetricConfig `mapstructure:"sshcheck.sftp_status"`
	SshcheckStatus       MetricConfig `mapstructure:"sshcheck.status"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SshcheckDuration: MetricConfig{
			Enabled: true,
		},
		SshcheckError: MetricConfig{
			Enabled: true,
		},
		SshcheckSftpDuration: MetricConfig{
			Enabled: false,
		},
		SshcheckSftpError: MetricConfig{
			Enabled: false,
		},
		SshcheckSftpStatus: MetricConfig{
			Enabled: false,
		},
		SshcheckStatus: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool            `mapstructure:"enabled"`
	Include []filter.Config `mapstructure:"include"`
	Exclude []filter.Config `mapstructure:"exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for sshcheck resource attributes.
type ResourceAttributesConfig struct {
	SSHEndpoint ResourceAttributeConfig `mapstructure:"ssh.endpoint"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		SSHEndpoint: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for sshcheck metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
