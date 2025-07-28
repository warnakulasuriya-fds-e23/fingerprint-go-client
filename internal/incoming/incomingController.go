package incoming

import (
	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-interface-service/internal/outgoing"
)

type IncomingController struct {
	outgoingClient outgoing.OutgoingClient
}

// TODO: Remove mutex locks
func NewIncomingController(outgoingClient outgoing.OutgoingClient) *IncomingController {
	return &IncomingController{outgoingClient: outgoingClient}
}
func (controller *IncomingController) FileSubmitHandler(c *gin.Context) {
	controller.fileSubmitHandler(c)
}
