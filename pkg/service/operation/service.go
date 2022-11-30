package operation

import (
	"fmt"
	sdkutil "github.com/TBD54566975/ssi-sdk/util"
	"github.com/pkg/errors"
	"github.com/tbd54566975/ssi-service/internal/util"
	"github.com/tbd54566975/ssi-service/pkg/service/framework"
	opstorage "github.com/tbd54566975/ssi-service/pkg/service/operation/storage"
	"github.com/tbd54566975/ssi-service/pkg/storage"
)

type Service struct {
	storage opstorage.Storage
}

func (s Service) Type() framework.Type {
	return framework.Operation
}

func (s Service) Status() framework.Status {
	ae := sdkutil.NewAppendError()
	if s.storage == nil {
		ae.AppendString("no storage configured")
	}
	if !ae.IsEmpty() {
		return framework.Status{
			Status:  framework.StatusNotReady,
			Message: fmt.Sprintf("operation service is not ready: %s", ae.Error().Error()),
		}
	}
	return framework.Status{Status: framework.StatusReady}
}

func NewOperationService(s storage.ServiceStorage) (*Service, error) {
	opStorage, err := opstorage.NewOperationStorage(s)
	if err != nil {
		return nil, util.LoggingErrorMsg(err, "creating operation storage")
	}
	service := &Service{storage: opStorage}
	if !service.Status().IsReady() {
		return nil, errors.New(service.Status().Message)
	}
	return service, nil
}