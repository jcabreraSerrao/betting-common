package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	RoleId      int      `json:"role_id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	UserId      int      `json:"user_id"`
	GroupId     int      `json:"group_id"`
	NameGroup   string   `json:"name_group"`
}

type HasuraClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	UserId       string   `json:"x-hasura-user-id"`
	GroupId      string   `json:"x-hasura-group-id"`
	RoleId       string   `json:"x-hasura-role-id"`
}

func (data *JwtClaims) GenerateToken() (string, error) {
	config := GetConfig()
	hasuraClaims := HasuraClaims{
		AllowedRoles: append(data.Permissions, "prueba"),
		DefaultRole:  "prueba",
		UserId:       fmt.Sprintf("%d", data.UserId),
		GroupId:      fmt.Sprintf("%d", data.GroupId),
		RoleId:       fmt.Sprintf("%d", data.RoleId),
	}

	if len(hasuraClaims.AllowedRoles) == 0 {
		hasuraClaims.AllowedRoles = []string{data.Role}
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"email":                        data.Email,
			"role":                         data.Role,
			"role_id":                      data.RoleId,
			"name":                         data.Name,
			"permissions":                  data.Permissions,
			"groupId":                      data.GroupId,
			"nameGroup":                    data.NameGroup,
			"https://hasura.io/jwt/claims": hasuraClaims,
		},
	)

	tokenString, err := token.SignedString([]byte(config.JWT.SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (data *JwtClaims) ValidateToken(tokenString string) (bool, error) {
	config := GetConfig()
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWT.SECRET), nil
		},
	)
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if err := mapstructure.Decode(claims, &data); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// DecodeToken decodes the JWT token y devuelve los claims
func (data *JwtClaims) DecodeToken(tokenString string) (*JwtClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	decodedClaims := &JwtClaims{}
	if err := mapstructure.Decode(claims, decodedClaims); err != nil {
		return nil, fmt.Errorf("error decoding claims: %v", err)
	}

	return decodedClaims, nil
}
