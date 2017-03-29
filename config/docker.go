package config

import (
	"os"
)

const (
	dockerHubServerEvn = "DOCKER_HUB_SERVER_EVN"
	dockerHubUserEvn   = "DOCKER_HUB_USER_EVN"
	dockerHubPwdEvn    = "DOCKER_HUB_PWD_EVN"
)

var (
	DockerHubServer string
	DockerHubUser   string
	DockerHubPwd    string
)

const (
	DockerImageLib    = "library"
	DockerImagePrefix = "serverless_"
)

func init() {
	DockerHubServer = os.Getenv(dockerHubServerEvn)
	if DockerHubServer == "" {
		DockerHubServer = "dcatalog.hnaresearch.com"
	}
	DockerHubUser = os.Getenv(dockerHubUserEvn)
	if DockerHubUser == "" {
		DockerHubUser = "qian.tang@hnair.com"
	}
	DockerHubPwd = os.Getenv(dockerHubPwdEvn)
}
