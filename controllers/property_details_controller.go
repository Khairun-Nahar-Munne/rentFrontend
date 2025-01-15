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

    // Create icon mapping
    iconMapping := map[string]string{
        "Non-smoking rooms":                "fa-ban",
        "Smoke-free property":              "fa-ban",
        "Fitness center":                   "fa-dumbbell",
        "Room service":                     "fa-concierge-bell",
        "Restaurant":                       "fa-utensils",
        "Wifi in all areas":                "fa-wifi",
        "Free Wifi":                        "fa-wifi",
        "Internet":                         "fa-wifi",
        "Pet friendly":                     "fa-paw",
        "Facilities for disabled guests":   "fa-wheelchair",
        "1 swimming pool":                  "fa-swimming-pool",
        "2 swimming pools":                 "fa-swimming-pool",
        "Free parking":                     "fa-square-parking",
        "Private Parking":                  "fa-square-parking",
        "Parking on site":                  "fa-square-parking",
        "Airport shuttle":                  "fa-shuttle-van",
        "Airport drop-off":                 "fa-plane-departure",
        "Air conditioning":                 "fa-snowflake",
        "Heating":                          "fa-temperature-high",
        "Daily housekeeping":               "fa-broom",
        "Security alarm":                   "fa-shield-alt",
        "Fire extinguishers":               "fa-fire-extinguisher",
        "Smoke alarms":                     "fa-bell",
        "Cycling":                          "fa-bicycle",
        "Concierge":                        "fa-concierge-bell",
    }

    // Set the data for the template
    c.Data["Property"] = propertyResp.Property
    c.Data["IconMapping"] = iconMapping
    c.Data["Title"] = propertyResp.Property.HotelName
    c.TplName = "property-details.tpl"
}