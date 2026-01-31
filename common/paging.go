package common

type Paging struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p *Paging) Preset() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 2 {
		p.Limit = 2
	}
	if p.Limit > 50 {
		p.Limit = 50
	}
}
