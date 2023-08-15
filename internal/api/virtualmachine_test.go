package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
)

func TestQuery_IPAddress(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vm1 := (&VirtualMachineBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.VirtualMachine
		errorMsg string
	}{
		{
			name:     "happy path virtual machine",
			queryID:  vm1.ID,
			expected: vm1,
		},
		{
			name:     "not found id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "virtual_machine not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetVirtualMachineByID(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.VirtualMachine)
			require.NotNil(t, resp.VirtualMachine.Location)
			require.NotNil(t, resp.VirtualMachine.Owner)

			assert.Equal(t, tc.expected.ID, resp.VirtualMachine.ID)
			assert.Equal(t, tc.expected.Name, resp.VirtualMachine.Name)
			assert.Equal(t, tc.expected.OwnerID, resp.VirtualMachine.Owner.ID)
			assert.Equal(t, tc.expected.LocationID, resp.VirtualMachine.Location.ID)
		})
	}
}

// TODO: test crud operations
