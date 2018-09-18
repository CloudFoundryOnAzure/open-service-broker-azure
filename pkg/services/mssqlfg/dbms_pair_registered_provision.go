package mssqlfg

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (d *dbmsPairRegisteredManager) GetProvisioner(
	service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("preProvision", d.preProvision),
		service.NewProvisioningStep("validatePriServer", d.validatePriServer),
		service.NewProvisioningStep("validateSecServer", d.validateSecServer),
		service.NewProvisioningStep("testPriConnection", d.testPriConnection),
		service.NewProvisioningStep("testSecConnection", d.testSecConnection),
	)
}

func (d *dbmsPairRegisteredManager) preProvision(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pp := instance.ProvisioningParameters
	priServerName := pp.GetString("primaryServer")
	secServerName := pp.GetString("secondaryServer")
	if priServerName == secServerName {
		return nil, fmt.Errorf("The primary server and the secondary server " +
			"should be different servers")
	}
	priLocation := pp.GetString("primaryLocation")
	secLocation := pp.GetString("secondaryLocation")
	if priLocation == secLocation {
		return nil, fmt.Errorf("The primary server and the secondary server " +
			"should be in different locations")
	}
	return &dbmsPairInstanceDetails{
		PriServerName:                 priServerName,
		PriAdministratorLogin:         pp.GetString("primaryAdministratorLogin"),
		PriAdministratorLoginPassword: service.SecureString(pp.GetString("primaryAdministratorLoginPassword")), // nolint: lll
		SecServerName:                 secServerName,
		SecAdministratorLogin:         pp.GetString("secondaryAdministratorLogin"),
		SecAdministratorLoginPassword: service.SecureString(pp.GetString("secondaryAdministratorLoginPassword")), // nolint: lll
	}, nil
}

func (d *dbmsPairRegisteredManager) validatePriServer(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pp := instance.ProvisioningParameters
	dt := instance.Details.(*dbmsPairInstanceDetails)
	fqdn, err := validateServer(
		ctx,
		&d.serversClient,
		pp.GetString("primaryResourceGroup"),
		dt.PriServerName,
		instance.Service.GetProperties().Extended["version"].(string),
		pp.GetString("primaryLocation"),
	)
	if err != nil {
		return nil, err
	}
	dt.PriFullyQualifiedDomainName = fqdn
	return dt, nil
}

func (d *dbmsPairRegisteredManager) validateSecServer(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pp := instance.ProvisioningParameters
	dt := instance.Details.(*dbmsPairInstanceDetails)
	fqdn, err := validateServer(
		ctx,
		&d.serversClient,
		pp.GetString("secondaryResourceGroup"),
		dt.SecServerName,
		instance.Service.GetProperties().Extended["version"].(string),
		pp.GetString("secondaryLocation"),
	)
	if err != nil {
		return nil, err
	}
	dt.SecFullyQualifiedDomainName = fqdn
	return dt, nil
}

func (d *dbmsPairRegisteredManager) testPriConnection(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*dbmsPairInstanceDetails)
	if err := testConnection(
		dt.PriFullyQualifiedDomainName,
		dt.PriAdministratorLogin,
		string(dt.PriAdministratorLoginPassword),
	); err != nil {
		return nil, err
	}
	return instance.Details, nil
}

func (d *dbmsPairRegisteredManager) testSecConnection(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*dbmsPairInstanceDetails)
	if err := testConnection(
		dt.SecFullyQualifiedDomainName,
		dt.SecAdministratorLogin,
		string(dt.SecAdministratorLoginPassword),
	); err != nil {
		return nil, err
	}
	return instance.Details, nil
}
