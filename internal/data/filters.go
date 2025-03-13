package data

import (
	"strings"

	"github.com/somashekhar/greenlight/internal/validator"
)

type Filters struct{
	Page int
	PageSize int
	Sort string
	SortSafelist []string
}

func (f Filters) sortColumn()string{
	for _, safeValue := range f.SortSafelist{
		if f.Sort == safeValue{
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: "+ f.Sort)
}

func(f Filters)sortDirection() string{
	if strings.HasPrefix(f.Sort, "-"){
		return "DESC"
	}
	return "ASC"
}

func ValidateFilters(v *validator.Validator, f Filters){
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maxium of 100")
	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort","invalid sort value")
}

// curl "localhost:4000/v1/movies?page=-1&pagesize=-1&sort=foo"
// curl "localhost:4000/v1/movies?page=abc&page_size=abc"
// curl "localhost:4000/v1/movies?title=black+panther"
// curl "localhost:4000/v1/movies?genres=adventure"
// curl "localhost:4000/v1/movies?sort=-title"
// curl "localhost:4000/v1/movies?sort=-runtime"