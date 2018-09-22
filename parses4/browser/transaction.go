package browser

import (
	"github.com/lpuig/scopelecspi/spis4"
	"time"
)

type Transaction struct {
	Name string
	Date time.Time
	Req  *spis4.S4ReqZek
	Resp *spis4.S4RespZek
}

type Transactions []Transaction
