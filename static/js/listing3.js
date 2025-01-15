let currentLocationId = 1; // Default location ID
let selectedLocation = null;
let slideIndexes = {}; 
let xDown = null;
let yDown = null;


$(document).ready(function() {
    loadPropertiesByLocation(currentLocationId);
    
    // Search toggle functionality
    $('#searchToggle').on('click', function() {
        setTimeout(() => {
            $('.search-dropdown').removeClass('transform -translate-y-full');
        }, 100);
        $('.search-dropdown').removeClass('hidden');
        $('.search-dropdown').addClass('active');
        $('.navigation').removeClass('active');
        $('.navigation').addClass('hidden');
        $('#locationSearch').focus();
    });

    // Close search functionality
    $('#closeSearch').on('click', function() {
        $('.search-dropdown').addClass('transform -translate-y-full');
        setTimeout(() => {
            $('.search-dropdown').removeClass('active');
            $('.search-dropdown').addClass('hidden');
            $('.navigation').removeClass('hidden');
            $('.navigation').addClass('active');
        }, 100); // Match this duration with Tailwind's duration-300
    
        hideLocationDropdown();
    });

    // Location search functionality
    $('#locationSearch').on('input', function() {
        const searchQuery = $(this).val().toLowerCase();
        if (searchQuery.length > 1) {
            showLocationDropdown(searchQuery);
        } else {
            hideLocationDropdown();
        }
    });

    // Search button click handler
    $('.container button').on('click', function() {
        if (selectedLocation) {
            loadPropertiesByLocation(selectedLocation.id);
            $('.search-dropdown').removeClass('active');
        }
    });


    // ESC key handler to close search
    $(document).on('keydown', function(e) {
        if (e.key === 'Escape') {
            $('.search-dropdown').removeClass('active');
            hideLocationDropdown();
        }
    });
    $(document).on('touchstart', '.property-slider', handleTouchStart);
    $(document).on('touchmove', '.property-slider', handleTouchMove);
});

function showLocationDropdown(query) {
    $.ajax({
        url: '/api/list/fetch',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const matchedLocations = response.locations.filter(loc => 
                    loc.value.toLowerCase().includes(query.toLowerCase())
                );
                
                displayLocationDropdown(matchedLocations);
            }
        }
    });
}

function displayLocationDropdown(locations) {
    // Create dropdown if it doesn't exist
    let dropdown = $('#locationDropdown');
    if (dropdown.length === 0) {
        dropdown = $('<div id="locationDropdown" class="absolute z-50 w-64 bg-white shadow-lg rounded-lg mt-1"></div>');
        $('#locationSearch').parent().addClass('relative').append(dropdown);
    }

    // Clear and populate dropdown
    dropdown.empty();
    locations.forEach(location => {
        const item = $(`<div class="p-2 hover:bg-gray-100 cursor-pointer">${location.value}</div>`);
        item.on('click', () => {
            selectedLocation = location;
            $('#locationSearch').val(location.value);
            hideLocationDropdown();
        });
        dropdown.append(item);
    });
}

function hideLocationDropdown() {
    $('#locationDropdown').remove();
}

function loadPropertiesByLocation(locationId) {
    $.ajax({
        url: '/api/list/fetch',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const location = response.locations.find(loc => loc.id === locationId);
                if (location) {
                    updateBreadcrumb(location);
                    updatePageTitle(location);
                    displayProperties(location.properties);
                    // Update search input with current location
                    $('#locationSearch').val(location.value);
                    selectedLocation = location;
                }
            }
        },
        error: function(error) {
            console.error('Error loading properties:', error);
        }
    });
}

