package token

type Ext struct {
	UserId      int64   `json:"user_id"`
	Permissions []int64 `json:"permissions"`
	ProjectId   int     `json:"project_id"`
	Namespace   string  `json:"namespace"`
}
