package iterator

type Collection interface {
	CreateIterator() Iterator
}

type DataCollection struct {
	Data []*Data
}

func (d *DataCollection) CreateIterator() Iterator {
	return &DataIterator{
		Index: 0,
		Data:  d.Data,
	}
}
