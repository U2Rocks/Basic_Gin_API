package itemfeed

type Cvsitem struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Date string  `json:"date"`
}

type Itemlist struct {
	Cvsitems []Cvsitem `json:"Cvsitems"`
}

func New() *Itemlist {
	return &Itemlist{
		Cvsitems: []Cvsitem{},
	}
}

func (i *Itemlist) GetItemsBought() []Cvsitem {
	return i.Cvsitems
}

func (i *Itemlist) AddItemtoList(item Cvsitem) {
	i.Cvsitems = append(i.Cvsitems, item)
}
