// Amenity to Font Awesome icon mapping
const amenityIcons = {
  'Non-smoking rooms': 'fa-ban smoking',
  'Smoke-free property': 'fa-ban smoking',
  'Fitness center': 'fa-dumbbell',
  'Room service': 'fa-concierge-bell',
  'Restaurant': 'fa-utensils',
  'Wifi in all areas': 'fa-wifi',
  'Free Wifi': 'fa-wifi',
  'Internet': 'fa-wifi',
  'Pet friendly': 'fa-paw',
  'Facilities for disabled guests': 'fa-wheelchair',
  '1 swimming pool': 'fa-swimming-pool',
  '2 swimming pools': 'fa-swimming-pool',
  'Free parking': 'fa-square-parking',
  'Private Parking': 'fa-square-parking',
  'Parking on site': 'fa-square-parking',
  'Airport shuttle': 'fa-shuttle-van',
  'Airport drop-off': 'fa-plane-departure',
  'Air conditioning': 'fa-snowflake',
  'Heating': 'fa-temperature-high',
  'Daily housekeeping': 'fa-broom',
  'Security alarm': 'fa-shield-alt',
  'Fire extinguishers': 'fa-fire-extinguisher',
  'Smoke alarms': 'fa-bell',
  'Cycling': 'fa-bicycle',
  'Concierge': 'fa-concierge-bell'
};

// Function to get Font Awesome icon
function getAmenityIcon(amenity) {
  const iconClass = amenityIcons[amenity];
  if (iconClass) {
      return `<i class="fas ${iconClass} "></i>`;
  }
  return '<i class="fas fa-circle-check"></i>'; // Default icon for unmapped amenities
}


$(document).ready(function () {
    const urlParams = new URLSearchParams(window.location.search);
    const propertyId = urlParams.get("id");
    if (propertyId) {
      loadPropertyDetails(propertyId);
    }
});
  
function loadPropertyDetails(propertyId) {
    $.ajax({
      url: `http://localhost:8000/v1/property/details?property_id=${propertyId}`,
      method: "GET",
      crossDomain: true, 
      success: function (response) {
        if (response.success) {
          displayPropertyDetails(response.property);
        }
      },
      error: function (error) {
        console.error("Error loading property details:", error);
      },
    });
}
  
function displayPropertyDetails(property) {
    // Update breadcrumbs with clickable links
    document.getElementById("breadcrumbs").innerHTML = property.breadcrumbs
      .map(
        (crumb, index) => `
          <a 
            href="#" 
            class="cursor-pointer hover:text-blue-600" 
            data-level="${index}" 
            data-location="${crumb}"
            onclick="handleBreadcrumbClick(event, '${crumb}', ${index})"
          >
              ${crumb}
          </a>
          ${index < property.breadcrumbs.length - 1 ? " > " : ""}
      `
      )
      .join("");
  
    let roomsInfo = "";
    let bedsInfo = "";
    // Conditionally add bedrooms info
    if (property.bedrooms > 0) {
      roomsInfo += `<span><i class="fa-solid fa-bed mr-1"></i>${property.bedrooms} Bedrooms </span>`;
      bedsInfo += `<span><i class="fa-solid fa-bed mr-1"></i>${property.bedrooms} Bedrooms</span>`;
    }
  
    // Conditionally add bathrooms info
    if (property.bathrooms > 0) {
      roomsInfo += `<span><i class="fa-solid fa-bath mr-1"></i>${property.bathrooms} Bathrooms </span>`;
    }
  
    // Conditionally add guest count info
    if (property.guest_count > 0) {
      roomsInfo += `<span><i class="fa-solid fa-person mr-1"></i>${property.guest_count} Guests </span>`;
    }
  
    const detailsHTML = `
          <div class="grid grid-cols-1 gap-8">
              <div>
              <h1 class="text-2xl font-bold text-blue-900 mb-4">
                      ${bedsInfo} ${property.hotel_name} ${property.city_in_trans}
                  </h1>
                  <div class="md:items-center mb-4">
                      <span class="inline-flex items-center justify-center border-2 text-white bg-blue-900 rounded-full w-7 h-7 ml-1">
                        <i class="fa-solid fa-thumbs-up text-sm"></i>
                      </span>   
                      <span class="text-blue-900 font-bold text-xl">${property.rating}</span>
                      <span class=" text-gray-600 text-sm">(${property.review_count} Reviews)</span>
                      <span class="mx-1">|</span>
                      <span class=" text-gray-600 text-sm">${roomsInfo}</span>
                      <span class="md:ml-2 text-gray-600">
                          ${property.amenities.map(amenity => `
                            <span class="mr-2 text-sm">
                                ${getAmenityIcon(amenity)}
                                ${amenity}
                            </span>
                        `).join(' ')}
                      </span>
                  </div>
                  <div class="md:grid md:grid-cols-2 md:gap-2">
                      <img src="${property.images[0]}" 
                           alt="${property.hotel_name}" 
                           class="md:w-full w-full h-96 object-cover rounded-lg">
                      <div class="grid grid-cols-2 gap-2 mt-2 md:mt-0">
                          ${property.images
                            .slice(1, 5)
                            .map(
                              (img) => `
                              <img src="${img}" 
                                   alt="" 
                                   class="w-full h-[11.75rem] object-cover rounded-lg cursor-pointer"
                                   onclick="updateMainImage(this.src)">
                          `
                            )
                            .join("")}
                      </div>
                  </div>
              </div>
              <div>
                  

                  <div class="mb-6">
                      <h2 class="text-xl font-bold mb-2">Description</h2>
                      <p class="text-gray-600">
                          ${property.description || "No description available"}
                      </p>
                  </div>
                  <div class="mb-6">
                      <h2 class="text-xl font-bold mb-2">Amenities</h2>
                      <div class="grid grid-cols-2 gap-2">
                          ${property.amenities.map(amenity => `
                            <span class=" mr-2">
                                ${getAmenityIcon(amenity)}
                                ${amenity}
                            </span>
                        `).join(' ')}
                      </div>
                  </div>
                  <button class="w-full bg-emerald-500 text-white py-3 rounded-lg text-lg font-bold hover:bg-emerald-600 transition-colors">
                      Check Availability
                  </button>
              </div>
          </div>
      `;
  
    document.getElementById("propertyDetails").innerHTML = detailsHTML;
}

function handleBreadcrumbClick(event) {
    event.preventDefault(); // Prevent default anchor behavior
    
    // Get the base URL dynamically
    const baseUrl = window.location.origin;
    
  
    
    // Construct the final URL
    const newUrl = `${baseUrl}`;
    
    // Navigate to the new URL
    window.location.href = newUrl;
}
  
function updateMainImage(src) {
    const mainImage = document.querySelector(".h-96");
    if (mainImage) {
      mainImage.src = src;
    }
}