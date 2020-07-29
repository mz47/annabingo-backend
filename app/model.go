package app

type BingoMatrix struct {
	Fields [][]string `json:"fields"`
}

func TestData() [][]string {
	return [][]string{
		{"test", "test", "test", "test"},
		{"rest", "rest", "rest", "rest"},
		{"rest", "rest", "rest", "rest"},
		{"rest", "rest", "rest", "rest"},
	}
}

func ApiData() [4][4]string {
	return [4][4]string{
		{"1", "2", "3", "4"},
		{"5", "6", "7", "8"},
		{"9", "10", "11", "12"},
		{"13", "14", "15", "16"},
	}
}
