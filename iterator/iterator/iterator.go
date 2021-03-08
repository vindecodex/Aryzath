package iterator

type Iterator interface {
	HasNext() bool
	GetNext() *Data
}

type DataIterator struct {
	Index int
	Data  []*Data
}

func (d *DataIterator) HasNext() bool {
	if d.Index < len(d.Data) {
		return true
	}
	return false
}

func (d *DataIterator) GetNext() *Data {
	if d.HasNext() {
		data := d.Data[d.Index]
		d.Index = d.Index + 1
		return data
	}
	return nil
}
