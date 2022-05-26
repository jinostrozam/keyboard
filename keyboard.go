package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	// capturar entrada teclado
	reader := bufio.NewReader(os.Stdin)
	// leer captura hasta presionar enter
	input, err := reader.ReadString('\n')
	// revisar errores captura
	if err != nil {
		return 0, err
	}

	// eliminar espacios en blanco
	input = strings.TrimSpace(input)
	// transformar texto captura en numero
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil

}
