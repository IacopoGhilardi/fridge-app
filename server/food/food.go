package food

type Food struct {
	Image    string `bson:image,omitempty`
	Name     string `bson:name,omitempty`
	Quantity int    `bson:quantity,omitempty`
}
