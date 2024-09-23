package frete_test

import (
	"context"
	"gofrete/frete"
	"gofrete/fretetypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularFrete(t *testing.T) {
	input := fretetypes.Params{
		CepOrigem:        "30626380",
		CepDestino:       "30626370",
		Peso:             "300",
		Comprimento:      1.2,
		Altura:           1.3,
		Largura:          1.2,
		ValorDeclarado:   10.00,
		AvisoRecebimento: "",
		CodigoEmpresa:    "",
		CodigoFormato:    1,
		Senha:            "",
	}

	ctx := context.Background()
	r := frete.MakeFreteRequest(input, fretetypes.SvcSEDEXVarejo)
	resp, err := frete.CalcularFrete(ctx, r)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
