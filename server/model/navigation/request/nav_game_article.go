package request

// NavGameArticleRequest 游戏文章请求
type NavGameArticleRequest struct {
	GameId uint `json:"gameId" form:"gameId" binding:"required"` // 游戏ID
}
