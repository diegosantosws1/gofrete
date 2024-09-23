package fretetypes

// FreteEndpoint o endpoint a ser utilizado para calcular o frete
var FreteEndpoint = "https://shopping.correios.com.br/wbm/shopping/script/calcprecoprazo.asmx"

// Request ...
type Request struct {
	CepOrigem        string        `url:"sCepOrigem"`
	CepDestino       string        `url:"sCepDestino"`
	PesoKg           string        `url:"nVlPeso"`
	ComprimentoCm    float64       `url:"nVlComprimento"`
	AlturaCm         float64       `url:"nVlAltura"`
	LarguraCm        float64       `url:"nVlLargura"`
	Servicos         []TipoServico `url:"nCdServico"`
	ValorDeclarado   float64       `url:"nVlValorDeclarado"`
	AvisoRecebimento string        `url:"sCdAvisoRecebimento"`
	CdEmpresa        string        `url:"nCdEmpresa"`
	DsSenha          string        `url:"sDsSenha"`
	CdFromato        int           `url:"nCdFormato"`
	Mode             RequestMode
}

// SetServicos troca os tipos de serviço a serem consultados
func (r *Request) SetServicos(srvs ...TipoServico) *Request {
	r.Servicos = make([]TipoServico, 0)
	r.Servicos = append(r.Servicos, srvs...)
	return r
}

// AppendServico anexa o serviço srv aos serviços a serem consultados
func (r *Request) AppendServico(srv TipoServico) *Request {
	for _, v := range r.Servicos {
		if v == srv {
			// already exists
			return r
		}
	}
	r.Servicos = append(r.Servicos, srv)
	return r
}

// FreteResponse resposta dos correios
type FreteResponse struct {
	Servicos map[TipoServico]ServicoResponse `json:"servicos,omitempty"`
}

// ServicoResponseError é a resposta de erro da API dos Correios
type ServicoResponseError struct {
	Codigo TipoErro `json:"code,omitempty"`
}

// ServicoResponse representa os dados retornados para um tipo de serviço
type ServicoResponse struct {
	Tipo                  TipoServico           `json:"type,omitempty"`
	Preco                 float64               `json:"price,omitempty"` // preço != valor != custo; deveria ser tudo preço
	PrazoEntregaDias      int                   `json:"delivery_days,omitempty"`
	PrecoSemAdicionais    float64               `json:"additional_price,omitempty"`
	PrecoMaoPropria       float64               `json:"price_mao_own,omitempty"`
	PrecoAvisoRecebimento float64               `json:"price_notice_receipt,omitempty"`
	PrecoValorDeclarado   float64               `json:"price_declared_value,omitempty"`
	EntregaDomiciliar     bool                  `json:"home_delivery,omitempty"`
	EntregaSabado         bool                  `json:"saturday_delivery,omitempty"`
	Erro                  *ServicoResponseError `json:"error,omitempty"`
	ErroMsg               string                `json:"error_msg,omitempty"`
}

// ServicoResp xml wrapper for ServicoResponse
type ServicoResp struct {
	Codigo                string `xml:"Codigo" json:"code,omitempty"`
	Valor                 string `xml:"Valor" json:"value,omitempty"`
	PrazoEntrega          int    `xml:"PrazoEntrega" json:"delivery_time,omitempty"`
	ValorSemAdicionais    string `xml:"ValorSemAdicionais" json:"additionals_values,omitempty"`
	ValorMaoPropria       string `xml:"ValorMaoPropria" json:"price_mao_own,omitempty"`
	ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento" json:"price_notice_receipt,omitempty"`
	ValorValorDeclarado   string `xml:"ValorValorDeclarado" json:"price_declared_value,omitempty"`
	EntregaDomiciliar     string `xml:"EntregaDomiciliar" json:"home_delivery,omitempty"`
	EntregaSabado         string `xml:"EntregaSabado" json:"saturday_delivery,omitempty"`
	Erro                  int    `xml:"Erro" json:"erro,omitempty"`
	MsgErro               string `xml:"MsgErro" json:"erro_msg,omitempty"`
}

// Params é a struct com os parametros da requisição
type Params struct {
	CodigoEmpresa    string
	Senha            string
	CepOrigem        string
	CepDestino       string
	Peso             string
	CodigoFormato    int
	Comprimento      float64
	Altura           float64
	Largura          float64
	Diametro         float64
	MaoPropria       string
	ValorDeclarado   float64
	AvisoRecebimento string
}
