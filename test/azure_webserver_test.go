package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
)

const (
	subscriptionID = "ce81f054-340d-4d08-88ed-80f98650eb94"
	labelPrefix    = "nada0038"
	resourceGroup  = labelPrefix + "-A05-RG"
	vmName         = labelPrefix + "A05VM"
	nicName        = labelPrefix + "A05Nic"
	expectedSize   = "Standard_B2s"
)

// TestAzureLinuxVMCreation validates that the VM was deployed successfully.
func TestAzureLinuxVMCreation(t *testing.T) {
	t.Parallel()

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroup, subscriptionID))

	// Confirm VM size
	actualSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroup, subscriptionID)
	assert.Equal(t, expectedSize, string(actualSize))
}

// TestAzureLinuxVMNICConnection confirms the NIC exists and is connected to the VM.
func TestAzureLinuxVMNICConnection(t *testing.T) {
	t.Parallel()

	// Confirm the NIC resource exists
	assert.True(t, azure.NetworkInterfaceExists(t, nicName, resourceGroup, subscriptionID))

	// Confirm the VM has the NIC attached (returns list of NIC IDs)
	vmNics := azure.GetVirtualMachineNics(t, vmName, resourceGroup, subscriptionID)
	assert.NotEmpty(t, vmNics, "VM should have at least 1 NIC attached")
}

// TestAzureLinuxVMUbuntuVersion confirms the VM is running the correct Ubuntu version.
func TestAzureLinuxVMUbuntuVersion(t *testing.T) {
	t.Parallel()

	// GetVirtualMachineImage returns a VMImage struct with Publisher, Offer, Sku, Version fields
	image := azure.GetVirtualMachineImage(t, vmName, resourceGroup, subscriptionID)

	assert.Equal(t, "Canonical", image.Publisher)
	assert.Equal(t, "0001-com-ubuntu-server-jammy", image.Offer)
	assert.Equal(t, "22_04-lts-gen2", image.SKU)
}
