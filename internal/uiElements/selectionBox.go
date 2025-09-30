package uiElements

type Menu struct {
	XLocation int
	YLocation int
	Width     int
	Height    int
	Current   int
	Items     []string
}

func NewMenu(xLoc, yLoc, w, h int, items []string) *Menu {
	var menu *Menu = new(Menu)
	menu.XLocation = xLoc
	menu.YLocation = yLoc
	menu.Width = w
	menu.Height = h
	menu.Current = 0
	menu.Items = items
	return menu
}
