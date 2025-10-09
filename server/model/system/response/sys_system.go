package response

import "goLinker-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
