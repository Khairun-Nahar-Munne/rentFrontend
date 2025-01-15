package controllers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    beego "github.com/beego/beego/v2/server/web"
    "rentFrontend/models"
)

type PropertyController struct {
    beego.Controller
}

func (c *PropertyController) Get() {
    // Get property_id from query parameters
    propertyId := c.GetString("id")
    if propertyId == "" {
        c.Redirect("/", 302)
        return
    }

    // Fetch property details from your API
    apiURL := "http://localhost:8080/api/details/fetch?property_id=" + propertyId
    resp, err := http.Get(apiURL)
    if err != nil {
        c.Data["Error"] = "Failed to fetch property details"
        c.TplName = "error.tpl"
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.Data["Error"] = "Failed to read response"
        c.TplName = "error.tpl"
        return
    }

    var propertyResp models.PropertyDetailsResponse
    err = json.Unmarshal(body, &propertyResp)
    if err != nil {
        c.Data["Error"] = "Failed to parse property details"
        c.TplName = "error.tpl"
        return
    }

    // Format breadcrumbs for better display
    for i, crumb := range propertyResp.Property.Breadcrumbs {
        propertyResp.Property.Breadcrumbs[i] = formatBreadcrumb(crumb)
    }

    // Set the data for the template
    c.Data["Property"] = propertyResp.Property
    c.Data["Title"] = propertyResp.Property.HotelName
    c.TplName = "property-details.tpl"
}

