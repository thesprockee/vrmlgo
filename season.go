package vrmlgo

type Season struct {
	ID           string `json:"seasonID"`
	Name         string `json:"seasonName"`
	GameName     string `json:"gameName"`
	GameURLShort string `json:"gameUrlShort"`
	GameActive   bool   `json:"gameActive"`
	IsCurrent    bool   `json:"isCurrent"`
}