function updateBreadcrumb(location) {
    if (location.properties && location.properties.length > 0 && location.properties[0].breadcrumbs) {
        const breadcrumbs = location.properties[0].breadcrumbs;
        const lastLocation = breadcrumbs[breadcrumbs.length - 1];

        const breadcrumbHTML = `
            <span class="items-center text-sm">
                <span class="text-blue-900 font-bold">Vacation Rentals in ${lastLocation}</span>
                <span class="mx-2">|</span>
                ${breadcrumbs.map((crumb, index) => `
                    <span class="text-gray-600 cursor-pointer hover:text-blue-600" 
                          onclick="handleBreadcrumbClick('${crumb}', ${index})">
                        ${crumb}
                    </span>
                    ${index < breadcrumbs.length - 1 ? 
                        '<span class="mx-2 text-gray-400"></span>' : ''}
                `).join('>')}
            </span>
        `;
        document.getElementById('locationBreadcrumb').innerHTML = breadcrumbHTML;
        
    } else {
        // Fallback if no breadcrumbs available
        const breadcrumbHTML = `
            <div class="flex items-center text-sm">
                <span class="text-blue-900 font-bold">Vacation Rentals in ${location.value}</span>
                <span class="mx-2">|</span>
                <span class="text-gray-600 cursor-pointer hover:text-blue-600" 
                      onclick="loadPropertiesByLocation(${location.id})">
                    ${location.value}
                </span>
            </div>
        `;
        document.getElementById('locationBreadcrumb').innerHTML = breadcrumbHTML;
        
    }
}

function updatePageTitle(location) {
    const pageTitle = document.getElementById('pageTitle');
    if (location.properties.length > 0 && location.properties[0].breadcrumbs.length > 0) {
        pageTitle.textContent = `Rent By Owner ${location.properties[0].breadcrumbs[0]} - Vacation Rentals in ${location.properties[0].breadcrumbs[0]}`;
    } else {
        pageTitle.textContent = `Rent By Owner - Vacation Rentals`;
    }
}

function handleTouchStart(evt) {
    const firstTouch = evt.touches[0];
    xDown = firstTouch.clientX;
    yDown = firstTouch.clientY;
}

function handleTouchMove(evt) {
    if (!xDown || !yDown) return;

    const xUp = evt.touches[0].clientX;
    const yUp = evt.touches[0].clientY;

    const xDiff = xDown - xUp;
    const yDiff = yDown - yUp;

    if (Math.abs(xDiff) > Math.abs(yDiff)) {
        const slider = $(evt.target).closest('.property-slider');
        const propertyId = slider.data('property-id');
        
        if (xDiff > 0) {
            // Swiped left
            changeSlide(propertyId, 1);
        } else {
            // Swiped right
            changeSlide(propertyId, -1);
        }
    }

    xDown = null;
    yDown = null;
}

function changeSlide(propertyId, direction) {
    const totalSlides = 5; // Assuming each property has 5 images
    if (!slideIndexes[propertyId]) {
        slideIndexes[propertyId] = 0;
    }

    slideIndexes[propertyId] += direction;

    // Handle wrapping
    if (slideIndexes[propertyId] >= totalSlides) {
        slideIndexes[propertyId] = 0;
    }
    if (slideIndexes[propertyId] < 0) {
        slideIndexes[propertyId] = totalSlides - 1;
    }

    updateSlider(propertyId);
}

function updateSlider(propertyId) {
    const slider = $(`.property-slider[data-property-id="${propertyId}"]`);
    const slides = slider.find('.slide');
    const dots = slider.find('.slider-dot');

    // Update slides
    slides.removeClass('opacity-100').addClass('opacity-0');
    $(slides[slideIndexes[propertyId]]).removeClass('opacity-0').addClass('opacity-100');

    // Update dots
    dots.removeClass('bg-white').addClass('bg-gray-300');
    $(dots[slideIndexes[propertyId]]).removeClass('bg-gray-300').addClass('bg-white');
}

