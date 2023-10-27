package Model

type User struct {
	Id          int    `json:"id" gorm:"primaryKey;column:id"`
	Username    string `json:"username" gorm:"column:username"`
	Password    string `json:"password" gorm:"column:password"`
	NickName    string `json:"nick_name" gorm:"column:nick_name"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time"`
}

func (u User) TableName() string {
	return "user"
}
