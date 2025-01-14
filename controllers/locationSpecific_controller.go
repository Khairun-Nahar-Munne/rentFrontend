// controllers/property_client_controller.go
package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rentFrontend/models"
	"strings"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertySpecificController struct {
    beego.Controller
}



func (c *PropertySpecificController) Get() {
    // Get the full path and split it into breadcrumb components
    path := c.Ctx.Input.Param(":splat")
    breadcrumbs := strings.Split(path, "/")

    apiURL := "http://localhost:8000/v1/property/list"
    
    // Create channels for communication
    responseChan := make(chan models.PropertyResponse)
    errorChan := make(chan error)
    
    var wg sync.WaitGroup
    wg.Add(1)
    
    // Fetch data asynchronously
    go func() {
        defer wg.Done()
        
        resp, err := http.Get(apiURL)
        if err != nil {
            errorChan <- err
            return
        }
        defer resp.Body.Close()
        
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            errorChan <- err
            return
        }
        
        var propertyResp models.PropertyResponse
        err = json.Unmarshal(body, &propertyResp)
        if err != nil {
            errorChan <- err
            return
        }
        
        // Filter locations based on breadcrumbs
        filteredLocations := make([]models.Location, 0)
        for _, location := range propertyResp.Locations {
            filteredProperties := make([]models.Property, 0)
            
            for _, property := range location.Properties {
                // Check if property breadcrumbs match the requested path
                matches := true
                for i, crumb := range breadcrumbs {
                    if i >= len(property.Breadcrumbs) || !strings.EqualFold(property.Breadcrumbs[i], crumb) {
                        matches = false
                        break
                    }
                }
                
                if matches {
                    filteredProperties = append(filteredProperties, property)
                }
            }
            
            if len(filteredProperties) > 0 {
                location.Properties = filteredProperties
                filteredLocations = append(filteredLocations, location)
            }
        }
        
        propertyResp.Locations = filteredLocations
        responseChan <- propertyResp
    }()
    
    // Close channels after WaitGroup is done
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
        if len(response.Locations) == 0 {
            c.Data["json"] = map[string]interface{}{
                "success": false,
                "error":   "No properties found for the specified location",
            }
        } else {
            c.Data["json"] = response
        }
    }
    
    c.ServeJSON()
}