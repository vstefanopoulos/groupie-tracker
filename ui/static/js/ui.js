// Updates the artist-container with data
export function updateArtistsContainer(artists) {
    const container = document.getElementById("artists-container");
    container.innerHTML = "";

    if (artists.length === 0) {
        container.innerHTML = '<p>No artists found.</p>';
        return;
    }

    artists.forEach(artist => {
        const artistDiv = document.createElement("div");
        artistDiv.classList.add("artist-box");

        artistDiv.innerHTML = `
            <a href="details?id=${artist.id}">
                <img src="${artist.image}" alt="${artist.name}" loading="lazy">
            </a>
            <h2>${artist.name}</h2>
            <a href="details?id=${artist.id}">View Details</a>
        `;

        container.appendChild(artistDiv);
    });
}

export function updateArtistsContainerFromSearch(search_results) {
    const container = document.getElementById("artists-container");
    container.innerHTML = "";

    if (search_results.length === 0) {
        container.innerHTML = '<p>No artists found.</p>';
        return;
    }

    search_results.forEach(result => {
        const artistDiv = document.createElement("div");
        artistDiv.classList.add("artist-box");

        artistDiv.innerHTML = `
            <a href="details?id=${result.Artist.id}">
                <img src="${result.Artist.image}" alt="${result.Artist.name}" loading="lazy">
            </a>
            <h2>${result.Artist.name}</h2>
            <a href="details?id=${result.Artist.id}">View Details</a>
        `;

        container.appendChild(artistDiv);
    });
}
