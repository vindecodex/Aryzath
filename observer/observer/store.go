package observer

import "fmt"

type Store struct {
	Observer []Observer
	Product  string
}

func NewStore(product string) *Store {
	return &Store{Product: product}
}

func (s *Store) Register(observer Observer) {
	s.Observer = append(s.Observer, observer)
}

func (s *Store) Unregister(observerID int) {
	for index, observer := range s.Observer {
		if observer.GetId() == observerID {
			s.Observer = append(s.Observer[:index], s.Observer[index+1:]...)
		}
	}
}

func (s *Store) Notify() {
	fmt.Println("----------Notification----------")
	for _, observer := range s.Observer {
		observer.GetNotified(s.Product)
	}
}
