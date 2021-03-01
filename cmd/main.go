package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
	"xsuaa-cli/pkg/domain"

	"github.com/cloudfoundry-community/go-cfclient"
)

var verbose bool
var interactiveMode bool

func main() {

	//arch := runtime.GOOS

	var username string
	var password string
	var appName string
	var appGuid string
	var api string
	var help bool

	flag.StringVar(&username, "username", "", "User to login at api endpoint")
	flag.StringVar(&password, "password", "", "Password to login at api endpoint")
	flag.StringVar(&appName, "appname", "", "Name of your app. Must be unique. Either appguid or appname must be set.")
	flag.StringVar(&appGuid, "appguid", "", "App Guid - Setting this will ignore the app name parameter. Either appguid or appname must be set.")
	flag.StringVar(&api, "api", "", "API endpoint of CF controller")
	flag.BoolVar(&verbose, "verbose", false, "Verbose mode")
	flag.BoolVar(&interactiveMode, "i", false, "Interactive mode - When multiple apps are found for a given name you are able to choose")
	flag.BoolVar(&help, "help", false, "Prints out the list of arguments and their description")
	flag.BoolVar(&help, "h", false, "Prints out the list of arguments and their description")

	flag.Parse()

	if help {
		flag.PrintDefaults()
	}

	if (appGuid == "" || appName == "") && (username == "" || password == "" || api == "") {
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	// clear screen to circumvent credentials on terminal window when using during presentations
	if !verbose {
		// freebsd, openbsd,
		// plan9, windows...
		var clearCommand string
		switch os := runtime.GOOS; os {
		case "windows":
			clearCommand = "cls"
		case "darwin":
			clearCommand = "clear"
		}

		command := exec.Command(clearCommand)
		command.Stdout = os.Stdout
		command.Run()
	}

	c := &cfclient.Config{
		ApiAddress: api,
		Username:   username,
		Password:   password,
	}
	client, _ := cfclient.NewClient(c)

	// try to get guid by name
	if appGuid == "" {
		guid, err := getAppGuidByName(client, appName)
		if err != nil {
			log.Fatalln(err)
			return
		}
		appGuid = guid
	}

	// try to get app env by guid
	appEnv, err := getAppEnvByGuid(client, appGuid)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// get XSUAA Credentials
	xsuaaCredentials, err := getXSUAAFromAppEnv(*appEnv)
	if err != nil {

		log.Fatalln(err)
		return

	}

	xsuaaUrl := xsuaaCredentials.URL
	xsuaaClientId := xsuaaCredentials.Clientid
	xsuaaClientIdEscaped := fmt.Sprint(url.QueryEscape(xsuaaClientId))
	xsuaaClientSecret := xsuaaCredentials.Clientsecret

	authorizationCodeURL :=
		fmt.Sprintf("%s/oauth/authorize?response_type=code&redirect_uri=http%%3A%%2F%%2Flocalhost:8080&client_id=%s", xsuaaUrl, xsuaaClientIdEscaped)

	if verbose {
		fmt.Print("Generated URL to get authorization code: " + authorizationCodeURL)
	}
	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}

	var authorizationCode string

	wg := sync.WaitGroup{}
	wg.Add(1)

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// shutdown server when callback's been successful and call waitinggroup
		defer func() {
			wg.Done()
			s.Shutdown(context.Background())
		}()

		p := "." + r.URL.Path
		if p == "./" {
			p = "./static/callback.html"
		}
		http.ServeFile(w, r, p)

		authorizationCode = r.RequestURI[len(r.RequestURI)-10:]
		accesstoken, err2 := getAccesstoken(xsuaaUrl, authorizationCode, xsuaaClientId, xsuaaClientSecret)

		if err2 != nil {
			log.Fatalln(err2)
		}

		if verbose {
			fmt.Println("Received token")
			fmt.Println("-------------------------------")
		}
		fmt.Println("bearer " + accesstoken.AccessToken)

	})

	// start servier in another routine as listenandserve is blocking
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// go to sleep for 1 second to allow http server to start
	time.Sleep(1 * time.Second)

	var openCommand string
	switch os := runtime.GOOS; os {
	case "darwin":
		openCommand = "open"
	case "windows":
		openCommand = "explorer"
	}

	// use default command to open browser with url
	command := exec.Command(openCommand, authorizationCodeURL)
	err2 := command.Run()

	if err2 != nil {
		fmt.Println(err2)
	}

	wg.Wait()
	os.Exit(0)
}

