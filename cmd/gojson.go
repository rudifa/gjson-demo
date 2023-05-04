/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/spf13/cobra"
)

var gojsonCmd = &cobra.Command{
	Use:   "gojson",
	Short: "Decoding JSON to Structs and Maps",
	Long:  `Decoding JSON to Structs and Maps...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- gojson")

		JsonToStruct()
		JsonToArray()
		Invalid()
	},
}

func init() {
	rootCmd.AddCommand(gojsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gojsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gojsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func JsonToStruct() {
	type Bird struct {
		Species     string
		Description string
	}

	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println("--- Decoding JSON to Structs")
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)
	//Species: pigeon, Description: likes to perch on rocks
}

func JsonToArray() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds := result["birds"].(map[string]any)

	fmt.Println("--- Decoding JSON to Maps - Unstructured Data")
	for key, value := range birds {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
	//pigeon likes to perch on rocks
	//eagle bird of prey
}

func Invalid() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`

	fmt.Println("--- Invalid JSON:")
	if !json.Valid([]byte(birdJson)) {
		// handle the error here
		fmt.Println("invalid JSON string:", birdJson)
		return
	}
}
