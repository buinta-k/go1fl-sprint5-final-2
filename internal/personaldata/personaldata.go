package personaldata

type Personal struct {
	Name string
	Weight int
	Height int
}

func (p Personal) Print() {
	fmt.Printf(
		"Имя: %s\nВес: %.2f кг.\nРост: %.2fм.\n",
		p.Name,
		p.Weight,
		p.Height
	)
}
