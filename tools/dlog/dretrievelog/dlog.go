package dretrievelog

import (
	"github.com/filecoin-project/go-fil-markets/tools/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("retrieve")
}
