package param

type ReqRegister struct {
	Email    string `json:"email" binding:"required,max=255,email"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type ReqLogin struct {
	Email    string `json:"email" binding:"required,min=2,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type ReqUpdateProfile struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Avatar      string `json:"avatar" binding:"max=255"`
	Description string `json:"description" binding:"max=1024"`
	Background  string `json:"background" binding:"max=1024"`
}
