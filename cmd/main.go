package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/mahsamnz/check24/internal/factories"
	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/providers/acme"
	"github.com/mahsamnz/check24/internal/serializers"
	"github.com/mahsamnz/check24/internal/utils"
)

var factory *factories.ServiceProviderFactory

func init() {
	// Initialization factory
	factory = factories.NewServiceProviderFactory()

	// Register ACME service provider with Xml serializer
	acmeService := acme.NewACMEServiceProvider(serializers.NewXMLSerializer())
	factory.RegisterService(acmeService)
}

func main() {

	// Check if file argument and service provider is provided
	if len(os.Args) != 3 {
		log.Fatal("Arguments are not provided, Usage: main.go filename provider")
	}

	filename := os.Args[1]
	serviceProviderId := os.Args[2]

	// Create Service provider by its identifier
	serviceProvider, err := factory.GetProvider(serviceProviderId)
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	fileContent, err := utils.ReadJSONFile(filename)
	if err != nil {
		log.Fatalf("Failed to get file content: %v", err)
	}

	var carInsuranceRequest models.CarInsuranceRequest
	if err := json.Unmarshal(fileContent, &carInsuranceRequest); err != nil {
		log.Fatalf("invalid JSON: %v", err)
	}

	err = utils.Validate(carInsuranceRequest)
	if err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	serviceProvider.MapData(carInsuranceRequest)

	// Serialize to XML
	xmlData, err := serviceProvider.SerializeData()
	if err != nil {
		log.Fatalf("Failed to serialize to data: %v", err)
	}
	log.Printf("Data: %s", string(xmlData))
}
