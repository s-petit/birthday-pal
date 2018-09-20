package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_basicauth_authenticated_client(t *testing.T) {

	basic := BasicAuth{Username: "user", Password: "pass"}

	_, err := basic.Client()

	assert.NoError(t, err)
}
