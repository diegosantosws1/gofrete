package frete

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"gofrete/fretetypes"
	"io"
	"log"
	"net/http"
)

// CalcularFrete envia o request p/ calcular o frete utilizando
// um *Request
// http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?sCepOrigem=01243000&sCepDestino=04041002&nVlPeso=1&nCdFormato=1&nVlComprimento=16&nVlAltura=5&nVlLargura=11&StrRetorno=xml&nCdServico=40010,41106&nVlValorDeclarado=0
func CalcularFrete(ctx context.Context, req *fretetypes.Request) (*fretetypes.FreteResponse, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}

	if len(req.Servicos) > 1 && ((req.Mode == fretetypes.RequestModeAuto && req.CdEmpresa == "") || (req.Mode == fretetypes.RequestModeSingle)) {
		reqs := make([]*fretetypes.Request, len(req.Servicos))
		for k := range req.Servicos {
			clone := &fretetypes.Request{
				CepOrigem:        req.CepOrigem,
				CepDestino:       req.CepDestino,
				PesoKg:           req.PesoKg,
				ComprimentoCm:    req.ComprimentoCm,
				AlturaCm:         req.AlturaCm,
				LarguraCm:        req.LarguraCm,
				Servicos:         []fretetypes.TipoServico{req.Servicos[k]},
				ValorDeclarado:   req.ValorDeclarado,
				AvisoRecebimento: req.AvisoRecebimento,
				CdEmpresa:        req.CdEmpresa,
				DsSenha:          req.DsSenha,
			}
			reqs[k] = clone
		}

		r00 := &fretetypes.FreteResponse{
			Servicos: make(map[fretetypes.TipoServico]fretetypes.ServicoResponse),
		}

		for i, v := range reqs {
			rsp, err := CalcularFrete(ctx, v)
			if err != nil && len(reqs) == i+1 {
				return r00, err
			} else if err != nil {
				continue
			}
			for k2, v2 := range rsp.Servicos {
				r00.Servicos[k2] = v2
			}
		}

		return r00, nil
	}

	vl := createQuery(req)
	urlStr := fmt.Sprintf("%s/CalcPrecoPrazo?%s", fretetypes.FreteEndpoint, vl.Encode())
	resp, err := doRequest(ctx, http.MethodGet, urlStr)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	data, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	vlov := struct {
		XMLName string                   `xml:"Servicos"`
		Values  []fretetypes.ServicoResp `xml:"cServico"`
	}{}

	err = xml.Unmarshal(data, &vlov)
	if err != nil {
		log.Println(fmt.Sprintf("CORREIOS: %s", err.Error()))
		return nil, err
	}

	//
	output := &fretetypes.FreteResponse{
		Servicos: make(map[fretetypes.TipoServico]fretetypes.ServicoResponse),
	}
	//

	//	for _, v := range vlov.Values {
	//		v2 := fretetypes.ServicoResponse{}
	//		v2.Tipo = fretetypes.TipoServico(v.Codigo)
	//		v2.Preco = v.Valor
	//		v2.PrazoEntregaDias = v.PrazoEntrega
	//		v2.PrecoSemAdicionais, _ = decimal.NewFromString(fixWrongDecimals(v.ValorSemAdicionais))
	//		v2.PrecoMaoPropria, _ = decimal.NewFromString(fixWrongDecimals(v.ValorMaoPropria))
	//		v2.PrecoAvisoRecebimento, _ = decimal.NewFromString(fixWrongDecimals(v.ValorAvisoRecebimento))
	//		v2.PrecoValorDeclarado, _ = decimal.NewFromString(fixWrongDecimals(v.ValorValorDeclarado))
	//		v2.EntregaDomiciliar = (v.EntregaDomiciliar == "S")
	//		v2.EntregaSabado = (v.EntregaSabado == "S")
	//		if v.Erro != 0 {
	//			er9 := &fretetypes.ServicoResponseError{
	//				Codigo: fretetypes.TipoErro(v.Erro),
	//			}
	//			v2.Erro = er9
	//			v2.ErroMsg = v.MsgErro
	//		}
	//		output.Servicos[v2.Tipo] = v2
	//	}
	//
	return output, nil
}
