package runners

//GetRunner returns an object that implements the Runner interface for the
//the given language. If the language is not supported it returns a nil
func GetRunner(lang string) Runner {
	runner, found := supportedLanguages[lang]
	if found {
		return runner
	}
	return nil
}
