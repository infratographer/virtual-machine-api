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

func TestQuery_MemoryConfig(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vmMemoryConfig := (&VirtualMachineMemoryConfigBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.VirtualMachineMemoryConfig
		errorMsg string
	}{
		{
			name:     "happy path virtual machine memory config",
			queryID:  vmMemoryConfig.ID,
			expected: vmMemoryConfig,
		},
		{
			name:     "not found id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "virtual_machine_memory_config not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetVirtualMachineMemoryConfigByID(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.VirtualMachineMemoryConfig)
			require.NotNil(t, resp.VirtualMachineMemoryConfig.Size)

			assert.Equal(t, tc.expected.ID, resp.VirtualMachineMemoryConfig.ID)
			assert.Equal(t, tc.expected.Size, int(resp.VirtualMachineMemoryConfig.Size))
		})
	}
}

// TODO: test crud operations
