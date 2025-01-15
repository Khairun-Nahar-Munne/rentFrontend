{{/* views/property/list.tpl */}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.PageTitle}}</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-50">
    <!-- Header -->
    <header class="bg-[#27357e] text-white py-4">
        <div class="container mx-auto px-4">
            <h1 class="text-2xl font-bold">{{.PageTitle}}</h1>
        </div>
    </header>

    <!-- Navigation Breadcrumbs -->
    <div class="bg-white shadow">
        <div class="container mx-auto px-4 py-3">
            <nav class="text-sm">
                {{range $index, $crumb := .Breadcrumbs}}
                    {{if gt $index 0}}
                        <span class="text-gray-500">></span>
                    {{end}}
                    <a href="/{{range $i, $part := $.Breadcrumbs}}{{if le $i $index}}{{if gt $i 0}}/{{end}}{{$part}}{{end}}{{end}}" 
                       class="text-[#27357e] hover:text-blue-700">
                        {{$crumb}}
                    </a>
                {{end}}
            </nav>
        </div>
    </div>

    <!-- Property Grid -->
    <div class="container mx-auto px-4 py-8">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {{if .Properties.Success}}
                {{range $location := .Properties.Locations}}
                    {{range $property := $location.Properties}}
                        <!-- Property Card -->
                        <div class="bg-white rounded-lg shadow overflow-hidden">
                            <!-- Image Container -->
                            <div class="relative h-48">
                                {{if and $property.Images (gt (len $property.Images) 0)}}
                                    <img src="{{index $property.Images 0}}" 
                                         alt="{{$property.HotelName}}"
                                         class="w-full h-full object-cover">
                                {{else}}
                                    <div class="w-full h-full bg-gray-200 flex items-center justify-center">
                                        <span class="text-gray-400">No image available</span>
                                    </div>
                                {{end}}
                                <div class="absolute bottom-4 left-4 bg-white px-3 py-1 rounded">
                                    <span class="font-medium">From {{$property.Price}}</span>
                                </div>
                            </div>

                            <!-- Property Info -->
                            <div class="p-4">
                                <!-- Rating -->
                                <div class="flex items-center justify-between gap-2 text-blue-600 mb-2">
                                <div>
                                    <span class="font-bold">{{$property.Rating}}</span>
                                    <span class="text-sm">({{$property.ReviewCount}} Reviews)</span>
                                </div>
                                    <span class="text-gray-600 text-sm">{{$property.Type}}</span>
                                </div>

                                <!-- Hotel Name -->
                                <a href="/property-details?id={{$property.PropertyID}}" target="_blank" 
                           class="text-gray-600 hover:underline block overflow-hidden text-ellipsis">
                                    {{$property.HotelName}}
                                </a>
                                  
                                <!-- Amenities -->
                                <div class="text-gray-600 text-xs my-2">
                                    {{if $property.Amenities}}
                                        {{range $amenity := $property.Amenities}}
                                            <span class="text-sm text-gray-600 bg-gray-100 py-1 rounded">
                                                {{$amenity}}
                                            </span>
                                        {{end}}
                                    {{end}}
                                </div>
                                <!-- Property Breadcrumbs -->
                                <div class="text-gray-600 mt-2 mb-3">
                                    {{if $property.Breadcrumbs}}
                                        {{range $index, $crumb := $property.Breadcrumbs}}
                                            {{if gt $index 0}}
                                                <span class="text-gray-400">></span>
                                            {{end}}
                                            <a href="/{{range $i, $part := $property.Breadcrumbs}}{{if le $i $index}}{{if gt $i 0}}/{{end}}{{$part}}{{end}}{{end}}"
                                               target="_blank"
                                               class="text-[#27357e] hover:text-blue-700">
                                                {{$crumb}}
                                            </a>
                                        {{end}}
                                    {{end}}
                                </div>
                              

                                <!-- View Button -->
                                <div class="flex items-center justify-between mt-4">
                                    <div>
                                        <img src="https://static.rentbyowner.com/release/28.0.6/static/images/booking.svg" alt="Booking.com">
                                    </div>
                                    <button onclick="window.open('/property-details?id=${property.property_id}', '_blank')" 
                                            class="bg-emerald-500 text-white px-4 py-2 rounded-lg">
                                        View Availability
                                    </button>
                                </div>
                            </div>
                        </div>
                    {{end}}
                {{end}}
            {{else}}
                <div class="col-span-full text-center py-8">
                    <p class="text-gray-600">No properties found for this location.</p>
                </div>
            {{end}}
        </div>
    </div>
</body>
</html>