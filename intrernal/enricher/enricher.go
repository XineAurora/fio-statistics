package enricher

type Enricher interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetNationality(name string) (string, error)
}
