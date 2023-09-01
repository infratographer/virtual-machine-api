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

func TestQuery_CPUConfig(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vmCPUConf := (&VirtualMachineCPUConfigBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.VirtualMachineCPUConfig
		errorMsg string
	}{
		{
			name:     "happy path virtual machine cpu config",
			queryID:  vmCPUConf.ID,
			expected: vmCPUConf,
		},
		{
			name:     "not found id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "virtual_machine_cpu_config not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetVirtualMachineCPUConfigByID(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.VirtualMachineCPUConfig)
			require.NotNil(t, resp.VirtualMachineCPUConfig.Cores)
			require.NotNil(t, resp.VirtualMachineCPUConfig.Sockets)

			assert.Equal(t, tc.expected.ID, resp.VirtualMachineCPUConfig.ID)
			assert.Equal(t, tc.expected.Cores, resp.VirtualMachineCPUConfig.Cores)
			assert.Equal(t, tc.expected.Sockets, resp.VirtualMachineCPUConfig.Sockets)
		})
	}
}

// TODO: test crud operations
