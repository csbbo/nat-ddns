package account


type RegistParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
