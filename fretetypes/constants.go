package fretetypes

import "fmt"

// TipoServico representa os tipos de serviço (numérico)
type TipoServico string

func (svct TipoServico) String() string {
	switch svct {
	case SvcSEDEXVarejo:
		return "SEDEX Varejo"
	case SvcSEDEXACobrarVarejo:
		return "SEDEX Varejo (a cobrar)"
	case SvcSEDEX10Varejo:
		return "SEDEX 10 Varejo"
	case SvcSEDEXHojeVarejo:
		return "SEDEX Hoje Varejo"
	case SvcSEDEXComContrato:
		return "SEDEX"
	case SvcPACVarejo:
		return "PAC Varejo"
	case SvcPACComContrato:
		return "PAC"
	default:
		return fmt.Sprintf("the %s type cannot found", string(svct))
	}
}

// update 05/05/2017
// Correios alterou os códigos de serviço:
// 04510 – PAC sem contrato (era 41106)
// 04669 – PAC com contrato
// 04014 – SEDEX sem contrato (era 40010)
// 04162 – SEDEX com contrato
//
// Não há nenhuma documentação sobre a mudança no momento.

// Todos os tipos de serviço
// https://www.correios.com.br/para-voce/correios-de-a-a-z/pdf/calculador-remoto-de-precos-e-prazos/manual-de-implementacao-do-calculo-remoto-de-precos-e-prazos
const (
	SvcSEDEXVarejo        TipoServico = "04014"
	SvcSEDEXACobrarVarejo TipoServico = "40045"
	SvcSEDEX10Varejo      TipoServico = "40215"
	SvcSEDEXHojeVarejo    TipoServico = "40290"
	SvcSEDEXComContrato   TipoServico = "04162"
	SvcPACVarejo          TipoServico = "04510"
	SvcPACComContrato     TipoServico = "04669"
)

// RequestMode altera o metodo utilizado p/ obter o frete
type RequestMode int

const (
	// RequestModeAuto escolher o modo automaticamente
	RequestModeAuto RequestMode = iota
	// RequestModeSingle sempre envia um request por tipo de serviço
	RequestModeSingle
	// RequestModeCombined sempre envia um único request com todos tipos de serviço
	RequestModeCombined
)

// TipoErro representa uma resposta de erro dos Correios (código interno)
type TipoErro int

// Todos os tipos de erros possíveis que a API dos Correios pode retornar
const (
	ErrTipoServicoInvalido          TipoErro = -1
	ErrCepOrigemInvalido            TipoErro = -2
	ErrCepDestinoInvalido           TipoErro = -3
	ErrCepPesoExcedido              TipoErro = -4
	ErrValorDeclaradoAlto10k        TipoErro = -5
	ErrServicoIndisponivelTrecho    TipoErro = -6
	ErrValorDeclaradoObrigatorio    TipoErro = -7
	ErrMaoPropriaIndisponivel       TipoErro = -8
	ErrAvisoRecebimentoIndisponivel TipoErro = -9
	ErrPrecificacaoIndisponivel     TipoErro = -10
	ErrInformarDimensoes            TipoErro = -11
	ErrComprimento                  TipoErro = -12
	ErrLargura                      TipoErro = -13
	ErrAltura                       TipoErro = -14
	ErrComprimento105               TipoErro = -15  // > 105cm
	ErrLargura105                   TipoErro = -16  // > 105cm
	ErrAltura105                    TipoErro = -17  // > 105cm
	ErrAlturaInferior               TipoErro = -18  // < 2cm
	ErrLarguraInferior              TipoErro = -20  // < 11cm
	ErrComprimentoInferior          TipoErro = -22  // < 16cm
	ErrDimensoesSoma                TipoErro = -23  // A soma resultante do comprimento + largura + altura não deve superar a 200 cm
	ErrComprimento2                 TipoErro = -24  // WTF (ver -12)
	ErrDiametro                     TipoErro = -25  // Diâmetro inválido
	ErrComprimento3                 TipoErro = -26  // WTF (ver -12)
	ErrDiametro2                    TipoErro = -27  // ?
	ErrComprimento4                 TipoErro = -28  // O comprimento não pode ser maior que 105 cm.
	ErrDiametro91                   TipoErro = -29  // O diâmetro não pode ser maior que 91 cm.
	ErrComprimento18                TipoErro = -30  // O comprimento não pode ser inferior a 18 cm.
	ErrDiametro5                    TipoErro = -31  // O diâmetro não pode ser inferior a 5 cm.
	ErrSomaDiametro                 TipoErro = -32  // A soma resultante do comprimento + o dobro do diâmetro não deve superar a 200 cm
	ErrSistemaIndisponivel          TipoErro = -33  // Sistema temporariamente fora do ar. Favor tentar mais tarde.
	ErrCodigoOuSenha                TipoErro = -34  // Código Administrativo ou Senha inválidos.
	ErrSenha                        TipoErro = -35  // Senha incorreta.
	ErrSemContrato                  TipoErro = -36  // Cliente não possui contrato vigente com os Correios.
	ErrSemServicoAtivo              TipoErro = -37  // Cliente não possui serviço ativo em seu contrato.
	ErrServicoIndisponivelAdmin     TipoErro = -38  // Serviço indisponível para este código administrativo.
	ErrPesoExcedidoEnvelope         TipoErro = -39  // Peso excedido para o formato envelope
	ErrInformarDimensoes2           TipoErro = -40  // Para definicao do preco deverao ser informados, tambem, o comprimento e a largura e altura do objeto em centimetros (cm).
	ErrComprimento60                TipoErro = -41  // O comprimento nao pode ser maior que 60 cm.
	ErrComprimento16                TipoErro = -42  // (repetido) O comprimento nao pode ser inferior a 16 cm.
	ErrComprimentoLargura120        TipoErro = -43  // A soma resultante do comprimento + largura nao deve superar a 120 cm
	ErrLarguraInferior2             TipoErro = -44  // < 11cm
	ErrLarguraSuperior60            TipoErro = -44  // > 60cm
	ErrErroCalculoTarifa            TipoErro = -888 // Erro ao calcular a tarifa
	ErrLocalidadeOrigem             TipoErro = 6    // Localidade de origem não abrange o serviço informado
	ErrLocalidadeDestino            TipoErro = 7    // Localidade de destino não abrange o serviço informado
	ErrServicoIndisponivelTrecho2   TipoErro = 8    // 008 Serviço indisponível para o trecho informado
	ErrAreaDeRiscoCEPInicial        TipoErro = 9    // 009 CEP inicial pertencente a Área de Risco.
	ErrAreaPrazoDiferenciado        TipoErro = 10   // Área com entrega temporariamente sujeita a prazo diferenciado.
	ErrAreaDeRiscoCEPs              TipoErro = 11   // CEP inicial e final pertencentes a Área de Risco
	ErrIndisponivel                 TipoErro = 7    // Serviço indisponível, tente mais tarde
	ErrIndeterminado                TipoErro = 99   // Outros erros diversos do .Net // ¯\_(ツ)_/¯
)
