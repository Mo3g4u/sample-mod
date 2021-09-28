package data

var (
	dataList = map[string]string{}
)

func SetData(k, v string) map[string]string {
	dataList[k] = v
	return dataList
}
