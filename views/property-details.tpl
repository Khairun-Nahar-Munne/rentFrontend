{{/* views/property_details.tpl */}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Property.HotelName}}</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <!-- Add Font Awesome -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css">
</head>
<body class="bg-gray-100">
    <!-- Breadcrumb Navigation -->
    <div class="bg-white shadow">
        <div class="contain mx-auto px-4 py-3">
            <nav class="text-sm">
                {{range $index, $crumb := .Property.Breadcrumbs}}
                    {{if gt $index 0}}
                        <span class="text-gray-500">></span>
                    {{end}}
                    <a href="/{{range $i, $part := $.Property.Breadcrumbs}}{{if le $i $index}}{{if gt $i 0}}/{{end}}{{$part}}{{end}}{{end}}" 
                       class="text-[#27357e] hover:text-blue-700">
                        {{$crumb}}
                    </a>
                {{end}}
            </nav>
        </div>
    </div>

    <div class="container mx-auto px-4 py-8">
        <!-- Hotel Name, Rating, and Details -->
        <div class="mb-6">
         <!-- First line -->
            <div class="text-blue-900 text-2xl">
                <span class=" font-bold mb-2">{{.Property.HotelName}}</span> 
                <span class="mx-1 mb-2">|</span>
                  {{if .Property.Type}}
                        <span class="font-bold mb-2 mr-1">{{.Property.Type}}</span>
                {{end}}
                    {{if .Property.CityInTrans}}
                            <span class="font-bold mb-2">{{.Property.CityInTrans}}</span>
                    {{end}}
            </div>
            <!-- Second line -->
            <div class="flex items-center mt-4 mb-4">
                    <span class="inline-flex items-center justify-center border-2 text-white bg-blue-900 rounded-full w-7 h-7 ml-1">
                            <i class="fa-solid fa-thumbs-up text-sm"></i>
                    </span> 
                    <span class="text-gray-600 px-2 py-1 rounded">{{.Property.Rating}}</span>
                    <span class=" text-gray-600">({{.Property.ReviewCount}} Reviews)</span>
                    <span class="mx-2 mb-2 text-xl">|</span>
                    <!-- Property Details -->
                    <div class="flex flex-wrap gap-6 text-gray-700">
                        <div class="flex flex-wrap gap-4">
                            {{range .Property.Amenities}}
                                {{$iconClass := index $.IconMapping .}}
                                <span class="rounded flex items-center gap-1">
                                    <i class="fas {{$iconClass}}"></i>
                                    <span>{{.}}</span>
                                </span>
                            {{end}}
                        </div>
                        {{if ne .Property.Bedrooms 0}}
                            <div class="flex items-center">
                                <i class="fa-solid fa-bed mr-1"></i>
                                <span>{{.Property.Bedrooms}} Bedroom{{if ne .Property.Bedrooms 1}}s{{end}}</span>
                            </div>
                        {{end}}

                        {{if ne .Property.Bathrooms 0}}
                            <div class="flex items-center">
                                <i class="fa-solid fa-bath mr-1"></i>
                                <span>{{.Property.Bathrooms}} Bathroom{{if ne .Property.Bathrooms 1}}s{{end}}</span>
                            </div>
                        {{end}}

                        {{if ne .Property.GuestCount 0}}
                            <div class="flex items-center">
                                <i class="fa-solid fa-users mr-1"></i>
                                <span>{{.Property.GuestCount}} Guest{{if ne .Property.GuestCount 1}}s{{end}}</span>
                            </div>
                        {{end}}
                    </div>
            </div>

           
        </div>

        <!-- Image Gallery -->
        <div class="md:grid md:grid-cols-2 md:gap-2">
            {{range $index, $image := .Property.Images}}
                {{if eq $index 0}}
                    <img 
                        src="{{$image}}" 
                        alt="Property Image {{$index}}" 
                        class="md:w-full w-full h-96 object-cover rounded-lg">
                {{end}}
            {{end}}

            <!-- Right grid of 4 smaller images -->
            <div class="grid grid-cols-2 gap-2 mt-2 md:mt-0">
                {{range $index, $image := .Property.Images}}
                    {{if gt $index 0}}
                        <img 
                            src="{{$image}}" 
                            alt="Property Image {{$index}}" 
                            class="w-full h-[11.75rem] object-cover rounded-lg">
                    {{end}}
                {{end}}
            </div>
        </div>

        <!-- Description -->
        <div class="mt-8  mb-8">
            <div class="flex font-bold text-blue-900 text-2xl mb-2">
                {{if ne .Property.Bedrooms 0}}
                    <div class="mr-1">
                        <span>{{.Property.Bedrooms}} Bedroom{{if ne .Property.Bedrooms 1}}s{{end}}</span>
                    </div>
                {{end}}
                  {{if .Property.Type}}
                    <div class="mr-1">
                        <span>{{.Property.Type}}</span>
                    </div>
                {{end}}
                {{if .Property.CityInTrans}}
                    <span class="font-bold">{{.Property.CityInTrans}}</span>
                {{end}}
            </div>
            <p class="text-gray-700">{{.Property.Description}}</p>
            <button class="mt-4 w-full bg-green-500 text-white px-8 py-3 rounded-lg hover:bg-green-600 transition-colors">
                Check Availability
            </button>
        </div>
    </div>
</body>
</html>