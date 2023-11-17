package models

// I have no clue whether this is the right way to do enums in Golang...
type ProficiencyLevel int

const (
	Beginner ProficiencyLevel = iota
	Intermediate
	Experienced
)

func (pl ProficiencyLevel) String() string {
	return [...]string{"Beginner", "Intermediate", "Experienced"}[pl]
}

type DesiredOutcome int

const (
	Learning DesiredOutcome = iota
	Portfolio
	Production
)

func (do DesiredOutcome) String() string {
	return [...]string{"Learning", "Portfolio", "Production"}[do]
}

type ProjectRequest struct {
	ProjectName     string           `json:"project_name"`
	Description     string           `json:"description"`
	Language        string           `json:"language"`
	Experience      ProficiencyLevel `json:"proficiency_level"`
	LangProficiency ProficiencyLevel `json:"language_proficiency"`
	DesiredOutcome  DesiredOutcome   `json:"desired_outcome"`
	MakeDirs        bool             `json:"make_dirs"`
	MakeBoilerplate bool             `json:"make_boilerplate"`
}
