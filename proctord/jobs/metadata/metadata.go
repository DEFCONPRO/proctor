package metadata

import "proctor/proctord/jobs/metadata/env"

type Metadata struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	ImageName        string   `json:"image_name"`
	EnvVars          env.Vars `json:"env_vars"`
	AuthorizedGroups []string `json:"authorized_groups"`
	Author           string   `json:"author"`
	Contributors     string   `json:"contributors"`
	Organization     string   `json:"organization"`
}
