package api_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/gidx"
	"go.infratographer.com/x/testing/auth"
)

func TestGetVirtualMachineWithAuth(t *testing.T) {
	oauthCLI, issuer, oAuthClose := auth.OAuthTestClient("urn:test:virtualmachine", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	vm1 := (&VirtualMachineBuilder{
		OwnerID: gidx.MustNewID("testown"),
	}).MustNew(ctx)

	resp, err := graphTestClient(
		withHTTPClient(oauthCLI),
		withSrvURL(srv.URL+"/query"),
	).GetVirtualMachineByID(ctx, vm1.ID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, vm1.ID, resp.VirtualMachine.ID)
}

func TestGetVirtualMachinerNoAuth(t *testing.T) {
	_, issuer, oAuthClose := auth.OAuthTestClient("urn:test:virtualmachine", "")
	defer oAuthClose()

	srv, err := newTestServer(
		withAuthConfig(
			&echojwtx.AuthConfig{
				Issuer: issuer,
			},
		),
		withPermissions(
			permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
		),
	)

	require.NoError(t, err)
	require.NotNil(t, srv)

	defer srv.Close()

	ctx := context.Background()
	vm1 := (&VirtualMachineBuilder{
		OwnerID: gidx.MustNewID("testown"),
	}).MustNew(ctx)

	resp, err := graphTestClient(
		withHTTPClient(http.DefaultClient),
		withSrvURL(srv.URL+"/query"),
	).GetVirtualMachineByID(ctx, vm1.ID)

	require.Error(t, err, "Expected an authorization error")
	require.Nil(t, resp)
	assert.ErrorContains(t, err, `{"networkErrors":{"code":401`)
}
