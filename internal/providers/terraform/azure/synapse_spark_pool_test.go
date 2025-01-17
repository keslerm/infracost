package azure_test

import (
	"testing"

	"github.com/infracost/infracost/internal/providers/terraform/tftest"
)

func TestNewAzureRMSynapseSparkPool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTestsWithOpts(t, "synapse_spark_pool_test", &tftest.GoldenFileOptions{
		CaptureLogs: true,
	})
}
