package functions

import (
	"context"
	"strings"
	"time"

	ironClient "github.com/mutemaniac/functions_go"
	"github.com/mutemaniac/steel/config"

	"fmt"

	"net/http"

	"encoding/json"

	"bytes"

	"github.com/mutemaniac/steel/models"
)

func doCallBack(ctx context.Context, taskid string, route models.ExRouteWrapper, retErr error, callback string) error {
	// Do call back
	fmt.Printf("Enter doCallBack. \n")
	br, err := json.Marshal(struct {
		Route  models.ExRouteWrapper `json:"route"`
		RetErr error                 `json:"error"`
		Taskid string                `json:"taskid"`
	}{
		Route:  route,
		RetErr: retErr,
		Taskid: taskid,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", callback, bytes.NewBuffer(br))
	if err != nil {
		return err
	}
	tr := &http.Transport{}
	client := http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() {
		c <- func(resp *http.Response, err error) error {
			if err == nil {
				defer resp.Body.Close()
			}
			return err
		}(client.Do(req))
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
func AsyncCreateRoute(ctx context.Context, taskid string, args interface{}) {
	//route models.AsyncRouteWrapper
	route, ok := args.(models.AsyncRouteWrapper)
	if !ok {
		// TODO
	}
	retRoute, err := CreateRoute(ctx, route.ExRouteWrapper)
	if err != nil {
		fmt.Printf("CreateRoute error, %v\n", err)
		//panic(err)
	}
	// Do call back
	err = doCallBack(ctx, taskid, retRoute, err, route.Callback)
	fmt.Printf("doCallBack error, %v .\n", err)
}

// CreateRoute Create iron functions route wirh source code.
// @param route is the parameters that passed by http interface.
// @return
func CreateRoute(ctx context.Context, route models.ExRouteWrapper) (models.ExRouteWrapper, error) {
	fmt.Printf("Enter CreateRoute.\n")
	if route.Image == "" {
		route.Image = config.DockerHubServer + `/` +
			config.DockerImageLib + `/` +
			config.DockerImagePrefix +
			route.AppName + "_" +
			strings.TrimPrefix(route.Path, `/`)
	}
	// Build image & push from code.
	// err := docker.Build(ctx, route.Code, route.Runtime, route.Image, route.AppName)
	// if err != nil {
	// 	// TODO ceate Route failure & callback.
	// 	return route, err
	// }
	time.Sleep(time.Second * 30)

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
		_, _, err := appClinet.AppsPost(appwrap)
		if nil != err {
			// TODO ceate Route failure & callback.
			return route, err
		}
		//fmt.Printf("New functions app: %v \n", retAppwrap)
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
		fmt.Printf("before AppsAppRoutesPost route: %v\n", route)
		fmt.Printf("before AppsAppRoutesPost routewrap: %v\n", routewrap)
		retRoute, _, err := routeClient.AppsAppRoutesPost(route.AppName, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
		fmt.Printf("AppsAppRoutesPost retRoute: %v\n", retRoute)
		route.Route = retRoute.Route
	} else {
		fmt.Println("Update the old route.")
		routewrap := prepareUpdateModel(route.Route, routewrapper.Route)
		//route.SetDefaults(&routewrapper.Route)
		//route.Route.Path = ""
		routewrap.Route = route.Route
		routewrap.Route.Path = ""
		fmt.Printf("AppsAppRoutesRoutePatch: %v \n", routewrap)
		retRoute, _, err := routeClient.AppsAppRoutesRoutePatch(route.AppName, route.Path, routewrap)
		if err != nil {
			// TODO ceate Route failure & callback.
			return route, err
		}
		//fmt.Printf("AppsAppRoutesRoutePatch retRoute = %v \n resp=: %v \n", retRoute, resp.Message)
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
