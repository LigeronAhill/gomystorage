package gomystorage

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/LigeronAhill/gomystorage/models"
)

func TestApiClient_SearchEntity(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
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
