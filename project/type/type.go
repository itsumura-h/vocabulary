package Type

type WordRow struct {
	Key     int
	Word    string
	Pron    string
	Meaning string
}

type Request struct {
	Words []string `json:"words"`
}
