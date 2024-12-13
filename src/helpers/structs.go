package structs

// Constituency represents the main structure
type Constituency struct {
	Name            string `json:"constituency"`
	TotalPopulation int    `json:"total_population"`
	Wards           []Ward `json:"wards"`
}

// Ward represents a ward within the constituency
type Ward struct {
	Name       string    `json:"name"`
	Population int       `json:"population"`
	Amenities  Amenities `json:"amenities"`
}

// Amenities represent categorized lists of amenities
type Amenities struct {
	Health          []Health         `json:"Health,omitempty"`
	Education       []Education      `json:"Education,omitempty"`
	Markets         []Market         `json:"Markets,omitempty"`
	Recreation      []Recreation     `json:"Recreation,omitempty"`
	Security        []Security       `json:"Security,omitempty"`
	Transportation  []Transportation `json:"Transportation,omitempty"`
	ShoppingCenters []ShoppingCenter `json:"ShoppingCenters,omitempty"`
	PlacesOfWorship []PlaceOfWorship `json:"PlacesOfWorship,omitempty"`
	WaterSupply     []WaterSupply    `json:"WaterSupply,omitempty"`
}

// Health represents a health facility
type Health struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Services []string `json:"services,omitempty"`
}

// Education represents an educational institution
type Education struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Level    string `json:"level"`
}

// Market represents a market
type Market struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Recreation represents a recreational site
type Recreation struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Security represents a security facility
type Security struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Transportation represents a transportation hub
type Transportation struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// ShoppingCenter represents a shopping center
type ShoppingCenter struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// PlaceOfWorship represents a place of worship
type PlaceOfWorship struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Denomination string `json:"denomination"`
}

// WaterSupply represents a water supply project
type WaterSupply struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Services []string `json:"services,omitempty"`
}
