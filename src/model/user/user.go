package user

type User struct {
	Id              string `json:"id" column:"id"`
	Password        string `json:"password" column:"password"`
	Name            string `json:"name" column:"name"`
	Email           string `json:"email" column:"email"`
	HeadPhoto       string `json:"head_photo" column:"head_photo"`
	Gender          bool   `json:"gender" column:"gender"`
	PermissionLevel int    `json:"permission_level" column:"u_type"`
}
