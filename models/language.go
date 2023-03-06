package models

type Language int

const (
	Cpp = iota
	Java
	Python
)

func (lang Language) GetLanguage() string {
	return []string{"cpp", "java", "python"}[lang]
}

func (lang Language) GetExtension() string {
	return []string{".cpp", ".java", ".py"}[lang]
}
