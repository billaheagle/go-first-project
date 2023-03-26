package go_first_project

type Nasabah struct {
	Name string
}

func SayHelloToNasabah(nasabah Nasabah) string {
	return "Hello Nasabah " + nasabah.Name
}
