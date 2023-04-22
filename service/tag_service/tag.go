package tag_service

import "github.com/monegim/product-api-go/models"

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
	PageNum    int
	PageSize   int
}

func (t *Tag)GetAll([]models.Tag, error)  {
	
}
