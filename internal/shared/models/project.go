// The various structs for the project outline.
package models

// Tips are more heavily favoured in beginner projects
type BeginnerProject struct {
	Title           string        `json:"title"`
	Objective       string        `json:"objective"`
	Steps           []ProjectStep `json:"steps"`
	WatchOuts       []string      `json:"watchOuts"`
	ExtraChallenges []string      `json:"extraChallenges"`
}

// IntermediateProject has a larger focus on features
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

// Tips should be optionable (this nilable)
type ProjectStep struct {
	Description string  `json:"description"`
	Tips        *string `json:"tips,omitempty"`
}

// ExperiencedProject focuses more on DetailedFeatures
type ExperiencedProject struct {
	Title              string             `json:"title"`
	Objective          string             `json:"objective"`
	DetailedFeatures   []DetailedFeature  `json:"detailedFeatures"`
	DevelopmentProcess DevelopmentProcess `json:"developmentProcess"`
	Challenges         []string           `json:"challenges"`
}

// DetailedFeature is broken down into clear implementation steps
type DetailedFeature struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	ImplementationSteps []string `json:"implementationSteps"`
}

// DevelopmentProcess is sometimes useful, sometimes not
// If you know what you need to do I don't really know why you might be using this
type DevelopmentProcess struct {
	Setup     string   `json:"setup"`
	Phases    []string `json:"phases"`
	Testing   string   `json:"testing"`
	Debugging string   `json:"debugging"`
}
