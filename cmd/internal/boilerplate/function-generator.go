package boilerplate

import "fmt"

type FunctionGenerator interface {
	GenerateFunction(functionName string) string
}

type JavaScriptFunctionGenerator struct{}

func (j JavaScriptFunctionGenerator) GenerateFunction(functionName string) string {
	return fmt.Sprintf("function %s() {\n    // TODO: Implement\n}\n\n", functionName)
}

type GoFunctionGenerator struct{}

func (g GoFunctionGenerator) GenerateFunction(functionName string) string {
	return fmt.Sprintf("func %s() {\n    // TODO: Implement\n}\n\n", functionName)
}

type PythonFunctionGenerator struct{}

func (p PythonFunctionGenerator) GenerateFunction(functionName string) string {
	return fmt.Sprintf("def %s():\n    # TODO: Implement\n\n", functionName)
}

type TypeScriptFunctionGenerator struct{}

func (t TypeScriptFunctionGenerator) GenerateFunction(functionName string) string {
	return fmt.Sprintf("function %s() {\n    // TODO: Implement\n}\n\n", functionName)
}
