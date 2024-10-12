package endpoint_scheme

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int64                 `json:"game_index"`
	ID                   int64                 `json:"id"`
	Location             Location              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod Location                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`
	Version Location `json:"version"`
}

type Name struct {
	Language Location `json:"language"`
	Name     string   `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Location                        `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`
	Version          Location          `json:"version"`
}

type EncounterDetail struct {
	Chance          int64         `json:"chance"`
	ConditionValues []interface{} `json:"condition_values"`
	MaxLevel        int64         `json:"max_level"`
	Method          Location      `json:"method"`
	MinLevel        int64         `json:"min_level"`
}
