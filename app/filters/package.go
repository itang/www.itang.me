package filters

import (
	"app/utils"
)

var (
	Full = utils.Comp(CatchErr, AppengineContext)
)
