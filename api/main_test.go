package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/pzanwar/employee/db/sqlc"
	"github.com/pzanwar/employee/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
