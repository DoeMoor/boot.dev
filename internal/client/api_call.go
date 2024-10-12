package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DoeMoor/pokedexcli/internal/pokecache"
)

// accept (url string) and (scheme *T) generic pointer to Json_scheme
//
// Example:
//
//	var loc scheme.Locations
//	err = client.ApiCall(url,&loc)
func ApiCall[T any](url string, scheme *T) error {
	cached, ok := pokecache.GetCache().Read(url)
	if ok {
		err := json.Unmarshal(cached, scheme)
		fmt.Println("Cache hit")
		if err != nil {
			return err
		}
		return nil
	}

	client := http.Client{}
	client.Timeout = 5 * time.Second

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Not Found")
		return nil
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("apicall status code: %v", resp.Status)
	}


	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	pokecache.GetCache().Write(url, buffer)

	err = json.Unmarshal(buffer, scheme)
	if err != nil {
		return err
	}

	return nil
}
