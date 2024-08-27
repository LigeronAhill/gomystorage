package gomystorage

import (
	"encoding/json"
	"gomystorage/models"
	"os"
	"testing"
)

func TestApiClient_ListAllEntities(t *testing.T) {

	token := os.Getenv("MS_TOKEN")
	client := New(token)
	body, err := client.ListAllEntities(models.Variant{}, 1000, 0)
	if err != nil {
		t.Errorf("ApiClient.GetEntity() error = %v", err)
		return
	}
	var response models.EntityResponse[models.Variant]
	if err = json.Unmarshal(body, &response); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	products := response.Rows
	if len(products) == 0 {
		t.Errorf("Got empty result")
	}
}
