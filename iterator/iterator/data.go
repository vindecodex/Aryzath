package iterator

type Data struct {
	value string
}

func NewData(value string) *Data {
	return &Data{value: value}
}

func (d *Data) SetData(value string) {
	d.value = value
}

func (d *Data) GetData() string {
	return d.value
}
