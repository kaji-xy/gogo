// Package toli offers information about our toli
package toli

// Title returns the title of this book.
func Title() string {
	return "GoGo Toli: " + subtitle()
}

func subtitle() string {
	return "Tolders on skates"
}
