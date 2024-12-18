package param

type ReqCreatePost struct {
	Title   string `json:"title" binding:"required,min=1,max=255"`
	Content string `json:"content" binding:"required,min=1"`
	Camera  string `json:"camera" binding:"max=255"`
	Lens    string `json:"lens" binding:"max=255"`

	Images []PostImage `json:"images" binding:"required"`
}

type PostImage struct {
	Url    string `json:"url" binding:"required,url"`
	Width  int    `json:"width" binding:"required"`
	Height int    `json:"height" binding:"required"`
	Index  int    `json:"index" binding:"required"`
}
