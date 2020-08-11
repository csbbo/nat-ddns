package frp

type CreateFRPParm struct {
	LocalPort int `json:"localPort" binding:"required"`
}
