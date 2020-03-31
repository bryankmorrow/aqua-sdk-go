package assurance // import "github.com/BryanKMorrow/aqua-sdk-go/types/assurance"

type Scope struct {
	Expression string `json:"expression"`  // v1 is variable 1, v2 is variable 2 etc. AND=&& OR=||
	Variables  []ScopeVariable `json:"variables"`
}

type ScopeVariable struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}