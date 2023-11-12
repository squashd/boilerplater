package models

type ProjectStructure struct {
	RootDirectory string      `json:"rootDirectory,omitempty"`
	Directories   []Directory `json:"directories"`
	MainFiles     []string    `json:"mainFiles"`
}

type Directory struct {
	Name    string      `json:"name"`
	SubDirs []Directory `json:"subDirs,omitempty"`
	Files   []string    `json:"files"`
}

type FunctionBoilerplate struct {
	Language  string   `json:"language"`
	FilePath  string   `json:"filePath"`
	Functions []string `json:"functions"`
}
