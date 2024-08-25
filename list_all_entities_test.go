package gomystorage

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/LigeronAhill/gomystorage/models"
	"github.com/joho/godotenv"
)

func TestApiClient_ListAllEntities(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("MS_TOKEN")
	client := New(token)
	body, err := client.ListAllEntities(models.ProductFolder{}, 10, 0)
	if err != nil {
		t.Errorf("ApiClient.GetEntity() error = %v", err)
		return
	}
	var response models.EntityResponse[models.ProductFolder]
	if err = json.Unmarshal(body, &response); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	folders := response.Rows
	t.Log(folders)
	if len(folders) == 0 {
		t.Errorf("Got empty result")
	}
}
