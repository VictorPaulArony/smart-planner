  // Initialize the map
  const map = L.map("map").setView([-0.0917, 34.763], 10); // Center on Kisumu

  // Add OpenStreetMap tile layer
  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    maxZoom: 19,
    attribution: "© OpenStreetMap contributors",
  }).addTo(map);

  // Fetch dynamic data from the backend
  let data = {};
  fetch("https://joabowala.github.io/Data/data.json")
    .then((response) => response.json())
    .then((jsonData) => {
      data = jsonData;
      console.log("Data loaded successfully:", data);
    })
    .catch((err) => {
      console.error("Failed to fetch data:", err);
    });

  // DOM elements
  const countySelect = document.getElementById("county-search");
  const subcountySelect = document.getElementById("subcounty-search");
  const wardSelect = document.getElementById("ward-search");
  const resourceSelect = document.getElementById("resource-search");
  const searchButton = document.getElementById("search-button");

  // Event listener for county selection
  countySelect.addEventListener("change", () => {
    console.log("Selected county:", countySelect.value);

    subcountySelect.innerHTML =
      '<option value="">Select a sub-county...</option>';
    wardSelect.innerHTML = '<option value="">Select a ward...</option>';
    resourceSelect.value = "";

    if (countySelect.value && data[countySelect.value]) {
      const subcounties = Object.keys(data[countySelect.value]);
      console.log("Available sub-counties:", subcounties);

      subcounties.forEach((subcounty) => {
        const option = document.createElement("option");
        option.value = subcounty;
        option.textContent = subcounty;
        subcountySelect.appendChild(option);
      });
      subcountySelect.disabled = false;
    } else {
      console.error(
        "No data found for selected county or invalid county selection."
      );
      subcountySelect.disabled = true;
      wardSelect.disabled = true;
      resourceSelect.disabled = true;
      searchButton.disabled = true;
    }
  });

  // Populate ward dropdown
  subcountySelect.addEventListener("change", () => {
    console.log("Selected sub-county:", subcountySelect.value);

    wardSelect.innerHTML = '<option value="">Select a ward...</option>';
    resourceSelect.value = "";

    if (
      subcountySelect.value &&
      data[countySelect.value][subcountySelect.value]
    ) {
      const wards = Object.keys(
        data[countySelect.value][subcountySelect.value]
      );
      console.log("Available wards:", wards);

      wards.forEach((ward) => {
        const option = document.createElement("option");
        option.value = ward;
        option.textContent = ward;
        wardSelect.appendChild(option);
      });
      wardSelect.disabled = false;
    } else {
      console.error(
        "No data found for selected sub-county or invalid sub-county selection."
      );
      wardSelect.disabled = true;
      resourceSelect.disabled = true;
      searchButton.disabled = true;
    }
  });

  // Enable resource selection and search button
  wardSelect.addEventListener("change", () => {
    console.log("Selected ward:", wardSelect.value);

    if (wardSelect.value) {
      resourceSelect.disabled = false;
      searchButton.disabled = false;
    } else {
      resourceSelect.disabled = true;
      searchButton.disabled = true;
    }
  });

  // Search functionality
  searchButton.addEventListener("click", () => {
    const county = countySelect.value;
    const subcounty = subcountySelect.value;
    const ward = wardSelect.value;
    const resource = resourceSelect.value;

    console.log(
      `Filters - County: ${county}, Sub-County: ${subcounty}, Ward: ${ward}, Resource: ${resource}`
    );

    // Clear existing markers
    map.eachLayer((layer) => {
      if (layer instanceof L.Marker || layer instanceof L.CircleMarker) {
        map.removeLayer(layer);
      }
    });

    // Add markers
    if (
      data[county] &&
      data[county][subcounty] &&
      data[county][subcounty][ward] &&
      data[county][subcounty][ward][resource]
    ) {
      const locations = data[county][subcounty][ward][resource];
      console.log("Locations to display:", locations);

      locations.forEach((location) => {
        L.marker(location.coords)
          .addTo(map)
          .bindPopup(`<strong>${location.title}</strong>`);
      });

      // Zoom to the first location
      if (locations.length > 0) {
        map.setView(locations[0].coords, 15);
      }
    } else {
      console.error("No locations found for the selected filters.");
    }
  });