package domain

import "time"

type GetServiceBinding struct {
	GUID string `json:"guid"`
	Type string `json:"type"`
	Data struct {
		Name         string      `json:"name"`
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
	} `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		ServiceInstance struct {
			Href string `json:"href"`
		} `json:"service_instance"`
		App struct {
			Href string `json:"href"`
		} `json:"app"`
	} `json:"links"`
}

type ListServiceBinding struct {
	Pagination struct {
		TotalResults int `json:"total_results"`
		TotalPages   int `json:"total_pages"`
		First        struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"pagination"`
	Resources []struct {
		GUID string `json:"guid"`
		Type string `json:"type"`
		Data struct {
			Name         string      `json:"name"`
			InstanceName string      `json:"instance_name"`
			BindingName  interface{} `json:"binding_name"`
			Credentials  struct {
				RedactedMessage string `json:"redacted_message"`
			} `json:"credentials"`
			SyslogDrainURL interface{}   `json:"syslog_drain_url"`
			VolumeMounts   []interface{} `json:"volume_mounts"`
		} `json:"data"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Links     struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			ServiceInstance struct {
				Href string `json:"href"`
			} `json:"service_instance"`
			App struct {
				Href string `json:"href"`
			} `json:"app"`
		} `json:"links"`
	} `json:"resources"`
}
