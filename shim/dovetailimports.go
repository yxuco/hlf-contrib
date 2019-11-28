package main

import (
	_ "github.com/yxuco/hlf-contrib/fabric/activity/cid"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/delete"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/endorsement"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/get"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/getbycompositekey"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/gethistory"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/getrange"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/invokechaincode"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/put"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/putall"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/query"
	_ "github.com/yxuco/hlf-contrib/fabric/activity/setevent"
	_ "github.com/yxuco/hlf-contrib/fabric/trigger/transaction"
	_ "github.com/project-flogo/core/app/propertyresolver"
)
