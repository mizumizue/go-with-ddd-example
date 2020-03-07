package response

type UserResponse struct {
	UserID   string `json:"id"`
	UserName string `json:"name"`
}

func (res *UserResponse) SetUserID(userID string) {
	res.UserID = userID
}

func (res *UserResponse) SetUserName(userName string) {
	res.UserName = userName
}
