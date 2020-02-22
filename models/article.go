package models

// Article 文章模型
type Article struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Remark     string `json:"remark"`
	VideoURL   string `json:"video_url"`
	VideoCover string `json:"video_cover"`
	Picture    string `json:"picture"`
	Type       string `json:"type"`
	Model
}
