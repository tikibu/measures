package measures

import "time"

var Measures = CreateMeasureStore(NewMemIncrementer(nil), time.Second)
