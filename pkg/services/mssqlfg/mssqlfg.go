package mssqlfg

import (
	sqlSDK "github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2017-03-01-preview/sql" // nolint: lll
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/open-service-broker-azure/pkg/azure/arm"
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

type module struct {
	dbmsPairRegisteredManager     *dbmsPairRegisteredManager
	databasePairManager           *databasePairManager
	databasePairRegisteredManager *databasePairRegisteredManager
	databasePairFePrimaryManager  *databasePairFePrimaryManager
	databasePairFeManager         *databasePairFeManager
}

type dbmsPairRegisteredManager struct {
	sqlDatabaseDNSSuffix string
	armDeployer          arm.Deployer
	serversClient        sqlSDK.ServersClient
}

type databasePairManager struct {
	armDeployer          arm.Deployer
	databasesClient      sqlSDK.DatabasesClient
	failoverGroupsClient sqlSDK.FailoverGroupsClient
}

type databasePairRegisteredManager struct {
	armDeployer          arm.Deployer
	databasesClient      sqlSDK.DatabasesClient
	failoverGroupsClient sqlSDK.FailoverGroupsClient
}

type databasePairFePrimaryManager struct {
	armDeployer          arm.Deployer
	databasesClient      sqlSDK.DatabasesClient
	failoverGroupsClient sqlSDK.FailoverGroupsClient
}

type databasePairFeManager struct {
	armDeployer          arm.Deployer
	databasesClient      sqlSDK.DatabasesClient
	failoverGroupsClient sqlSDK.FailoverGroupsClient
}

// New returns a new instance of a type that fulfills the service.Module
// interface and is capable of provisioning MS SQL servers and databases
// using "Azure SQL Database"
func New(
	azureEnvironment azure.Environment,
	armDeployer arm.Deployer,
	serversClient sqlSDK.ServersClient,
	databasesClient sqlSDK.DatabasesClient,
	failoverGroupsClient sqlSDK.FailoverGroupsClient,
) service.Module {
	return &module{
		dbmsPairRegisteredManager: &dbmsPairRegisteredManager{
			sqlDatabaseDNSSuffix: azureEnvironment.SQLDatabaseDNSSuffix,
			armDeployer:          armDeployer,
			serversClient:        serversClient,
		},
		databasePairManager: &databasePairManager{
			armDeployer:          armDeployer,
			databasesClient:      databasesClient,
			failoverGroupsClient: failoverGroupsClient,
		},
		databasePairRegisteredManager: &databasePairRegisteredManager{
			armDeployer:          armDeployer,
			databasesClient:      databasesClient,
			failoverGroupsClient: failoverGroupsClient,
		},
		databasePairFePrimaryManager: &databasePairFePrimaryManager{
			armDeployer:          armDeployer,
			databasesClient:      databasesClient,
			failoverGroupsClient: failoverGroupsClient,
		},
		databasePairFeManager: &databasePairFeManager{
			armDeployer:          armDeployer,
			databasesClient:      databasesClient,
			failoverGroupsClient: failoverGroupsClient,
		},
	}
}

func (m *module) GetName() string {
	return "mssqlfg"
}