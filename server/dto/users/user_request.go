package usersdto

type UpdateUserPicture struct {
	Email     string `json:"email" form:"email"`
	Fullname  string `json:"fullname" form:"fullname"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}
