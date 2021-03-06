package storage

import (
	"github.com/kyma-project/kyma/components/remote-environment-broker/internal"
)

// RemoteEnvironment is an interface that describe storage layer operations for Charts
type RemoteEnvironment interface {
	Upsert(re *internal.RemoteEnvironment) (bool, error)
	Get(name internal.RemoteEnvironmentName) (*internal.RemoteEnvironment, error)
	FindAll() ([]*internal.RemoteEnvironment, error)
	FindOneByServiceID(id internal.RemoteServiceID) (*internal.RemoteEnvironment, error)
	Remove(name internal.RemoteEnvironmentName) error
}

// Instance is an interface that describe storage layer operations for Instances
type Instance interface {
	Insert(i *internal.Instance) error
	Remove(id internal.InstanceID) error
	Get(id internal.InstanceID) (*internal.Instance, error)
	FindOne(func(i *internal.Instance) bool) (*internal.Instance, error)
	UpdateState(iID internal.InstanceID, state internal.InstanceState) error
}

// InstanceOperation is an interface that describe storage layer operations for InstanceOperations
type InstanceOperation interface {
	// Insert is inserting object into storage.
	// Object is modified by setting CreatedAt.
	Insert(*internal.InstanceOperation) error
	Get(internal.InstanceID, internal.OperationID) (*internal.InstanceOperation, error)
	GetAll(internal.InstanceID) ([]*internal.InstanceOperation, error)
	UpdateState(internal.InstanceID, internal.OperationID, internal.OperationState) error
	UpdateStateDesc(internal.InstanceID, internal.OperationID, internal.OperationState, *string) error
	Remove(internal.InstanceID, internal.OperationID) error
}

// IsNotFoundError checks if given error is NotFound error
func IsNotFoundError(err error) bool {
	nfe, ok := err.(interface {
		NotFound() bool
	})
	return ok && nfe.NotFound()
}

// IsAlreadyExistsError checks if given error is AlreadyExist error
func IsAlreadyExistsError(err error) bool {
	aee, ok := err.(interface {
		AlreadyExists() bool
	})
	return ok && aee.AlreadyExists()
}
