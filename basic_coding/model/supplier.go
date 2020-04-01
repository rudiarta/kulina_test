package model

type SupplierInterface interface {
	Insert() Supplier
}

type Supplier struct {
	Id      int
	Name    string
	Address string
}

func (s *Supplier) Insert() Supplier {
	return *s
}
