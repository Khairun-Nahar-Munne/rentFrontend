package models

type PropertyResponse struct {
    Success   bool       `json:"success"`
    Locations []Location `json:"locations"`
}

type Location struct {
    ID         int         `json:"id"`
    DestID     string      `json:"dest_id"`
    DestType   string      `json:"dest_type"`
    Value      string      `json:"value"`
    Properties []Property  `json:"properties"`
}

type Property struct {
    ID              int      `json:"id"`
    PropertyID      int      `json:"property_id"`
    PropertySlugID  string   `json:"property_slug_id"`
    CityInTrans     string  `json:"city_in_trans"`
    HotelName       string   `json:"hotel_name"`
    Bedrooms        int      `json:"bedrooms"`
    Bathrooms       int      `json:"bathrooms"`
    GuestCount      int      `json:"guest_count"`
    Rating          float64  `json:"rating"`
    ReviewCount     int      `json:"review_count"`
    Price           string   `json:"price"`
    Breadcrumbs     []string `json:"breadcrumbs"`
    DisplayLocation []string `json:"display_location"`
    Amenities       []string `json:"amenities"`
    Type           string   `json:"type"`
    Images         []string `json:"images"`
}