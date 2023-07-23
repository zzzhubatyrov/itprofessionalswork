package role

type HR struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserTag   string `json:"user_tag"`
}
