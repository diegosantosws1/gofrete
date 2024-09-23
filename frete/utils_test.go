package frete_test

import (
	"gofrete/frete"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterCEP(t *testing.T) {
	caseTests := []struct {
		Name string
		Cep  string
		Exp  string
	}{
		{
			Name: "case 1",
			Cep:  "30626-360",
			Exp:  "30626360",
		},
		{
			Name: "case 2",
			Cep:  "30626380",
			Exp:  "30626380",
		},
	}

	for i := range caseTests {
		t.Run(caseTests[i].Name, func(t *testing.T) {
			got := frete.FilterCEP(caseTests[i].Cep)
			assert.Equal(t, caseTests[i].Exp, got, caseTests[i].Name)
		})
	}
}
