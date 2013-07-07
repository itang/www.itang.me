package filters

import (
  u "app/utils"
)

var (
  Full = u.Comp(CatchErr, AppengineContext)
)
