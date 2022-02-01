package tests

import (
	"crawler/app/database"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	database, err := database.Connect()

	t.Run("Test connect database", func(t *testing.T) {
		if database != nil {
			t.Error()
		}
	})
	t.Run("Test failed datavavase", func(t *testing.T) {
		if err == nil {
			t.Error()
		}
	})
}
