package stapi

// Page represents pagination for STAPI
type Page struct {
	PageNumber       uint64 `json:"pageNumber"`
	PageSize         uint64 `json:"pageSize"`
	NumberOfElements uint64 `json:"numberOfElements"`
	TotalElements    uint64 `json:"totalElements"`
	TotalPages       uint64 `json:"totalPages"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
}

// Sort represents sort options for STAPI
type Sort struct {
	Clauses []Clause `json:"clauses"`
}

// Clause model, child of Sort
type Clause struct {
	Name        string `json:"name"`
	Direction   string `json:"direction"`
	ClauseOrder uint64 `json:"clauseOrder"`
}