func getAccesstoken(xsuaaUrl, authorizationCode, xsuaaClientId, xsuaaClientSecret string) (*domain.XSUAAToken, error) {

	tokenUrl :=
		fmt.Sprintf("%v/oauth/token?grant_type=authorization_code&code=%v&redirect_uri=http%%3A%%2F%%2Flocalhost:8080", xsuaaUrl, authorizationCode)

	if verbose {
		fmt.Println("\nGenerated URL to get access token: " + tokenUrl)
	}

	hClient := http.Client{}
	tokenReq, _ := http.NewRequest("GET", tokenUrl, nil)
	tokenReq.SetBasicAuth(xsuaaClientId, xsuaaClientSecret)
	res, err := hClient.Do(tokenReq)

	if err != nil {
		return nil, errors.New("Can't access token endpoint")
	}

	jwtData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Can't read body of payload: %s", jwtData))
	}
	res.Body.Close()

	var result domain.XSUAAToken
	// Unmarshal or Decode the JSON to the interface.
	err = json.Unmarshal(jwtData, &result)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Can't parse json of payload: %s", jwtData))
	}

	return &result, nil

}

func getXSUAAFromAppEnv(env domain.AppEnv) (*domain.XSUAACredentials, error) {

	if env.SystemEnvJSON.VCAPSERVICES.Xsuaa == nil || len(env.SystemEnvJSON.VCAPSERVICES.Xsuaa) == 0 {
		return nil, errors.New("Can not found xsuaa as a service binding in your app")
	}

	xsuaa := env.SystemEnvJSON.VCAPSERVICES.Xsuaa[0].Credentials

	return &domain.XSUAACredentials{
		Tenantmode:      xsuaa.Tenantmode,
		Sburl:           xsuaa.Sburl,
		Clientid:        xsuaa.Clientid,
		Xsappname:       xsuaa.Xsappname,
		Clientsecret:    xsuaa.Clientsecret,
		URL:             xsuaa.URL,
		Uaadomain:       xsuaa.Uaadomain,
		Verificationkey: xsuaa.Verificationkey,
		Apiurl:          xsuaa.Apiurl,
		Identityzone:    xsuaa.Identityzone,
		Identityzoneid:  xsuaa.Identityzoneid,
		Tenantid:        xsuaa.Tenantid,
	}, nil

}

func getAppEnvByGuid(client *cfclient.Client, guid string) (*domain.AppEnv, error) {

	appEnvUrl := fmt.Sprintf("%v/v3/apps/%s/env", client.Config.ApiAddress, guid)

	req, _ := http.NewRequest("GET", appEnvUrl, nil)

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New(err.(cfclient.CloudFoundryError).Description)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var result domain.AppEnv
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(data, &result)

	if result.ApplicationEnvJSON.VCAPAPPLICATION.ApplicationName == "" {
		return &result, errors.New("No app found for given guid")
	}

	return &result, nil

}

func getAppGuidByName(client *cfclient.Client, name string) (string, error) {

	listAppsUrl := fmt.Sprintf("%s/v3/apps?names=%s&include=space,space.organization", client.Config.ApiAddress, name)

	req, _ := http.NewRequest("GET", listAppsUrl, nil)
	res, _ := client.Do(req)

	appData, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// Unmarshal or Decode the JSON to the interface.
	var appResult domain.ListApp

	json.Unmarshal(appData, &appResult)

	if appResult.Resources == nil || len(appResult.Resources) == 0 {
		return "", errors.New("Could not find an app with given name")
	}

	if len(appResult.Resources) > 1 {

		if interactiveMode {
			input := showAppsAndGetInput(appResult)
			return appResult.Resources[input].GUID, nil

		}

		return "", errors.New("Too many results for given name")
	}

	resources := appResult.Resources
	return resources[0].GUID, nil

}

func showAppsAndGetInput(apps domain.ListApp) int {

	// include sucks. GIMME EXPAND IN AN API FFS!
	spaces := make(map[string]domain.SpaceAndOrg)
	orgs := make(map[string]string)

	// get spaces
	for _, v := range apps.Included.Spaces {

		spaces[v.GUID] = domain.SpaceAndOrg{
			Name:    v.Name,
			OrgGuid: v.Relationships.Organization.Data.GUID,
		}

	}

	// get orgs
	for _, v := range apps.Included.Organizations {

		orgs[v.GUID] = v.Name

	}

	fmt.Println("Multiple apps found for given name - please choose the app from the given list")

	for ok := true; ok; {

		printAppList(apps, spaces, orgs)

		input, err := readInput(len(apps.Resources))

		if err != nil {
			ok = true
			fmt.Println("Wrong input. Must be number and a selection from the list")

			continue
		}

		return input

	}

	return -1

}

func readInput(itemCount int) (int, error) {

	fmt.Print("Choose: ")

	var value int
	_, err2 := fmt.Scanf("%d", &value)

	if err2 != nil {
		return -1, err2
	}

	if value > itemCount || value < 1 {
		return -1, errors.New("Value must be in range")
	}

	return value - 1, nil

}

func printAppList(apps domain.ListApp, spaces map[string]domain.SpaceAndOrg, orgs map[string]string) {

	for i, v := range apps.Resources {

		fmt.Println(fmt.Sprintf("%v. Org: %s Space: %s GUID: %s", i+1, orgs[spaces[v.Relationships.Space.Data.GUID].OrgGuid], spaces[v.Relationships.Space.Data.GUID].Name, v.GUID))

	}

}
