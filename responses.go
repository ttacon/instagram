package instagram

type Meta struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}

type UserResponse struct {
	Meta Meta `json:"meta"`
	User User `json:"data"`
}

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string `json:"bio"`
	Website        string `json:"website"`
	Counts         Counts `json:"counts"`
}

type Counts struct {
	Media      int64 `json:"media"`
	Follows    int64 `json:"follows"`
	FollowedBy int64 `json:"followed_by"`
}
