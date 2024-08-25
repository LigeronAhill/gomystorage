package models

type EntityResponse[T Entity] struct {
	Context Context `json:"context,omitempty"`
	Meta    Meta    `json:"meta"`
	Rows    []T     `json:"rows"`
}

type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}
