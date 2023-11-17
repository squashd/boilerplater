// generator functions for the multiple choice component
package input

// SkillLevelOptions returns choices for skill levels.
func SkillLevelOptions() []Choice {
	return []Choice{
		{"Beginner", "I'm comfortable with the basics", 0},
		{"Intermediate", "I can develop functional code independently", 1},
		{"Experienced", "I'm adept in advanced coding", 2},
	}
}

// TargetOutcomeOptions returns choices for the target outcome.
func TargetOutcomeOptions() []Choice {
	return []Choice{
		{"Learning", "I want this to be a learning experience", 0},
		{"Portfolio", "I want to build a project for my portfolio", 1},
		{"Produciton", "I want to build a production-ready project", 2},
	}
}

// LanguageSkillLevelOptions returns choices for language skill levels.
func LanguageSkillLevelOptions() []Choice {
	return []Choice{
		{"Beginner", "I'm new to this language", 0},
		{"Intermediate", "I'm comfortable with this language", 1},
		{"Experienced", "I'm highly proficient in this language", 2},
	}
}

// FilterChoices filters choices based on the an int value.
// Because the choices are ordered, we can filter them by the overall skill level.
// A beginner shouldn't worry about production-readiness or be able to select
// 'experienced' in a programming language.
func FilterChoices(overallSkillLevel int, choices []Choice) []Choice {
	filtered := []Choice{}

	for _, choice := range choices {
		if overallSkillLevel >= choice.Value {
			filtered = append(filtered, choice)
		}
	}

	return filtered
}
