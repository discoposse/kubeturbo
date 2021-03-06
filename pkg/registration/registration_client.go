package registration

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	TargetIdentifierField string = "targetIdentifier"
	Username              string = "username"
	Password              string = "password"
)

type K8sRegistrationClient struct{}

func NewK8sRegistrationClient() *K8sRegistrationClient {
	return &K8sRegistrationClient{}
}

func (rClient *K8sRegistrationClient) GetSupplyChainDefinition() []*proto.TemplateDTO {
	supplyChainFactory := NewSupplyChainFactory()
	supplyChain, err := supplyChainFactory.createSupplyChain()
	if err != nil {
		// TODO error handling
	}
	return supplyChain
}

func (rClient *K8sRegistrationClient) GetAccountDefinition() []*proto.AccountDefEntry {
	var acctDefProps []*proto.AccountDefEntry

	// target ID
	targetIDAcctDefEntry := builder.NewAccountDefEntryBuilder(TargetIdentifierField, "Address",
		"IP of the Kubernetes master", ".*", false, false).Create()
	acctDefProps = append(acctDefProps, targetIDAcctDefEntry)

	// username
	usernameAcctDefEntry := builder.NewAccountDefEntryBuilder(Username, "Username",
		"Username of the Kubernetes master", ".*", false, false).Create()
	acctDefProps = append(acctDefProps, usernameAcctDefEntry)

	// password
	passwordAcctDefEntry := builder.NewAccountDefEntryBuilder(Password, "Password",
		"Password of the Kubernetes master", ".*", false, true).Create()
	acctDefProps = append(acctDefProps, passwordAcctDefEntry)

	return acctDefProps
}

func (rClient *K8sRegistrationClient) GetIdentifyingFields() string {
	return TargetIdentifierField
}
