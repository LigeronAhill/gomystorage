package gomystorage

import (
	"encoding/json"
	"gomystorage/models"

	"os"
	"testing"
)

func TestApiClient_SearchEntity(t *testing.T) {

	token := os.Getenv("MS_TOKEN")
	client := New(token)

	search := "formation"
	resp, err := client.SearchEntity(models.ProductFolder{}, search)
	if err != nil {
		t.Errorf("ApiClient.SearchEntity() error = %v", err)
	}
	var response models.EntityResponse[models.ProductFolder]
	if err = json.Unmarshal(resp, &response); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	folders := response.Rows
	t.Log(folders)
	if len(folders) == 0 {
		t.Errorf("Got empty result")
	}
}
