package functions

import (
	"strings"

	ironClient "github.com/mutemaniac/functions_go"
	"github.com/mutemaniac/steel/config"

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
	// err := docker.Build(route.Code, route.Runtime, route.Image, route.AppName)
	// if err != nil {
	// 	// TODO ceate Route failure & callback.
	// 	return route, err
	// }

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
	fmt.Printf("AppsAppRoutesRouteGet: %v \n", routewrapper.Route)
	//FIXME If there is no route.
	if routewrapper.Route.Path == "" {
		fmt.Println("Create a new route.")
		var routewrap ironClient.RouteWrapper
		routewrap.Route = route.Route
		retRoute, _, err := routeClient.AppsAppRoutesPost(route.AppName, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
		route.Route = retRoute.Route
	} else {
		fmt.Println("Update the old route.")
		routewrap := prepareUpdateModel(route.Route, routewrapper.Route)
		//route.SetDefaults(&routewrapper.Route)
		//route.Route.Path = ""
		routewrap.Route = route.Route
		routewrap.Route.Path = ""
		fmt.Printf("AppsAppRoutesRoutePatch: %v \n", routewrap)
		retRoute, resp, err := routeClient.AppsAppRoutesRoutePatch(route.AppName, route.Path, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
		fmt.Printf("AppsAppRoutesRoutePatch retRoute = %v \n resp=: %v \n", retRoute, resp.Message)
		route.Route = retRoute.Route
	}

	return route, nil
}

func prepareUpdateModel(newRoute ironClient.Route, oldR ironClient.Route) ironClient.RouteWrapper {
	var routewrap ironClient.RouteWrapper
	if newRoute.Path == "" {
		newRoute.Path = oldR.Path
	}
	if newRoute.Image == "" {
		newRoute.Image = oldR.Image
	}

	if newRoute.Memory == 0 {
		newRoute.Memory = oldR.Memory
	}

	if newRoute.Type_ == models.TypeNone {
		newRoute.Type_ = oldR.Type_
	}

	if newRoute.Format == "" {
		newRoute.Format = oldR.Format
	}

	if newRoute.MaxConcurrency == 0 {
		newRoute.MaxConcurrency = oldR.MaxConcurrency
	}

	if newRoute.Headers == nil {
		newRoute.Headers = oldR.Headers
	}

	if newRoute.Config == nil {
		newRoute.Config = oldR.Config
	}

	if newRoute.Timeout == 0 {
		newRoute.Timeout = oldR.Timeout
	}
	routewrap.Route = newRoute
	return routewrap
}
