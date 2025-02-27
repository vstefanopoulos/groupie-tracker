let artistsData = null;

import { updateArtistsContainer } from "./ui.js";
import { handleSearchInput, handleSearchSubmit, resetSearch } from "./search.js";
import { resetFilters, handleFilterSubmit } from "./filters.js";
import { setupSliders } from "./sliders.js"; 

document.addEventListener("DOMContentLoaded", () => {
    fetchArtists();
    
    document.getElementById("search-form").addEventListener("submit", handleSearchSubmit);
    document.getElementById("reset-search").addEventListener("click", resetSearch);   
    document.getElementById("filter-form").addEventListener("submit", handleFilterSubmit);
    document.getElementById("reset-filters").addEventListener("click", resetFilters);

    setupSliders();

    const inputElement = document.getElementById("search-input");
    inputElement.addEventListener('keyup', (event) => handleSearchInput(event, artistsData));
});

// Calls api/artists and updates the artist-container
export function fetchArtists() {
    if (artistsData) {
        updateArtistsContainer(artistsData);
        return;
    }

    fetch("/api/artists")
        .then(response => response.json())
        .then(data => {
            artistsData = data; 
            updateArtistsContainer(data);
        })
        .catch(error => console.error("Error fetching artists:", error));
}



