package stapi

import "testing"

func TestCharacterSearch(t *testing.T) {
	setup()
	defer teardown()

	responseBody = `
	{
		"page": {
			"pageNumber": 0,
			"pageSize": 50,
			"numberOfElements": 3,
			"totalElements": 3,
			"totalPages": 1,
			"firstPage": true,
			"lastPage": true
		},
		"sort": {
			"clauses": []
		},
		"characters": [
        {
            "uid": "CHMA0000115364",
            "name": "Nyota Uhura",
            "gender": null,
            "yearOfBirth": null,
            "monthOfBirth": null,
            "hologram": false,
            "fictionalCharacter": false,
            "mirror": false,
            "alternateReality": true
		}
		]
	}
	`

	resp, err := stapiClient.Character.Search(struct {
		Name string `url:"name"`
	}{"Uhura"})
	<-requestChan
	if err != nil {
		t.Error(err)
		return
	}

	if len(resp.Characters) <= 0 {
		t.Error("invalid character length, expected 1, got 0")
	}

	if resp.Page.PageSize != 50 {
		t.Errorf("invalid page size, exptected 50 got %v\n", resp.Page.PageSize)
	}

	if resp.Page.NumberOfElements != 3 {
		t.Errorf("invalid num of elements, exptected 3 got %v\n", resp.Page.NumberOfElements)
	}

	if resp.Page.TotalElements != 3 {
		t.Errorf("invalid total elements, exptected 3 got %v\n", resp.Page.TotalElements)
	}

	if resp.Page.TotalPages != 1 {
		t.Errorf("invalid total pages, exptected 1 got %v\n", resp.Page.TotalPages)
	}

	if !resp.Page.FirstPage {
		t.Error("expected first page to be true, got false")
	}

	if !resp.Page.LastPage {
		t.Error("expected last page to be true, got false")
	}

}

func TestCharacterGet(t *testing.T) {
	setup()
	defer teardown()

	responseBody = `
	{
		"character": {
			"uid": "CHMA0000068639",
			"name": "Nyota Uhura",
			"gender": "F",
			"yearOfBirth": 2239,
			"characterSpecies": [
				{
					"uid": "SPMA0000026314",
					"name": "Human",
					"numerator": 1,
					"denominator": 1
				}
			]
		}
	}
	`

	ch, err := stapiClient.Character.Get(struct {
		UID string `url:"uid"`
	}{"CHMA0000068639"})
	<-requestChan
	if err != nil {
		t.Error(err)
		return
	}

	if ch == nil {
		t.Error("error, character is nil")
		return
	}

	if len(ch.CharacterSpecies) <= 0 {
		t.Error("invalid length of character species, exptected > 0, got 0")
		return
	}

	if ch.CharacterSpecies[0].Name != "Human" {
		t.Errorf("invalid species name, expected %s, got %s", "Human", ch.CharacterSpecies[0].Name)
	}
}
