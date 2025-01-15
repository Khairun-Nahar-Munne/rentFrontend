// controllers/property_details_client_controller.go
package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentFrontend/models"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsClientController struct {
    beego.Controller
}


// api/details/fetch
func (c *PropertyDetailsClientController) Get() {
    // Get property_id from query parameters
    propertyId := c.GetString("property_id")
    if propertyId == "" {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "error":   "property_id is required",
        }
        c.ServeJSON()
        return
    }

    // Create channels for response and error
    responseChan := make(chan models.PropertyDetailsResponse)
    errorChan := make(chan error)

    var wg sync.WaitGroup
    wg.Add(1)

    // Fetch data from the API asynchronously
    go func() {
        defer wg.Done()

        // Construct the API URL with the property_id
        apiURL := "http://localhost:8000/v1/property/details?property_id=" + propertyId

        // Make HTTP GET request
        resp, err := http.Get(apiURL)
        if err != nil {
            errorChan <- err
            return
        }
        defer resp.Body.Close()

        // Read response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            errorChan <- err
            return
        }

        // Check if the response status is not 200
        if resp.StatusCode != http.StatusOK {
            errorChan <- fmt.Errorf("API returned status code: %d", resp.StatusCode)
            return
        }

        // Parse JSON response
        var propertyResp models.PropertyDetailsResponse
        err = json.Unmarshal(body, &propertyResp)
        if err != nil {
            errorChan <- err
            return
        }

        responseChan <- propertyResp
    }()

    // Create a goroutine to close channels after WaitGroup is done
    go func() {
        wg.Wait()
        close(responseChan)
        close(errorChan)
    }()

    // Handle response or error
    select {
    case err := <-errorChan:
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "error":   err.Error(),
        }
    case response := <-responseChan:
        c.Data["json"] = response
    }

    c.ServeJSON()
}