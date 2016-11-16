package cards

import "github.com/macnibblet/vcard-api/database"
import "github.com/auth0/go-jwt-middleware"

type Context struct {
	JwtMiddleware  *jwtmiddleware.JWTMiddleware
	UserRepository *database.UserRepository
	CardRepository *database.VCardRepository
}
