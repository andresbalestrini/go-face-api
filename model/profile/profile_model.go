package profile

type Familiar struct {
    Name string `json:"name"`
    Id string `json:"id"`
    Relationship string `json:"relationship"`

}

type Arrayfamily struct {
	Array []Familiar `json:"data"`
}

type Dataprofile struct {

    Nombre string `json:"name"`
    Genero string `json:"gender"`
    Ubicacion string `json:"locale"`
    Cumple string `json:"birthday"`
}

type Response struct {
    Profile Dataprofile
    Family  Arrayfamily
}