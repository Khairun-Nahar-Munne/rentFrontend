{{/* views/property_details.tpl */}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Property.HotelName}}</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/tailwindcss/2.2.19/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100">
    <!-- Breadcrumb Navigation -->
    <div class="bg-white shadow">
        <div class="container mx-auto px-4 py-3">
            <nav class="text-sm">
                {{range $index, $crumb := .Property.Breadcrumbs}}
                    {{if gt $index 0}}
                        <span class="mx-2 text-gray-500">/</span>
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
        <!-- Hotel Name and Rating -->
        <div class="mb-6">
            <h1 class="text-3xl font-bold mb-2">{{.Property.HotelName}}</h1>
            <div class="flex items-center">
                <span class="bg-blue-700 text-white px-2 py-1 rounded">{{.Property.Rating}}</span>
                <span class="ml-2 text-gray-600">({{.Property.ReviewCount}} Reviews)</span>
            </div>
        </div>

        <!-- Image Gallery -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mb-8">
            {{range $index, $image := .Property.Images}}
                <div class="{{if eq $index 0}}col-span-2 row-span-2{{end}}">
                    <img src="{{$image}}" alt="Property Image" class="w-full h-full object-cover rounded">
                </div>
            {{end}}
        </div>

        <!-- Amenities -->
        <div class="mb-8">
            <h2 class="text-2xl font-bold mb-4">Amenities</h2>
            <div class="flex flex-wrap gap-4">
                {{range .Property.Amenities}}
                    <span class="bg-gray-200 px-4 py-2 rounded">{{.}}</span>
                {{end}}
            </div>
        </div>

        <!-- Description -->
        <div class="mb-8">
            <h2 class="text-2xl font-bold mb-4">Description</h2>
            <p class="text-gray-700 whitespace-pre-line">{{.Property.Description}}</p>
            <button class="bg-green-500 text-white px-8 py-3 rounded-lg">
                Check Availability
            </button>
        </div>
    </div>
</body>
</html>