package domain

type Records struct {
	ID          int    `json:"id"`
	Title       string `json:"title,max=55" example:"Title"`
	Description string `json:"description,max=255" example:"Description"`
}

type UpdateRecord struct {
	Title       *string `json:"title" example:"Title"`
	Description *string `json:"description" example:"Description!"`
}

func (ur UpdateRecord) IsValid() bool {
	return ur.Title != nil && ur.Description != nil
}
