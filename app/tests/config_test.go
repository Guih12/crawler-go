package tests

import (
	"crawler/app/utils"
	"testing"
)

func TestCofigEnviroments(t *testing.T) {
	databaseUrl, _ := utils.LoadingEnviroments()
	urlExpected := databaseUrl

	if urlExpected != databaseUrl {
		t.Fail()
	}
}
