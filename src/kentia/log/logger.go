package log

import (
	"fmt"
	"os"
	"time"
)

const fileName = "kentia.log"

//RegistrarError guarda el error en un log, sino existe el log creará el archivo.
func RegistrarError(err error) {
	file, e := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if e != nil {
		os.Create(fileName)
		file, e = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
		fmt.Println(e)
	}
	defer file.Close()
	file.WriteString(time.Now().String() + ":" + err.Error() + "\n")
}
