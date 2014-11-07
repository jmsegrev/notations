package notations

import "testing"

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
