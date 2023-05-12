package main

type Car interface {
	Name() string
	Logo() string
}

type Motobike interface {
	Name() string
	Logo() string
}

type Vinfast struct {
}

type Merc struct {
}

func (Vinfast) Name(name string) string {
	name = "Vinfast"
	return name
}
func (Vinfast) Logo(logo string) string {
	logo = "Vinfast"
	return logo
}
