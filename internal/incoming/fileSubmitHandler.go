package incoming

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-client/internal/responseobjects"
)

func (controller *IncomingController) fileSubmitHandler(c *gin.Context) {
	file, errLoadFileFromForm := c.FormFile("file")
	if errLoadFileFromForm != nil {
		resObj := responseobjects.ErrorResObj{Message: "error  while loading file from form data field 'file' " + errLoadFileFromForm.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	if filepath.Ext(file.Filename) != ".png" {
		resObj := responseobjects.ErrorResObj{Message: "error  only png files can be uploaded to this endpoint "}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}

	currentDirectory, errGetWorkingDir := os.Getwd()
	if errGetWorkingDir != nil {
		resObj := responseobjects.ErrorResObj{Message: "error  unable to get working directory of biometric-interface-service : " + errGetWorkingDir.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	imageFilePath := filepath.Join(currentDirectory, "/tmp", file.Filename)
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, imageFilePath)
	status, errOutgoing := controller.outgoingClient.AuthorizeTemplateFilesMethod(imageFilePath)
	if errOutgoing != nil {
		resObj := responseobjects.ErrorResObj{Message: "error  sending outgoing : " + errOutgoing.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	c.IndentedJSON(http.StatusOK, map[string]string{"status": status})

}
