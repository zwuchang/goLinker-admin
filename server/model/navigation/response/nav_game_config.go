package response

import "goLinker-admin/server/model/navigation"

type NavGameConfigResponse struct {
	GameConfig navigation.NavGameConfig `json:"gameConfig"`
}
