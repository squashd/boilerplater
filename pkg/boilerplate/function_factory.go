package boilerplate

func GetFunctionGenerator(language string) FunctionGenerator {
	switch language {
	case "JavaScript":
		return JavaScriptFunctionGenerator{}
	case "TypeScript":
		return TypeScriptFunctionGenerator{}
	case "Go":
		return GoFunctionGenerator{}
	case "Python":
		return PythonFunctionGenerator{}
	default:
		return nil
	}
}
