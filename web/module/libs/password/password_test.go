package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	assert.Equal(t, "810130b55d71763f2c7fef7402a698b9b2a53617ebf136f4c6083cb301722571", Sha256("kovey"))
}

func TestPassword(t *testing.T) {
	assert.Equal(t, "a50d7108a6f65cb8a4f919fd359f489b3755d09d9c46c82d243159d0ed28efaa", Password("kovey", "123456"))
	t.Logf("%s", Password("administrator", Sha256("123456")))
}

func TestVerify(t *testing.T) {
	assert.True(t, Verify("kovey", "123456", "a50d7108a6f65cb8a4f919fd359f489b3755d09d9c46c82d243159d0ed28efaa"))
}
