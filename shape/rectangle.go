package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectangle *Rectangle) Area() float32 { //nama funcnya harus diawali huruf besar agar bisa di import di file lain
	return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Width + rectangle.Length)
}