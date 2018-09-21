package browser

import "github.com/lpuig/scopelecspi/spis4"

type Transaction struct {
	Req  *spis4.S4ReqZek
	Resp *spis4.S4RespZek
}
