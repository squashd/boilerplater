package models

type BeginnerProject struct {
	Title           string        `json:"title"`
	Objective       string        `json:"objective"`
	Steps           []ProjectStep `json:"steps"`
	WatchOuts       []string      `json:"watchOuts"`
	ExtraChallenges []string      `json:"extraChallenges"`
}

type IntermediateProject struct {
	Title     string        `json:"title"`
	Objective string        `json:"objective"`
	Features  []Feature     `json:"features"`
	Steps     []ProjectStep `json:"steps"`
	Setup     string        `json:"setup"`
	Testing   string        `json:"testing"`
	Debugging string        `json:"debugging"`
	Extras    []string      `json:"extras"`
}

type Feature struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tips        []string `json:"tips"`
}

type ProjectStep struct {
	Description string  `json:"description"`
	Tips        *string `json:"tips,omitempty"`
}

type AdvancedProject struct {
	Title              string             `json:"title"`
	Objective          string             `json:"objective"`
	DetailedFeatures   []DetailedFeature  `json:"detailedFeatures"`
	DevelopmentProcess DevelopmentProcess `json:"developmentProcess"`
	Challenges         []string           `json:"challenges"`
}

type DetailedFeature struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	ImplementationSteps []string `json:"implementationSteps"`
}

type DevelopmentProcess struct {
	Setup     string   `json:"setup"`
	Phases    []string `json:"phases"`
	Testing   string   `json:"testing"`
	Debugging string   `json:"debugging"`
}
