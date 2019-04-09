package measures

type Incrementer interface {
	IncAndGet(key string, by float64) (res float64, err error)
	Get(key string) (res float64, err error)
}

type Measure interface {
	Inc(v float64)
	Get() (ret float64)
}

/*
type MeasureStoreInterface interface{
	Inc(key string, v float64)
	Get(key string) (ret float64, err error)
}*/
