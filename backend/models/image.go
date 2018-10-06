package models

// Image ...
type Image struct {
	UUID      string `json:"UUID"`
	Name      string `json:"Name"`
	Materials string `json:"Materials"`
	Year      string `json:"Year"`
	Size      string `json:"Size"`
	Type      string `json:"Type"`
	Ext       string `json:"Ext"`
	IsForSale bool   `json:"IsForSale"`
}
