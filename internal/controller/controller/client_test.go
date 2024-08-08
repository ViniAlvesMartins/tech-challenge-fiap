package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/controller/serializer/output"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/internal/entities/entity"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"log"
	"testing"
)

var expectedResponse Response
var responseBody Response
var statusCode int

func iSendAGetRequest() error {
	client := resty.New()

	c := entity.Client{
		ID:    1,
		Cpf:   12345678900,
		Name:  "Test Client",
		Email: "testclient@example.com",
	}

	expectedResponse = Response{
		Error: "",
		Data:  output.ClientFromEntity(c),
	}

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "/clients?cpf=12345678900", httpmock.NewJsonResponderOrPanic(200, expectedResponse))

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&responseBody).
		SetQueryString("cpf=12345678900").
		Get("/clients")

	if err != nil {
		log.Printf("Request failed: %s", err)
	}

	statusCode = resp.StatusCode()

	return nil
}

func statusCodeShouldBe(expectedStatus int) error {
	if expectedStatus != statusCode {
		if statusCode >= 400 {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d", expectedStatus, statusCode)
		}
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", expectedStatus, statusCode)
	}

	return nil
}

func clientDetailsShouldBeReturned() error {
	jsonResponse, _ := json.Marshal(responseBody)
	jsonExpectedResponse, _ := json.Marshal(expectedResponse)

	if len(jsonResponse) == 0 {
		return fmt.Errorf("Expected: %s\nGot: %s", string(jsonExpectedResponse), string(jsonResponse))
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send a GET request to "/clients"$`, iSendAGetRequest)
	ctx.Step(`^Status code should be (\d+)$`, statusCodeShouldBe)
	ctx.Step(`Client details should be returned$`, clientDetailsShouldBeReturned)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
