package frete

import (
	"fmt"
	"gofrete/fretetypes"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"context"
)

// MakeFreteRequest ...
func MakeFreteRequest(input fretetypes.Params, svs ...fretetypes.TipoServico) (nReq *fretetypes.Request) {
	nReq = &fretetypes.Request{
		CepOrigem:        FilterCEP(input.CepOrigem),
		CepDestino:       FilterCEP(input.CepDestino),
		PesoKg:           input.Peso,
		ComprimentoCm:    input.Comprimento,
		AlturaCm:         input.Altura,
		LarguraCm:        input.Largura,
		Servicos:         svs,
		ValorDeclarado:   input.ValorDeclarado,
		AvisoRecebimento: input.AvisoRecebimento,
		CdEmpresa:        input.CodigoEmpresa,
		DsSenha:          input.Senha,
	}

	return
}

// FilterCEP removes non numbers from a CEP.
func FilterCEP(v string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, v)
}

func doRequest(ctx context.Context, method string, url string, headers ...map[string]string) (io.ReadCloser, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	for _, values := range headers {
		for name, value := range values {
			req.Header.Add(name, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("The status code is not correct. url [%s] - status_code [%d]", url, resp.StatusCode))
		err = fmt.Errorf("The status code is not correct. status_code [%d]", resp.StatusCode)
	}

	return resp.Body, err
}

// createQuery cria a query de consulta a partir da interface informada
func createQuery(req *fretetypes.Request) (v url.Values) {
	v = url.Values{}
	v.Set("sCepOrigem", req.CepOrigem)
	v.Set("sCepDestino", req.CepDestino)
	v.Set("nVlPeso", req.PesoKg)
	v.Set("nCdFormato", "1")
	v.Set("nVlComprimento", strconv.FormatFloat(req.ComprimentoCm, 'f', -1, 64))
	v.Set("nVlAltura", strconv.FormatFloat(req.AlturaCm, 'f', -1, 64))
	v.Set("nVlLargura", strconv.FormatFloat(req.LarguraCm, 'f', -1, 64))
	// v.Set("StrRetorno", "xml")
	svcs := make([]string, len(req.Servicos))
	for k, v := range req.Servicos {
		svcs[k] = string(v)
	}
	v.Set("nCdServico", strings.Join(svcs, ","))
	v.Set("nVlValorDeclarado", strconv.FormatFloat(req.ValorDeclarado, 'f', -1, 64))
	if len(req.AvisoRecebimento) != 0 {
		v.Set("sCdAvisoRecebimento", req.AvisoRecebimento)
	}

	if len(req.CdEmpresa) > 0 {
		v.Set("nCdEmpresa", req.CdEmpresa)
		v.Set("sDsSenha", req.DsSenha)
	}

	return
}