function displayProperties(properties) {
    const propertyGrid = document.getElementById('propertyGrid');
    propertyGrid.innerHTML = '';

    properties.forEach(property => {
        // Initialize slide index for this property
        slideIndexes[property.property_id] = 0;

        const imageSlider = `
            <div class="property-slider relative h-64 overflow-hidden" data-property-id="${property.property_id}">
                ${property.images.map((img, index) => `
                    <div class="slide absolute w-full h-full transition-opacity duration-300 ease-in-out ${index === 0 ? 'opacity-100' : 'opacity-0'}"
                         style="transition-property: opacity">
                        <img src="${img }" 
                             alt="${property.hotel_name}" 
                             class="w-full h-full object-cover">
                    </div>
                `).join('')}
                
                <button onclick="changeSlide('${property.property_id}', -1)" 
                        class="absolute left-2 top-1/2 -translate-y-1/2 bg-white rounded-full w-7 h-7 shadow-md z-10 opacity-80 hover:opacity-100 transition-opacity duration-200">
                    <i class="fas fa-chevron-left"></i>
                </button>
                <button onclick="changeSlide('${property.property_id}', 1)" 
                        class="absolute right-2 top-1/2 -translate-y-1/2 bg-white rounded-full w-7 h-7  shadow-md z-10 opacity-80 hover:opacity-100 transition-opacity duration-200">
                    <i class="fas fa-chevron-right"></i>
                </button>
                
                <div class="absolute bottom-4 left-1/2 -translate-x-1/2 flex space-x-2">
                    ${Array(5).fill(0).map((_, index) => `
                        <div class="slider-dot w-2 h-2 rounded-full ${index === 0 ? 'bg-white' : 'bg-gray-300'} 
                             cursor-pointer transition-colors duration-200"
                             onclick="slideIndexes['${property.property_id}']=${index};updateSlider('${property.property_id}')">
                        </div>
                    `).join('')}
                </div>
            </div>
        `;

        const propertyCard = `
            <div class="bg-white rounded-lg shadow-md overflow-hidden">
                
                <div class="relative">
                    ${imageSlider}
                    <div class="absolute text-xs bottom-4 left-4 bg-white text-black-200 font-bold px-2 py-1 rounded-lg shadow-md flex items-center">
                        From ${property.price}
                        <span class="inline-flex items-center justify-center border-2 border-black rounded-full p-2 w-3 h-3 ml-1">
                            <i class="fa-solid fa-info text-xs text-black"></i>
                        </span>
                    </div>
                </div>
                <div class="p-4">
                    <div class="flex items-center justify-between mb-2 text-sm">
                        <div>
                            <span class="inline-flex items-center justify-center border-2 text-white bg-blue-900 rounded-full w-6 h-6 ml-1">
                               <i class="fa-solid fa-thumbs-up text-xs"></i>
                            </span>
                            <span class="text-blue-900 font-bold text-sm">${property.rating}</span>
                            <span class="text-gray-500 text-xs">(${property.review_count} Reviews)</span>
                        </div>
                        <span class="ml-2 text-gray-600 text-xs">${property.type}</span>
                    </div>
                    <h3 class="text-md font-bold mb-2 overflow-hidden whitespace-nowrap">
                        <a href="/property-details?id=${property.property_id}" target="_blank" 
                           class="text-gray-600 hover:underline block overflow-hidden text-ellipsis">
                            ${property.hotel_name}
                        </a>
                    </h3>
                    <div class="text-gray-600 text-xs my-2">
                        ${property.amenities.join('<i class="fas fa-circle text-[8px] mx-2 align-middle"></i>')}
                    </div>
                    <div class="text-blue-900 mb-2 text-xs">
                        ${property.breadcrumbs.map((crumb, index) => `
                            <span class="cursor-pointer hover:text-blue-600 text-sm" 
                                  onclick="handleBreadcrumbClick('${crumb}', ${index})">
                                ${crumb}
                            </span>
                            ${index < property.breadcrumbs.length - 1 ? ' > ' : ''}
                        `).join('')}
                    </div>
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
        `;
        propertyGrid.innerHTML += propertyCard;
    });
}

function handleBreadcrumbClick(locationName, level) {
    $.ajax({
        url: '/api/list/fetch',
        method: 'GET',
        success: function(response) {
            if (response.success) {
                const location = response.locations.find(loc => 
                    loc.value.toLowerCase() === locationName.toLowerCase()
                );
                if (location) {
                    currentLocationId = location.id;
                    loadPropertiesByLocation(location.id);
                }
            }
        }
    });
}