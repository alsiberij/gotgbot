package tg

type (
	Photo struct {
		Id       string `json:"file_id"`
		UniqueId string `json:"file_unique_id"`
		Size     int64  `json:"size"`
		Width    int64  `json:"width"`
		Height   int64  `json:"height"`
	}
)
