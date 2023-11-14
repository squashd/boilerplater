package models

import (
	"encoding/json"
)

type ProficiencyLevel int

const (
	Beginner ProficiencyLevel = iota
	Intermediate
	Experienced
)

func (pl ProficiencyLevel) String() string {
	return [...]string{"Beginner", "Intermediate", "Experienced"}[pl]
}

func (pl ProficiencyLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(pl.String())
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

func (do DesiredOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(do.String())
}

type ProjectRequest struct {
	ProjectName         string           `json:"project_name"`
	Description         string           `json:"description"`
	Language            string           `json:"language"`
	ProficiencyLevel    ProficiencyLevel `json:"proficiency_level"`
	LanguageProficiency ProficiencyLevel `json:"language_proficiency"`
	DesiredOutcome      DesiredOutcome   `json:"desired_outcome"`
	MakeDirs            bool             `json:"make_dirs"`
	MakeBoilerplate     bool             `json:"make_boilerplate"`
}
