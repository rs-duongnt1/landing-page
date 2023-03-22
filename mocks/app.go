package mocks

type Skill struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
type Slide struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Year        int     `json:"year"`
	Title       string  `json:"title"`
	Skills      []Skill `json:"skills"`
}

type Work struct {
	Title     string  `json:"title"`
	Skills    []Skill `json:"skills"`
	Thumbnail string  `json:"thumbnail"`
}

type People struct {
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Avatar      string `json:"avatar"`
	PositionJob string `json:"position_job"`
}
