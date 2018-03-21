package router

type Payload struct {
	Id             string      `json:"id"`
	Type           string      `json:"type"`
	Version        int         `json:"version"`
	OrganisationId string      `json:"organisation_id"`
	Attributes     interface{} `json:"attributes"`
}
