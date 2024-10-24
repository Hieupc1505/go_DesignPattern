package observer

type Item struct {
	name      string
	observers []Observer
}

func NewItem(n string) *Item {
	return &Item{
		name: n,
	}
}

func (i *Item) Register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) deregister(o Observer) {
	for idx, ctm := range i.observers {
		if ctm == o {
			i.observers = append(i.observers[:idx], i.observers[idx+1:]...)
			return
		}
	}
}

func (i *Item) UpdateStockItem() {
	i.notifyAll()
}

func (i *Item) notifyAll() {
	for _, ctm := range i.observers {
		ctm.updateItem(i.name)
	}
}
