# Logify

## Características

Logify es una biblioteca de registro simple y configurable para aplicaciones Go. Ofrece:
- Colores configurables para diferentes tipos de log.
- Estructura de logger para gestionar módulos, consumidores y entornos.
- Salida de log formateada con fecha y hora, tipo de log, consumidor y entorno.
- Diferentes modos de registro para adaptarse a entornos de desarrollo y producción.

## Instalación

Para instalar Logify, puedes usar `go get`:

```sh
go get github.com/Jdcabreradev/logify/v3
```

## Casos de uso

Logify se puede utilizar para registrar información, advertencias, errores y mensajes de depuración en aplicaciones Go. Es útil para:

- Monitoreo y depuración de aplicaciones.
- Registro de eventos importantes.
- Análisis post-mortem de errores.

## Funciones y Tipo

### Funciones

- New(module, consumer, logDirPath string, env logify_enums.LogifyMode) *Logger: Crea una nueva instancia de Logger.
- (*Logger) Log(logType logify_enums.LogType, message string): Imprime un mensaje formateado en la salida estándar y opcionalmente lo guarda en un archivo de log.
- SetColor(logType logify_enums.LogType, color string): Permite cambiar el color de un tipo de log específico.


### Tipo

- LogType: Un tipo enumerado para los diferentes tipos de logs (INFO, WARNING, ERROR, DEBUG).
- LogifyMode: Un tipo enumerado para los diferentes modos de registro (LogModeDev, LogModeRelease, LogModeVerbose, LogModeHidden).

## Ejemplos de Uso

### Crear un Logger y registrar mensajes

```go
package main

import (
    "github.com/Jdcabreradev/logify/v3"
    "github.com/Jdcabreradev/logify/v3/enums"
)

func main() {
    logger := logify.New("MiModulo", "MiConsumidor", "./logs/", enums.LogModeDev)

    logger.Log(enums.INFO, "Este es un mensaje informativo.")
    logger.Log(enums.WARNING, "Este es un mensaje de advertencia.")
    logger.Log(enums.ERROR, "Este es un mensaje de error.")
    logger.Log(enums.DEBUG, "Este es un mensaje de depuración.")
}
```

### Cambiar el color de un tipo de log

```go
package main

import (
	"github.com/Jdcabreradev/logify/v3"
	"github.com/Jdcabreradev/logify/v3/enums"
	"github.com/Jdcabreradev/logify/v3/colors"
)

func main() {
	logify.SetColor(enums.INFO, colors.Cyan) // Cambiar color de INFO a cyan
	logger := logify.New("MiModulo", "MiConsumidor", "./logs/", enums.LogModeDev)

	logger.Log(enums.INFO, "Este es un mensaje informativo en cyan.")
}
```

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request en GitHub.

1. Haz un fork del proyecto.
2. Crea una nueva rama (git checkout -b feature/nueva-funcionalidad).
3. Realiza los cambios y haz commits (git commit -am 'Añadir nueva funcionalidad').
4. Empuja la rama (git push origin feature/nueva-funcionalidad).
5. Abre un pull request.

## Licencia

Este proyecto está licenciado bajo la [Licencia MIT](https://github.com/Jdcabreradev/logify/blob/main/LICENSE). Consulta el archivo LICENSE para más detalles.