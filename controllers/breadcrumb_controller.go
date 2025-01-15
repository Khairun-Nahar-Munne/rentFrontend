package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentFrontend/models"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyViewController struct {
    beego.Controller
}

func (c *PropertyViewController) Get() {
    // Get the path and decode it
    path := c.Ctx.Request.URL.Path
    path = strings.TrimPrefix(path, "/")
    
    // Skip if path is empty (home page)
    if path == "" {
        c.Data["PageTitle"] = "Vacation Rentals"
        c.Data["Breadcrumbs"] = []string{"Home"}
        c.Data["Properties"] = models.PropertyResponse{Success: true}
        c.TplName = "category.tpl"
        return
    }
    
    // Construct the API URL with the correct prefix
    apiURL := fmt.Sprintf("http://localhost:8080/api/list/fetch/%s", path)
    
    // Add debug logging
    fmt.Printf("Fetching from API URL: %s\n", apiURL)

    // Make API request
    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Printf("Error fetching API: %v\n", err)
        c.handleError("Failed to fetch properties")
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response: %v\n", err)
        c.handleError("Failed to read response")
        return
    }


    var apiResp map[string]interface{}
    if err := json.Unmarshal(body, &apiResp); err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        c.handleError("Failed to parse response")
        return
    }

    // Extract the data part from the response
    var propertyResp models.PropertyResponse
    if data, ok := apiResp["data"].(map[string]interface{}); ok {
        // Marshal and unmarshal to ensure proper type conversion
        dataBytes, err := json.Marshal(data)
        if err != nil {
            c.handleError("Failed to process response data")
            return
        }
        
        if err := json.Unmarshal(dataBytes, &propertyResp); err != nil {
            c.handleError("Failed to process property data")
            return
        }
        
        propertyResp.Success = true
    } else {
        propertyResp.Success = false
    }

    // Create breadcrumbs from the path
    breadcrumbs := strings.Split(path, "/")
    // Capitalize and format breadcrumbs
    for i, crumb := range breadcrumbs {
        breadcrumbs[i] = formatBreadcrumb(crumb)
    }
    
    // Set page title using the last breadcrumb
    pageTitle := "Vacation Rentals"
    if len(breadcrumbs) > 0 {
        pageTitle = fmt.Sprintf("Vacation Rentals in %s", breadcrumbs[len(breadcrumbs)-1])
    }

    c.Data["PageTitle"] = pageTitle
    c.Data["Breadcrumbs"] = breadcrumbs
    c.Data["Properties"] = propertyResp
    // In your Go code where you set up the template
    c.TplName = "category.tpl"
}

func (c *PropertyViewController) handleError(message string) {
    c.Data["Error"] = message
    c.Data["Properties"] = models.PropertyResponse{Success: false}
    c.TplName = "category.tpl"
}

func formatBreadcrumb(crumb string) string {
    // Replace hyphens with spaces and capitalize words
    words := strings.Split(strings.ReplaceAll(crumb, "-", " "), " ")
    for i, word := range words {
        words[i] = strings.Title(strings.ToLower(word))
    }
    return strings.Join(words, " ")
}