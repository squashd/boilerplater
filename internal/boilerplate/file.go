package boilerplate

import (
	"fmt"
	"os"
	"strings"

	"github.com/SQUASHD/boilerplater/shared/models"
)

func GenerateFunctionBoilerplate(boilerplate []models.FunctionBoilerplate) {
	for _, boilerplate := range boilerplate {
		generator := GetFunctionGenerator(boilerplate.Language)
		if generator == nil {
			fmt.Println("Unsupported language")
			return
		}

		var contentBuilder strings.Builder
		for _, functionName := range boilerplate.Functions {
			contentBuilder.WriteString(generator.GenerateFunction(functionName))
		}

		if err := os.WriteFile(boilerplate.FilePath, []byte(contentBuilder.String()), 0644); err != nil {
			fmt.Println(err)
		}
	}
}
