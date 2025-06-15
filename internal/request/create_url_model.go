package request

type CreateUrl struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func NewCreateUrl(source string, destination string) (string, string, CreateUrl) {
	return "https://lat.sh/api/generate", "application/json", CreateUrl{
		Source:      source,
		Destination: destination,
	}
}
