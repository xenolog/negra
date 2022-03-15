module github.com/xenolog/negra

go 1.17

require (
	// github.com/juju/juju v0.0.0-20211104194817-89faeee95080
	github.com/urfave/cli v0.0.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/klog v1.0.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
)

replace (
	github.com/urfave/cli => github.com/urfave/cli/v2 v2.4.0
	k8s.io/klog => k8s.io/klog v1.0.0
)
