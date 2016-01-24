package controlador

import (
	"kentia/modelo"
)

//MapaInfo se utiliza para mandarle datos a los templates.
type MapaInfo map[string]interface{}

//SetInformacion a un mapa dado le agrega la informacion recibida.
func (mapa MapaInfo) SetInformacion(args ...interface{}) {
	for i := 0; i < len(args); i += 2 {
		mapa[args[i].(string)] = args[i+1]
	}
}

//ObtenerDatosRegistroPrenda le agrega datos necesarios al mapa para el registro de una prenda.
func (mapa MapaInfo) ObtenerDatosRegistroPrenda() {
	mapa.SetInformacion(
		"climas", modelo.ConsultarClimas(),
		"colores", modelo.ConsultarColores(),
		"ocasiones", modelo.ConsultarOcasiones(),
		"tiposPrenda", modelo.ConsultarTiposPrenda(),
	)
}

//ObtenerDatosCombinacion obtiene los datos para mostrar en el template.
func (mapa MapaInfo) ObtenerDatosCombinacion(usuarioID string) {
	mapa.SetInformacion(
		"mejores", GenerarMejorCombinacion(usuarioID),
	)
}
