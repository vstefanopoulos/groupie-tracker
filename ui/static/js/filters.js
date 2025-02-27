import { fetchArtists } from "./main.js";
import { updateArtistsContainer } from "./ui.js";
import { updateCreationYearLabel, updateFirstAlbumLabel } from "./sliders.js";

export let filterValues = {
    creationYearMin: 1950,
    creationYearMax: 2020,
    firstAlbumMin: 1950,
    firstAlbumMax: 2020,
    members: null,
    location: "",
};

// Calls api/filters with POST method with filter values
export function handleFilterSubmit(event) {
    event.preventDefault();
    
    filterValues.creationYearMin = Number(document.getElementById('creation-year-slider-min').value);
    filterValues.creationYearMax = Number(document.getElementById('creation-year-slider-max').value);
    filterValues.firstAlbumMin = Number(document.getElementById('first-album-slider-min').value);
    filterValues.firstAlbumMax = Number(document.getElementById('first-album-slider-max').value);
    

    const selectedMember = document.querySelector('input[name="members"]:checked');
    
    filterValues.members = getSelectedMembers()

    filterValues.location = document.getElementById('location-filter').value;
    
        const filters = {
        creation: [
            filterValues.creationYearMin, filterValues.creationYearMax
        ],
        firstAlbum: [
            filterValues.firstAlbumMin, filterValues.firstAlbumMax
        ],
        members: filterValues.members,
        locations: filterValues.location
    }

    console.log(filters)
    
    event.preventDefault();
    fetch("/api/filter", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(filters)
    })
        .then(response => response.json())
        .then(data => updateArtistsContainer(data))
        .catch(error => console.error("Error fetching filtered artists:", error));
}

function getSelectedMembers() {
    return Array.from(document.querySelectorAll('input[name="members"]:checked'))
                .map(checkbox => Number(checkbox.value));
}

export function resetFilters() {
    document.getElementById("filter-form").reset();
    fetchArtists();
    updateCreationYearLabel();
    updateFirstAlbumLabel();

    console.log('Filters reset');
}
