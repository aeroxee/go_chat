package auth

//var secretKey = []byte("")
//
//type Credentials struct {
//	UserID uint `json:"user_id"`
//}
//
//type Claims struct {
//	Credentials
//	jwt.RegisteredClaims
//}
//
//func GetToken(credentials Credentials) (string, error) {
//	expirationTime := time.Now().Add(time.Hour * 24)
//	claims := &Claims{
//		Credentials: credentials,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(expirationTime),
//		},
//	}
//
//	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := jwtToken.SignedString(secretKey)
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}
//
//func VerifyToken(token string) (Claims, error) {
//	var claims Claims
//
//	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
//		return secretKey, nil
//	})
//	if err != nil {
//		return Claims{}, err
//	}
//
//	if !jwtToken.Valid {
//		return Claims{}, errors.New("this token is not valid")
//	}
//
//	return claims, nil
//}