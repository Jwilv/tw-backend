package routers

import (
	"errors"
	"strings"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// variable que almacena el email de user
var Email string

// variable que almacena el id del user
var IDUser string

// procesa si el token recivido es valido
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	//key del token
	key := []byte("oqwepFJO03-49RFSAFN0123@$WEWQAsfoajf")

	//puntero a la posicion de claim para guardar en objeto mas adelante
	claims := &models.Claim{}

	//split para dividir en token en 2 y separar la palabra Bearer
	splitToken := strings.Split(token, "Bearer")

	//medimos si el largo del split es de 2 para saber si tiene el formato  valido
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	//tomamos la pocicion donde esta el token y le sacamos los espacios
	token = strings.TrimSpace(splitToken[1])

	//verificamos el token, mandando el token, la estructura para guardar los datos que es un puntero al modelo
	// y la funcion de err po si el token no es valido
	tokenVerify, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return key, nil
	})

	//verificamos que no nos llegue error, si pasa retornamos error
	if err != nil {
		return claims, false, string(""), err
	}

	//verificamos que el token sea vakido, si no lo es devolvemos error
	if !tokenVerify.Valid {
		return claims, false, string(""), errors.New("token invaldio")
	}

	//tomamos el email que se guardo en claims y lo buscamos en la base de datos
	_, exist, id := db.CheckExisEmail(claims.Email)

	//si el email no existe devolvemos error
	if !exist {
		return claims, false, string(""), errors.New("no se encontro un user con el token")
	}

	//guardamos el email y el id, guardamos el id ya pasado a string por la funcion de CheckExisEmail
	Email = claims.Email
	IDUser = id

	//retirnamos el claim que contiene el email y el id, exist que nos dice que el use existe, IDUser que es el id como string
	// y el err lo ponemos como nil
	return claims, exist, IDUser, nil

}
