// // platform/authenticator/auth.go

package auth

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	db ports.CustomerRepoPort
}

// CustomToken extends oauth2.Token to include the ID token.
type CustomToken struct {
	*oauth2.Token
	IDToken string `json:"id_token"`
}

// New instantiates the *Authenticator.
func New(db ports.CustomerRepoPort) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		os.Getenv("AUTH0_URL"),
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENTID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		//RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint: provider.Endpoint(),
		Scopes:   []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		db:       db,
	}, nil
}

// VerifyIDToken verifies that an *CustomToken is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token string) (*oidc.IDToken, error) {
	rawIDToken := token

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Authenticator) IsAuthenticated(c *fiber.Ctx) error {
	ctx := c.Context() // Get the Go standard library context from Fiber's context

	// Get Token
	var token string
	startsWith := "Bearer"
	authHeader := c.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, startsWith) {
		// Split the Authorization Into Array
		token = strings.Fields(authHeader)[1]

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": 401,
			"message":     "Login To Continue",
			"Error":       "Unauthorized",
		})
	}

	// Verify Token
	idToken, err := a.VerifyIDToken(ctx, token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "Failedd To Verify Token",
			"Error":       err.Error(),
		})
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	//Check if user is present in DB
	//Save customer to db

	//Get Email FRom Locals
	emailInterface, ok := profile["email"]
	if !ok {
		fmt.Println("Error getting email from profile")
		return c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "internal error",
			"error":       ok,
		})
	}

	// Now, type assert emailInterface to get the email string
	email, ok := emailInterface.(string)
	if !ok {
		fmt.Println("Error converting email to string")
		return c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "internal error",
			"error":       "email type assertion failed",
		})
	}

	//Get Email FRom Locals
	nameInterface, ok := profile["name"]
	if !ok {
		fmt.Println("Error getting name from profile")
		return c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "internal error",
			"error":       ok,
		})
	}

	// Now, type assert emailInterface to get the email string
	name, ok := nameInterface.(string)
	if !ok {
		fmt.Println("Error converting name to string")
		return c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "internal error",
			"error":       "email type assertion failed",
		})
	}

	customer, err := a.db.GetCustomerByEmail(email)
	if err != nil {
		fmt.Println("Error", err)

		//Save Customer To DB
		cus, err := a.db.CreateCustomer(queries.CreateCustomerParams{
			CustomerID: domain.GenerateUUID(),
			Name:       name,
			Email:      email})

		if err != nil {
			fmt.Println("Error", err)
		}

		c.Locals("customer", cus)

	} else {
		c.Locals("customer", customer)
	}

	//Store In Locals
	c.Locals("user", profile)

	// Next Middleware/Route
	return c.Next()
}

func (a *Authenticator) AllowedRoles(allowedrole string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")

		// Get role claim from user
		rolesClaim, ok := user.(map[string]interface{})["http://localhost:3000/roles"]
		if !ok {
			fmt.Println("Roles claim not found or not a string slice")
			return c.Status(500).JSON(fiber.Map{
				"status_code": 500,
				"message":     "internal error",
				"error":       ok,
			})
		}

		// Convert rolesClaim to a slice of strings
		roles, ok := rolesClaim.([]interface{})
		if !ok {
			fmt.Println("Roles claim is not a slice of interfaces")
			return c.Status(500).JSON(fiber.Map{
				"status_code": 500,
				"message":     "internal error",
				"error":       ok,
			})
		}

		// Convert []interface{} to []string
		rolesStrings := make([]string, len(roles))
		for i, v := range roles {
			// Convert each interface{} to string
			rolesStrings[i] = v.(string)
		}

		fmt.Println(rolesStrings)

		// Check if the user's role is in the allowedRoles slice.
		allowed := false

		for _, role := range rolesStrings {
			if role == allowedrole {
				allowed = true
				break
			}
		}

		// If the user's role is not in the allowedRoles slice, return a Forbidden response.
		if !allowed {
			return c.Status(403).JSON(fiber.Map{
				"status_code": 403,
				"message":     "Forbidden!.Login with proper rights",
			})
		}

		return c.Next()
	}

}
