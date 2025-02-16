package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mahsamnz/check24/internal/factories"
	"github.com/mahsamnz/check24/internal/models"
	"github.com/mahsamnz/check24/internal/providers/acme"
	"github.com/mahsamnz/check24/internal/serializers"
	"github.com/mahsamnz/check24/internal/utils"
)

const (
	outputFileName = "output"
)

var factory *factories.ServiceProviderFactory

func handleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func initializeFactory() *factories.ServiceProviderFactory {
	factory := factories.NewServiceProviderFactory()
	acmeService := acme.NewACMEServiceProvider(serializers.NewXMLSerializer())
	factory.RegisterService(acmeService)
	return factory
}

func init() {
	factory = initializeFactory()
}

func main() {

	filename, serviceProviderId, err := validateCommandArguments(os.Args)
	handleError("Invalid arguments", err)

	// Create Service provider by its identifier
	serviceProvider, err := factory.GetProvider(serviceProviderId)
	handleError("Failed to get provider", err)

	fileContent, err := utils.ReadJSONFile(filename)
	handleError("Failed to read file content", err)

	var carInsuranceRequest models.CarInsuranceRequest
	handleError("Invalid JSON", json.Unmarshal(fileContent, &carInsuranceRequest))

	handleError("Validation failed", utils.Validate(carInsuranceRequest))

	xmlData, err := serviceProvider.SerializeData(carInsuranceRequest)
	handleError("Failed to serialize data", err)

	outputFile, err := utils.CreateFile(outputFileName, serviceProvider.GetSerializer().GetFormat())
	handleError("Failed to create output file", err)
	defer outputFile.Close()

	_, err = outputFile.Write(xmlData)
	handleError("Failed to write to output file", err)

	log.Printf("data is written on %s.%s file successfully.", outputFileName, serviceProvider.GetSerializer().GetFormat())
}

func validateCommandArguments(args []string) (filename, providerID string, err error) {
	if len(args) != 3 {
		return "", "", fmt.Errorf("invalid number of arguments. Usage: main.go filename provider")
	}

	filename = args[1]
	if !strings.HasSuffix(filename, ".json") {
		return "", "", fmt.Errorf("file must be a JSON file")
	}

	providerID = args[2]

	return filename, providerID, nil
}
