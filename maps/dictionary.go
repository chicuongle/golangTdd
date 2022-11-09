package maps

type Dictionary map[string]string

func (m Dictionary) Search(keyword string) string {
	return m[keyword]
}
func (d Dictionary) Add(key string,value string){
	d[key] = value
}
