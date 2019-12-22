package action

type Result struct {
	Status int
	Page   string
	Dto    interface{}
}

func NewResult(status int, page string, dto interface{}) *Result {
	return &Result{
		Status: status,
		Page:   page,
		Dto:    dto,
	}
}
