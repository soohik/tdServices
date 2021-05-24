package phone

type CreatePointRequestParams struct {
	Name string `json:"name" binding:"required"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Provider   string `json:"provider" binding:"required"`
	ProviderId string `json:"provider_id"`
}
