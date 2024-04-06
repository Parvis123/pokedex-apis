package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
    Name string `json:"name"`
}

func GetPokemon(name string) (*Pokemon, error) {
    resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var pokemon Pokemon
    err = json.Unmarshal(body, &pokemon)
    if err != nil {
        return nil, err
    }

    return &pokemon, nil
}