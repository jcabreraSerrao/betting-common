package dto

type FiberCustom struct {
	Token       string
	Group       uint
	Subgroup    []uint
	IdUser      uint
	Permissions []string
	NameGroup   string
}
