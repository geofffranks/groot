module code.cloudfoundry.org/groot

go 1.16

require (
	code.cloudfoundry.org/lager v2.0.0+incompatible
	github.com/containers/image/v5 v5.21.1
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.3-0.20211202193544-a5463b7f9c84
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.5
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/urfave/cli => github.com/urfave/cli v1.22.1
