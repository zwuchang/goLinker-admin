package response

import "goLinker-admin/server/model/navigation"

type NavContactConfigResponse struct {
	ContactConfig navigation.NavContactConfig `json:"contactConfig"`
}

type NavContactMethodResponse struct {
	ContactMethod navigation.NavContactMethod `json:"contactMethod"`
}
