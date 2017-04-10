package config

import (
	"flag"
)

var DockerHubServer = flag.String("dockerhub-server", "dcatalog.hnaresearch.com", "dockerhub server")
var DockerHubUser = flag.String("dockerhub-user", "qian.tang@hnair.com", "dockerhub user")
var DockerHubPwd = flag.String("dockerhub-pwd", "", "dockerhub password")
var DockerImageLib = flag.String("dockerhub-lib", "tangqian", "dockerhub image lib")

const (
	DockerImagePrefix = "serverless_"
)
