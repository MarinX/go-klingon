package stapi

// CharacterService is an interface for interfacing with the character endpoints
// of the STAPI.
type CharacterService interface {
	Search(interface{}) (*CharactersResource, error)
	Get(interface{}) (*Character, error)
}

// CharacterServiceOp handles communication with the character related methods of
// the STAPI.
type CharacterServiceOp struct {
	client *Client
}

// Character represents a STAPI character
type Character struct {
	UID                    string             `json:"uid"`
	Name                   string             `json:"name"`
	Gender                 string             `json:"gender"`
	YearOfBirth            uint64             `json:"yearOfBirth"`
	MonthOfBirth           uint64             `json:"monthOfBirth"`
	DayOfBirth             uint64             `json:"dayOfBirth"`
	PlaceOfBirth           string             `json:"placeOfBirth"`
	YearOfDeath            uint64             `json:"yearOfDeath"`
	MonthOfDeath           uint64             `json:"monthOfDeath"`
	DayOfDeath             uint64             `json:"dayOfDeath"`
	PlaceOfDeath           string             `json:"placeOfDeath"`
	Height                 uint64             `json:"height"`
	Weight                 uint64             `json:"weight"`
	Deceased               bool               `json:"deceased"`
	BloodType              string             `json:"bloodType"`
	MaritalStatus          string             `json:"maritalStatus"`
	SerialNumber           string             `json:"serialNumber"`
	HologramActivationDate string             `json:"hologramActivationDate"`
	HologramStatus         string             `json:"hologramStatus"`
	HologramDateStatus     string             `json:"hologramDateStatus"`
	Hologram               bool               `json:"hologram"`
	FictionalCharacter     bool               `json:"fictionalCharacter"`
	Mirror                 bool               `json:"mirror"`
	AlternateReality       bool               `json:"alternateReality"`
	CharacterSpecies       []CharacterSpecies `json:"characterSpecies,omitempty"`
}

// CharacterSpecies represents a specie of character
type CharacterSpecies struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Numerator   int    `json:"numerator"`
	Denominator int    `json:"denominator"`
}

// CharactersResource is the result from /search endpoint
type CharactersResource struct {
	Page       *Page       `json:"page"`
	Sort       *Sort       `json:"sort"`
	Characters []Character `json:"characters"`
}

// CharacterResource represents the results from character endpoint
type CharacterResource struct {
	Character *Character `json:"character"`
}

// Search the character
func (c *CharacterServiceOp) Search(options interface{}) (*CharactersResource, error) {
	resource := new(CharactersResource)
	err := c.client.Post(BuildURL("character/search"), options, resource)
	return resource, err
}

// Get the character
func (c *CharacterServiceOp) Get(options interface{}) (*Character, error) {
	resource := new(CharacterResource)
	err := c.client.Get(BuildURL("character"), options, resource)
	return resource.Character, err
}
