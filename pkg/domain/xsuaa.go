package domain

type XSUAACredentials struct {
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
}
