package prettifiers

type Prettifier interface {
	Prettify(string) string
}
