package command

import (
	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/metrics"
	metricsFreebsd "github.com/mackerelio/mackerel-agent/metrics/freebsd"
	"github.com/mackerelio/mackerel-agent/spec"
	specFreebsd "github.com/mackerelio/mackerel-agent/spec/freebsd"
)

func specGenerators() []spec.Generator {
	return []spec.Generator{
		&specFreebsd.KernelGenerator{},
		&specFreebsd.MemoryGenerator{},
		&specFreebsd.CPUGenerator{},
		&spec.FilesystemGenerator{},
	}
}

func interfaceGenerator() spec.InterfaceGenerator {
	return &specFreebsd.InterfaceGenerator{}
}

func metricsGenerators(conf *config.Config) []metrics.Generator {
	generators := []metrics.Generator{
		&metrics.LoadavgGenerator{},
		&metricsFreebsd.CPUUsageGenerator{},
		&metrics.FilesystemGenerator{IgnoreRegexp: conf.Filesystems.Ignore.Regexp, UseMountpoint: conf.Filesystems.UseMountpoint},
		&metricsFreebsd.MemoryGenerator{},
		&metrics.InterfaceGenerator{IgnoreRegexp: conf.Interfaces.Ignore.Regexp, Interval: metricsInterval},
	}

	return generators
}
