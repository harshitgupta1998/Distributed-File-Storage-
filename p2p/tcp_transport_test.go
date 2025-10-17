package p2p
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransportOptions(t *testing.T) {
	listenAddr := "localhost:8080"
	tr := NewTCPTransportOptions(listenAddr)
	assert.Equal(t, listenAddr, tr.listenAddress)

	tr.ListenAndAccept()
}

