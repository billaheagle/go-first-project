package go_first_project

type Nasabah struct {
	Name string
}

func SayHelloToNasabah(nasabah Nasabah) (string, bool) {
	if nasabah.Name == "" {
		return "Hello Guest", false
	}

	return "Hello Nasabah " + nasabah.Name + " Selamat Datang !!!", true
}
