/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var gjsonCmd = &cobra.Command{
	Use:   "gjson",
	Short: "Examples of gjson and sjson usage",
	Long:  `Examples of gjson and sjson usage ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- gjson")
		GetAValue()
		UsePathSyntax()
		XformDemo()
	},
}

func init() {
	rootCmd.AddCommand(gjsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetAValue() {

	fmt.Println("\n--- GetAValue:")
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

	// Get searches json for the specified path.
	// A path is in dot syntax, such as "name.last" or "age".
	// When the value is found it's returned immediately.
	value := gjson.Get(json, "name.last")
	fmt.Println(value.String())
}

func UsePathSyntax() {
	fmt.Println("\n--- UsePathSyntax:")

	const json = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

	// A path is a series of keys separated by a dot.
	// A key may contain special wildcard characters '*' and '?'.
	// To access an array value use the index as the key.
	// To get the number of elements in an array or to access a child path,
	// use the '#' character.
	// The dot and wildcard characters can be escaped with '\'.

	print := func(selector string) {
		value := gjson.Get(json, selector)
		fmt.Println(selector+":", value.String())
	}

	print("name.first")       // "Tom"
	print("name.last")        // "Anderson"
	print("age")              // "37"
	print("children")         // ["Sara","Alex","Jack"]
	print("children.#")       // "3"
	print("children.1")       // "Alex"
	print("child*.2")         // "Jack"
	print("c?ildren.0")       // "Sara"
	print("fav\\.movie")      // "Deer Hunter"
	print("friends.#.first")  // ["Dale","Roger","Jane"]
	print("friends.1.last")   // "Craig"
	print("friends.1.nets.#") // "2"
	print("friends.1.nets.1") // "fb"

}

func XformDemo() {
	inputJson := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "James", "last": "Murphy"},
    {"first": "Roger", "last": "Craig"}
  ]
}`
	outputJson, shouldReturn := Xform(inputJson)
	if shouldReturn {
		return
	}

	// print the inpt and output json
	fmt.Println("\n--- Xform:")
	fmt.Println("inputJson:", inputJson)
	fmt.Println("outputJson:", outputJson)
}

func Xform(inputJson string) (string, bool) {
	// Get the "name.first" and "children" fields from the JSON string
	// Map the "friends" array to an array of friend names
	// Create a new JSON object with the extracted fields
	nameFirst := gjson.Get(inputJson, "name.first").String()

	children := gjson.Get(inputJson, "children").Array()
	var childrenArr []interface{}
	for _, child := range children {
		childrenArr = append(childrenArr, child.String())
	}
	var friends []string
	gjson.Get(inputJson, "friends").ForEach(func(_, v gjson.Result) bool {
		friendName := v.Get("first").String()
		friends = append(friends, friendName)
		return true
	})

	newJSON, _ := sjson.Set("", "name", nameFirst)
	newJSON, _ = sjson.Set(newJSON, "children", childrenArr)
	newJSON, _ = sjson.Set(newJSON, "friends", friends)

	outputJson := strings.TrimSpace(newJSON)
	return outputJson, false
}
