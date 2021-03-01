# Disclaimer

Wrote this in my free time as a **private person**. **This is not affiliated with SAP!**

# Purpose

The purpose of this tool is to allow a **oAuth user grant flow JWT retrieval** been automated within a CLI tool.
Most tools only implemented client credentials flow.

This CLI tool will boostrap a small http server in the background that will take care of the oAuth redirect after a successfull authentication at the IdP. 
**After the callback has been triggered it will print the JWT with the user principal on the console.**

# Usage
The MacOS binary is already been shipped in this repo, but feel free to build it once again if you don't trust or if you need a Windows binary :) 
```zsh
go build -o /usr/local/bin/xsuaa-cli cmd/main.go
```
# How to use it? 
```zsh
xsuaa-cli -username user -password "password" -api https://api.cf.sap.hana.ondemand.com -appname yourappname
```
# arguments

```zsh
  -api string
        API endpoint of CF controller
  -appguid string
        App Guid - Setting this will ignore the app name parameter. Either appguid or appname must be set.
  -appname string
        Name of your app. Must be unique. Either appguid or appname must be set.
  -h    Prints out the list of arguments and their description
  -help
        Prints out the list of arguments and their description
  -i    Interactive mode - When multiple apps are found for a given name you are able to choose
  -password string
        Password to login at api endpoint
  -username string
        User to login at api endpoint
  -verbose
        Verbose mode
```

Guess the most important argument is either **appguid** or **appname**.
If you set **appguid** it will ignore the appname argument and tries to get the token for the given appguid. 
**appname** is a little bit more flexible. By default the CLI tries to find an app for the given name, but refuses to retrieve a JWT if multiple apps are found.

## Interactive mode
You can use the **-i** argument to tell the CLI tool to go in interactive mode. If the CLI finds more than one suitable app for a given name it will show you a list to choose from.

An example for interactive mode might go like this:
```zsh
xsuaa-cli -username user -password "password" -api https://api.cf.sap.hana.ondemand.com -appname yourappname -i
```

As a result you might then see something like this, showing you some meta about the app to make your life easier

```zsh
Multiple apps found for given name - please choose the app from the given list
1. Org: SAP_IT_Cloud_sapitcft Space: onboarding GUID: 451f2afb-c6b6-4dc9-8506-7d8f332805a8
2. Org: SAP_IT_Cloud_sapitcf Space: onboarding GUID: 0243b047-c34b-4a33-a2de-439df9367036
```

If only ONE app is found for a given name, no choice needs to be made. 
