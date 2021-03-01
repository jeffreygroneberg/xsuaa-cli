package domain

import "time"

type ListApp struct {
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
		GUID      string    `json:"guid"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		State     string    `json:"state"`
		Lifecycle struct {
			Type string `json:"type"`
			Data struct {
				Buildpacks []interface{} `json:"buildpacks"`
				Stack      string        `json:"stack"`
			} `json:"data"`
		} `json:"lifecycle"`
		Relationships struct {
			Space struct {
				Data struct {
					GUID string `json:"guid"`
				} `json:"data"`
			} `json:"space"`
		} `json:"relationships"`
		Metadata struct {
			Labels struct {
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
		} `json:"metadata"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			EnvironmentVariables struct {
				Href string `json:"href"`
			} `json:"environment_variables"`
			Space struct {
				Href string `json:"href"`
			} `json:"space"`
			Processes struct {
				Href string `json:"href"`
			} `json:"processes"`
			Packages struct {
				Href string `json:"href"`
			} `json:"packages"`
			CurrentDroplet struct {
				Href string `json:"href"`
			} `json:"current_droplet"`
			Droplets struct {
				Href string `json:"href"`
			} `json:"droplets"`
			Tasks struct {
				Href string `json:"href"`
			} `json:"tasks"`
			Start struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"start"`
			Stop struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"stop"`
			Revisions struct {
				Href string `json:"href"`
			} `json:"revisions"`
			DeployedRevisions struct {
				Href string `json:"href"`
			} `json:"deployed_revisions"`
			Features struct {
				Href string `json:"href"`
			} `json:"features"`
		} `json:"links"`
	} `json:"resources"`
	Included struct {
		Spaces []struct {
			GUID          string    `json:"guid"`
			CreatedAt     time.Time `json:"created_at"`
			UpdatedAt     time.Time `json:"updated_at"`
			Name          string    `json:"name"`
			Relationships struct {
				Organization struct {
					Data struct {
						GUID string `json:"guid"`
					} `json:"data"`
				} `json:"organization"`
				Quota struct {
					Data struct {
						GUID string `json:"guid"`
					} `json:"data"`
				} `json:"quota"`
			} `json:"relationships"`
			Metadata struct {
				Labels struct {
				} `json:"labels"`
				Annotations struct {
				} `json:"annotations"`
			} `json:"metadata"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Organization struct {
					Href string `json:"href"`
				} `json:"organization"`
				Features struct {
					Href string `json:"href"`
				} `json:"features"`
				ApplyManifest struct {
					Href   string `json:"href"`
					Method string `json:"method"`
				} `json:"apply_manifest"`
				Quota struct {
					Href string `json:"href"`
				} `json:"quota"`
			} `json:"links"`
		} `json:"spaces"`
		Organizations []struct {
			GUID          string    `json:"guid"`
			CreatedAt     time.Time `json:"created_at"`
			UpdatedAt     time.Time `json:"updated_at"`
			Name          string    `json:"name"`
			Suspended     bool      `json:"suspended"`
			Relationships struct {
				Quota struct {
					Data struct {
						GUID string `json:"guid"`
					} `json:"data"`
				} `json:"quota"`
			} `json:"relationships"`
			Metadata struct {
				Labels struct {
				} `json:"labels"`
				Annotations struct {
				} `json:"annotations"`
			} `json:"metadata"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Domains struct {
					Href string `json:"href"`
				} `json:"domains"`
				DefaultDomain struct {
					Href string `json:"href"`
				} `json:"default_domain"`
				Quota struct {
					Href string `json:"href"`
				} `json:"quota"`
			} `json:"links"`
		} `json:"organizations"`
	} `json:"included"`
}
