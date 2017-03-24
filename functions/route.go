package functions

import (
	"strings"

	ironClient "github.com/iron-io/functions_go"
	"github.com/mutemaniac/steel/config"
	"github.com/mutemaniac/steel/docker"

	"fmt"

	"github.com/mutemaniac/steel/models"
)

// CreateRoute Create iron functions route wirh source code.
// @param route is the parameters that passed by http interface.
// @return
func CreateRoute(route models.ExRouteWrapper) (models.ExRouteWrapper, error) {
	fmt.Printf("Enter CreateRoute: %v.\n", route)
	if route.Image == "" {
		route.Image = config.DockerHubServer + `/` +
			config.DockerImageLib + `/` +
			config.DockerImagePrefix +
			route.AppName + "_" +
			strings.TrimPrefix(route.Path, `/`)
	}
	// Build image & push from code.
	err := docker.Build(route.Code, route.Runtime, route.Image, route.AppName)
	if err != nil {
		// TODO ceate Route failure & callback.
		return route, err
	}

	// Create Functions's route
	appClinet := ironClient.NewAppsApiWithBasePath(config.IronFunciotnsServer)
	// FIXME maybe need apiClinet.Configuration
	appwrapper, _, err := appClinet.AppsAppGet(route.AppName)
	if nil != err {
		// TODO ceate Route failure & callback.
		return route, err
	}
	if appwrapper.App.Name == "" {
		//There is no app then create one.
		var appwrap ironClient.AppWrapper
		appwrap.App.Name = route.AppName
		appClinet.AppsPost(appwrap)
	}

	routeClient := ironClient.NewRoutesApiWithBasePath(config.IronFunciotnsServer)
	routewrapper, _, err := routeClient.AppsAppRoutesRouteGet(route.AppName, route.Path)
	if nil != err {
		// TODO ceate Route failure & callback.
		return route, err
	}
	//FIXME If there is no route.
	if routewrapper.Route.Path == "" {
		var routewrap ironClient.RouteWrapper
		routewrap.Route = route.Route
		_, _, err := routeClient.AppsAppRoutesPost(route.AppName, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
	} else {
		var routewrap ironClient.RouteWrapper
		routewrap.Route = route.Route
		_, _, err := routeClient.AppsAppRoutesRoutePatch(route.AppName, route.Path, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
	}
	return route, nil
}
