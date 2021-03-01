package domain

type AppEnv struct {
	EnvironmentVariables struct {
		GOPACKAGENAME        string `json:"GOPACKAGENAME"`
		GOVERSION            string `json:"GOVERSION"`
		GOINSTALLPACKAGESPEC string `json:"GO_INSTALL_PACKAGE_SPEC"`
	} `json:"environment_variables"`
	StagingEnvJSON struct {
		NpmConfigSapRegistry string `json:"npm_config_@sap:registry"`
	} `json:"staging_env_json"`
	RunningEnvJSON struct {
		CREDHUBAPI string `json:"CREDHUB_API"`
	} `json:"running_env_json"`
	SystemEnvJSON struct {
		VCAPSERVICES struct {
			Hana []struct {
				Label        string      `json:"label"`
				Provider     interface{} `json:"provider"`
				Plan         string      `json:"plan"`
				Name         string      `json:"name"`
				Tags         []string    `json:"tags"`
				InstanceName string      `json:"instance_name"`
				BindingName  interface{} `json:"binding_name"`
				Credentials  struct {
					Host        string `json:"host"`
					Port        string `json:"port"`
					Driver      string `json:"driver"`
					URL         string `json:"url"`
					Schema      string `json:"schema"`
					User        string `json:"user"`
					Password    string `json:"password"`
					Certificate string `json:"certificate"`
				} `json:"credentials"`
				SyslogDrainURL interface{}   `json:"syslog_drain_url"`
				VolumeMounts   []interface{} `json:"volume_mounts"`
			} `json:"hana"`
			Xsuaa []struct {
				Label        string      `json:"label"`
				Provider     interface{} `json:"provider"`
				Plan         string      `json:"plan"`
				Name         string      `json:"name"`
				Tags         []string    `json:"tags"`
				InstanceName string      `json:"instance_name"`
				BindingName  interface{} `json:"binding_name"`
				Credentials  struct {
					Tenantmode      string `json:"tenantmode"`
					Sburl           string `json:"sburl"`
					Clientid        string `json:"clientid"`
					Xsappname       string `json:"xsappname"`
					Clientsecret    string `json:"clientsecret"`
					URL             string `json:"url"`
					Uaadomain       string `json:"uaadomain"`
					Verificationkey string `json:"verificationkey"`
					Apiurl          string `json:"apiurl"`
					Identityzone    string `json:"identityzone"`
					Identityzoneid  string `json:"identityzoneid"`
					Tenantid        string `json:"tenantid"`
				} `json:"credentials"`
				SyslogDrainURL interface{}   `json:"syslog_drain_url"`
				VolumeMounts   []interface{} `json:"volume_mounts"`
			} `json:"xsuaa"`
		} `json:"VCAP_SERVICES"`
	} `json:"system_env_json"`
	ApplicationEnvJSON struct {
		VCAPAPPLICATION struct {
			CfAPI  string `json:"cf_api"`
			Limits struct {
				Fds int `json:"fds"`
			} `json:"limits"`
			ApplicationName  string      `json:"application_name"`
			ApplicationUris  []string    `json:"application_uris"`
			Name             string      `json:"name"`
			SpaceName        string      `json:"space_name"`
			SpaceID          string      `json:"space_id"`
			OrganizationID   string      `json:"organization_id"`
			OrganizationName string      `json:"organization_name"`
			Uris             []string    `json:"uris"`
			Users            interface{} `json:"users"`
			ApplicationID    string      `json:"application_id"`
		} `json:"VCAP_APPLICATION"`
	} `json:"application_env_json"`
}
