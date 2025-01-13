<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Property Listings</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script src="/static/js/listing.js" defer></script>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css">
    <style>
        .search-dropdown {
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            transform: translateY(80%);
            transition: transform 0.5s ease-in;
            z-index: 50;
        }
        
        .search-dropdown.active {
            transform: translateY(0);
        }

        /* Add this to ensure nav stays on top */
        nav {
            position: relative;
            z-index: 51;
        }
    </style>
</head>
<body class="bg-gray-100">
    <!-- Navigation -->
    <nav class="bg-blue-900 p-4 relative">
        <div class="container mx-auto flex items-center justify-between">
            <div>
                <img src="https://static.rentbyowner.com/release/28.0.6/static/images/sites/rentbyowner.com/header_logo.svg" alt="rentbyowner logo">
            </div>
            <div class="navigation active flex items-center space-x-8">
                <button id="searchToggle" class="text-white hover:text-gray-200">
                    <i class="fas fa-search text-xl"></i>
                </button>
                <button class="text-white hover:text-gray-200">
                    <i class="fas fa-user text-xl"></i>
                </button>
            </div>
        </div>
    </nav>

    <!-- Search Dropdown Panel -->
    <div class="search-dropdown hidden absolute w-full bg-blue-900 p-4 shadow-lg">
        <div class="container mx-auto">
            <!-- Add flex justify-end to align the inner flex container to the right -->
            <div class="flex justify-center">
                <!-- Your existing flex container with a specific width -->
                <div class="flex items-center space-x-4 w-2/3">
                    <div class="flex-grow relative">
                        <input type="text" id="locationSearch" placeholder="United States of America" 
                            class="w-full p-2 rounded-lg">
                    </div>
                    <button class="bg-emerald-500 text-white px-6 py-2 rounded-lg">Search</button>
                    <button id="closeSearch" class="text-white hover:text-gray-200">Close</button>
                </div>
            </div>
        </div>
    </div>
    <!-- Main Content -->
    <div class="container mx-auto mt-8 p-4">
        <!-- Breadcrumb -->
        <div id="locationBreadcrumb" class="text-sm text-gray-600 mb-4"></div>
        
        <h1 id="pageTitle" class="text-3xl font-bold text-blue-900 mb-6">Rent By Owner</h1>
        

        <!-- Property Grid -->
        <div id="propertyGrid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <!-- Properties will be loaded here via AJAX -->
        </div>
    </div>
</body>
</html>