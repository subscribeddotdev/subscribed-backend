package jwks

import (
	"crypto/rsa"
	_ "embed"
	"testing"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	jwkclient "github.com/lestrrat-go/jwx/jwk"
	"github.com/stretchr/testify/require"
)

//go:embed publickey.test.json
var JwksPublicKey string

//go:embed privatekey.test.json
var JwksPrivateKey string

func JwtGenerator(t *testing.T, loginProviderID string) string {
	t.Helper()

	claims := map[string]any{
		"sid": "sess_123",
		"sub": loginProviderID,
		"iss": "https://clerk.com",
	}

	parsedPrivateKey, err := jwkclient.ParseKey([]byte(JwksPrivateKey))
	require.NoError(t, err)

	rsaPrivateKeY := &rsa.PrivateKey{}
	err = parsedPrivateKey.Raw(rsaPrivateKeY)
	require.NoError(t, err)

	signerOpts := &jose.SignerOptions{}
	signerOpts.WithType("JWT")
	signerOpts.WithHeader("kid", "kid-emulators")
	signer, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: jose.RS256,
			Key:       rsaPrivateKeY,
		},
		signerOpts,
	)
	require.NoError(t, err)

	builder := jwt.Signed(signer)
	builder = builder.Claims(claims)
	token, err := builder.CompactSerialize()
	require.NoError(t, err)

	return token
}
