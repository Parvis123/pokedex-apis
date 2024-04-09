package pokemon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
)

func TestGetPokemon(t *testing.T) {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    mockResponse := `{"abilities":[{"ability":{"name":"static","url":"https://pokeapi.co/api/v2/ability/9/"},"is_hidden":false,"slot":1},{"ability":{"name":"lightning-rod","url":"https://pokeapi.co/api/v2/ability/31/"},"is_hidden":true,"slot":3}],"name":"pikachu","sprites":{"other":{"home":{"front_default":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/home/25.png","front_shiny":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/home/shiny/25.png"}}},"height":4,"weight":60,"types":[{"slot":1,"type":{"name":"electric"}}],"cries":{"latest":"https://raw.githubusercontent.com/PokeAPI/cries/main/cries/pokemon/latest/25.ogg"}}`
    httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/pikachu",
        httpmock.NewStringResponder(200, mockResponse))

    req, err := http.NewRequest("GET", "/pokemon/pikachu", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router := mux.NewRouter()
    router.HandleFunc("/pokemon/{name}", GetPokemon)
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var expected, actual map[string]interface{}
    json.Unmarshal([]byte(mockResponse), &expected)
    json.Unmarshal(rr.Body.Bytes(), &actual)
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
}