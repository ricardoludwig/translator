package model

import (
	"testing"
)

func TestLanguagesSortByCode(t *testing.T) {

	pt := Language{
		Name: "Portugues",
		Code: "pt",
	}

	en := Language{
		Name: "English",
		Code: "en",
	}

	languages := Languages{pt, en}

	languages.SortByCode()
	//SortByCode(languages)

	if languages[0].Code != "en" {
		t.Error("Fail, expected english code en but got " + languages[0].Code)
		return
	}

}
