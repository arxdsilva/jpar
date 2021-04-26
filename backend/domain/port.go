package domain

type Port struct {
	ID          string    `json:"id",pg:"id"`
	Name        string    `json:"name",pg:"name"`
	Coordinates []float32 `json:"coordinates",pg:"coordinates"`
	City        string    `json:"city",pg:"city"`
	Province    string    `json:"province",pg:"province"`
	Country     string    `json:"country",pg:"country"`
	Alias       []string  `json:"alias",pg:"alias"`
	Regions     []string  `json:"regions",pg:"regions"`
	Timezone    string    `json:"timezone",pg:"timezone"`
	Unlocs      []string  `json:"unlocs",pg:"unlocs"`
	Code        string    `json:"code",pg:"code"`
}

type PortService interface {
	UpsertPort(Port) error
	ListPorts() ([]Port, error)
}
