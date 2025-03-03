import { updateArtistsContainerFromSearch } from "./ui.js";
import { fetchArtists } from "./main.js";

let debounceTimer;

// Calls getSuggestions with 'keyup'. Waits 500ms 
export function handleSearchInput(event, artistsData) {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {    
        getSuggestions(event.target.value, artistsData);
    }, 500);
}

// Calls for api/search clears and updates artists-container without reloading the page
export function handleSearchSubmit(event) {
    const suggestionsElement = document.getElementById("suggestions");
    suggestionsElement.innerHTML = ""; 

    event.preventDefault();
    const query = document.getElementById("search-input").value.trim();
    if (query) fetchSearchResults(query);
}

// Resets the artists-container with all artists and cleans suggestions and search query
export function resetSearch() {
    const suggestionsElement = document.getElementById("suggestions");
    suggestionsElement.innerHTML = ""; 
    document.getElementById("search-form").reset();
    fetchArtists();
}

// Searches through artistData and returns first 10 results with field found
// Uses a Set instead of array to exclude duplicate results
function getSuggestions(query, artistsData) {
    const suggestionsElement = document.getElementById("suggestions");
    suggestionsElement.innerHTML = "";
    if (query === "") {
        return
    }
    
    query = query.toLowerCase();

    const matchWithFields = new Set();

    const foundMatch = artistsData.filter(artist => {
        let match = false;

        if (artist.name.toLowerCase().includes(query)) {
            match = true;
            matchWithFields.add("artist " + artist.name);
        }

        artist.members.forEach(member => {
            if (member.toLowerCase().includes(query)) {
                match = true;
                matchWithFields.add("members " + member);
            }
        });

        if (artist.creationDate.toString().includes(query)) {
            match = true;
            matchWithFields.add("creation " + artist.creationDate);
        }

        if (artist.firstAlbum.includes(query)) {
            match = true;
            matchWithFields.add("first " + artist.firstAlbum);
        }

        Object.keys(artist.Relation.datesLocations).forEach(location => {  
            if (location.toLowerCase().includes(query)) {
                match = true;
                matchWithFields.add("locations " + location);
            }
        });
        return match
    }).slice(0, 10);

    if (!foundMatch) {
        matchWithFields.add("No matches found")
    }

    // debug
    console.log(matchWithFields)
    
    showSuggestions(matchWithFields)
}

// Iterates through suggestions Set and updates the container
function showSuggestions(matchWithFields) {
    const inputElement = document.getElementById("search-input");
    const suggestionsElement = document.getElementById("suggestions");
    if (matchWithFields.size > 0) {
        suggestionsElement.style.display = "block";
        matchWithFields.forEach(suggestion => {
            const suggestionDiv = document.createElement("div");
            suggestionDiv.classList.add("suggestion-item");
    
            suggestionDiv.innerHTML = `${suggestion}`;
            suggestionDiv.addEventListener("click", function () {
                fetchSearchResults(suggestion);
                inputElement.value = suggestion;  
                suggestionsElement.innerHTML = "";
                suggestionsElement.style.display = "none";
            });
    
            suggestionsElement.appendChild(suggestionDiv);
        });
    }
    else {
        suggestionsElement.style.display = "none";
    }
}

// Calls api/search with query and updates the artist-container
function fetchSearchResults(query) {
    fetch(`/api/search?q=${query}`)
        .then(response => response.json())
        .then(data => updateArtistsContainerFromSearch(data))
        .catch(error => console.error("Error fetching search results:", error));
}

