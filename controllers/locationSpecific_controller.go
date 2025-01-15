package controllers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
    "sync"
    "rentFrontend/models"

    beego "github.com/beego/beego/v2/server/web"
)

type PropertySpecificController struct {
    beego.Controller
}

func (c *PropertySpecificController) Get() {
    // Extract the location path from the URL
    fullPath := c.Ctx.Request.URL.Path
    // Remove the base path "/api/list/fetch/" to get just the location parts
    locationPath := strings.TrimPrefix(fullPath, "/api/list/fetch/")
    // Split the path into location components
    locationParts := strings.Split(strings.Trim(locationPath, "/"), "/")

    apiURL := "http://localhost:8000/v1/property/list"
    
    // Create channels for communication
    responseChan := make(chan models.PropertyResponse)
    errorChan := make(chan error)
    
    var wg sync.WaitGroup
    wg.Add(1)
    
    // Fetch and filter data asynchronously
    go func() {
        defer wg.Done()
        
        // Fetch data from API
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

        // Filter locations based on the location path
        filteredLocations := filterLocationsByPath(propertyResp.Locations, locationParts)
        propertyResp.Locations = filteredLocations
        propertyResp.Success = len(filteredLocations) > 0
        
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
        c.Data["json"] = map[string]interface{}{
            "success": response.Success,
            "data":    response,
        }
    }
    
    c.ServeJSON()
}

// filterLocationsByPath filters locations and properties based on the provided location path
func filterLocationsByPath(locations []models.Location, locationPath []string) []models.Location {
    if len(locationPath) == 0 {
        return locations
    }

    filteredLocations := make([]models.Location, 0)
    for _, location := range locations {
        filteredProperties := make([]models.Property, 0)
        
        for _, property := range location.Properties {
            if matchesLocationPath(property, locationPath) {
                filteredProperties = append(filteredProperties, property)
            }
        }
        
        if len(filteredProperties) > 0 {
            locationCopy := location
            locationCopy.Properties = filteredProperties
            filteredLocations = append(filteredLocations, locationCopy)
        }
    }
    
    return filteredLocations
}

// matchesLocationPath checks if a property matches the given location path
func matchesLocationPath(property models.Property, locationPath []string) bool {
    if len(locationPath) == 0 || len(property.Breadcrumbs) == 0 {
        return false
    }

    // For single-level path (e.g., "united-states-of-america"), 
    // match if it's anywhere in the breadcrumbs
    if len(locationPath) == 1 {
        searchTerm := normalizePathComponent(locationPath[0])
        for _, crumb := range property.Breadcrumbs {
            if normalizePathComponent(crumb) == searchTerm {
                return true
            }
        }
        return false
    }

    // For multi-level paths, ensure they match in sequence
    for i, pathComponent := range locationPath {
        if i >= len(property.Breadcrumbs) {
            return false
        }
        if normalizePathComponent(pathComponent) != normalizePathComponent(property.Breadcrumbs[i]) {
            return false
        }
    }
    
    return true
}

// normalizePathComponent normalizes a path component for comparison
func normalizePathComponent(component string) string {
    // Convert to lowercase and replace spaces with hyphens
    normalized := strings.ToLower(component)
    normalized = strings.ReplaceAll(normalized, " ", "-")
    return normalized
}