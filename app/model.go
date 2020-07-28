package app

type BingoMatrix struct {
	Fields [][]string `json:"fields"`
}

func TestData() [][]string {
	return [][]string{
		[]string{"test", "test", "test", "test"},
		[]string{"rest", "rest", "rest", "rest"},
		[]string{"rest", "rest", "rest", "rest"},
		[]string{"rest", "rest", "rest", "rest"},
	}
}

func ApiData() [][]string {
	return [][]string{
		[]string{"api", "api", "api", "api"},
		[]string{"api", "api", "api", "api"},
		[]string{"api", "api", "api", "api"},
		[]string{"api", "api", "api", "api"},
	}
}
