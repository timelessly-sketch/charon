package public

import "charon/internal/model/entity"

type KubernetesService struct {
	ServiceBase       entity.ServiceBase
	ServiceContainers entity.ServiceContainers
}
