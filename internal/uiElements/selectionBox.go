package uiElements

type Menu struct {
	XLocation int
	YLocation int
	Current   int
	Items     []string
}

func NewMenu(xLoc, yLoc int, items []string) *Menu {
	var menu *Menu = new(Menu)
	menu.XLocation = xLoc
	menu.YLocation = yLoc
	menu.Current = 0
	menu.Items = items
	return menu
}
