package cloudgenetics

import (
	//"encoding/json"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	auth0audience string
	auth0domain   string
)

// Jwks is JSON Web slices
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys Key values
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// Router returns a gin HTTP engine
func Router() *gin.Engine {
	// Initialize Auth0 variables
	setAuth0Variables()
	// Creates a router without any middleware by default
	r := gin.New()

	// Release mode
	gin.SetMode(gin.ReleaseMode)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// CORS middleware
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	return r
}

func setAuth0Variables() {
	auth0audience = os.Getenv("AUTH0_AUDIENCE")
	auth0domain = os.Getenv("AUTH0_DOMAIN")
}

// JSON Web Token middleware
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(auth0audience, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

// ValidateRequest will verify that a token received from an http request
// is valid and signed by Auth0 and the scope is valid
func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// CustomClaims Define JWT Scope
type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func checkScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	claims, ok := token.Claims.(*CustomClaims)

	hasScope := false
	if ok && token.Valid {
		result := strings.Split(claims.Scope, " ")
		for i := range result {
			if result[i] == scope {
				hasScope = true
			}
		}
	}

	return hasScope
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + auth0domain + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

// PublicRoutes define available public routes
func PublicRoutes(r *gin.Engine) {
	unauthorized := r.Group("/")
	unauthorized.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Cloudgenetics API")
	})
}

// APIV1Routes define API v1 private routes
func APIV1Routes(r *gin.Engine) {
	// Create an authorized group for API V1
	authorized := r.Group("/api/v1")
	// Info on version 1 of API
	authorized.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Cloudgenetics API Version 1")
	})

	// Anything below info should require authentication
	authorized.Use(checkJWT())
	// Get all projects
	authorized.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Cloudgenetics Protected API Version 1")
	})
}
