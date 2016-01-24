package logica

import (
	"kentia/modelo"
)

//MapaInfo se utiliza para mandarle datos a los templates
type MapaInfo map[string]interface{}

/*SetInformacion a un mapa dado le agrega la informacion recibida*/
func (mapa MapaInfo) SetInformacion(args ...interface{}) {
	for i := 0; i < len(args); i += 2 {
		mapa[args[i].(string)] = args[i+1]
	}
}

//ObtenerDatosRegistroPrenda le agraga datos necesarios al mapa para el registro de una prenda
func (mapa MapaInfo) ObtenerDatosRegistroPrenda() {
	mapa.SetInformacionMapa(
		"climas", modelo.ConsultarClimas(),
		"colores", modelo.ConsultarColores(),
		"ocasiones", modelo.ConsultarOcasiones(),
		"tiposPrenda", modelo.ConsultarTiposPrenda(),
	)
}
