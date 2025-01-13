// controllers/property_client_controller.go
package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rentFrontend/models"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyClientController struct {
    beego.Controller
}



// api/list/fetch
func (c *PropertyClientController) Get() {
    apiURL := "http://localhost:8000/v1/property/list" // Replace with your actual API URL
    
    // Create a channel to receive the API response
    responseChan := make(chan models.PropertyResponse)
    errorChan := make(chan error)
    
    // Create a WaitGroup to wait for the goroutine to complete
    var wg sync.WaitGroup
    wg.Add(1)
    
    // Fetch data from the API asynchronously
    go func() {
        defer wg.Done()
        
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
        
        // Parse JSON response
        var propertyResp models.PropertyResponse
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