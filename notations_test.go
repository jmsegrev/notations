package notations

import (
	"fmt"
	"testing"

	"github.com/jmsegrev/notations"
)

func TestExtract(t *testing.T) {

	text := `@mention @anothermention #tag $1223.93 some@email.com @ v @`

	notationsMap := Extract(text, "@#$")

	if len(notationsMap["@"]) != 2 {
		t.Error("Extracted invalid notations")
	}

	n := notationsMap["@"][0]
	if n.Value != "mention" && n.Slice.Begin != 0 &&
		text[n.Slice.Begin:n.Slice.End] != n.Value {
		t.Error("Failed to extract mention at the beginning of the text")
	}

	n = notationsMap["@"][1]
	if n.Value != "mention" && text[n.Slice.Begin:n.Slice.End] != n.Value {
		t.Error("Failed to extract mention in middle of the text")
	}

	n = notationsMap["#"][0]
	if n.Value != "tag" && text[n.Slice.Begin:n.Slice.End] != n.Value {
		t.Error("Failed to extract tag")
	}

	n = notationsMap["$"][0]
	if n.Value != "1223.93" && text[n.Slice.Begin:n.Slice.End] != n.Value {
		t.Error("Failed to extract amount")
	}

}

func ExampleExtract() {

	txt := `@username assume this is a reply, by @anotherusername that has
	some #tag1 and #tag2. It also works the same way for %any &other *symbol,
	but you provably will only use things like $2377.12.`

	notationsMap := notations.Extract(txt, "@#$%&*$")
	for key, notations := range notationsMap {
		fmt.Printf("%s: ", key)
		for _, notation := range notations {
			fmt.Printf("%s ", notation.Value)
		}
		fmt.Println("")
	}

	// Output:
	// @: username anotherusername
	// #: tag1 tag2
	// %: any
	// &: other
	// *: symbol
	// $: 2377.12

}
