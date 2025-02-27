const creationYearSliderMin = document.getElementById("creation-year-slider-min");
const creationYearSliderMax = document.getElementById("creation-year-slider-max");
const firstAlbumSliderMin = document.getElementById("first-album-slider-min");
const firstAlbumSliderMax = document.getElementById("first-album-slider-max");

const creationYearValue = document.getElementById("creation-year-value");
const firstAlbumValue = document.getElementById("first-album-value");

export function updateCreationYearLabel() {
    creationYearValue.textContent = `${creationYearSliderMin.value} - ${creationYearSliderMax.value}`;
}

export function updateFirstAlbumLabel() {
    firstAlbumValue.textContent = `${firstAlbumSliderMin.value} - ${firstAlbumSliderMax.value}`;
}

export const setupSliders = () => {
    creationYearSliderMin.addEventListener("input", () => {
        if (parseInt(creationYearSliderMin.value) > parseInt(creationYearSliderMax.value)) {
            creationYearSliderMin.value = creationYearSliderMax.value;
        }
        updateCreationYearLabel();
    });

    creationYearSliderMax.addEventListener("input", () => {
        if (parseInt(creationYearSliderMax.value) < parseInt(creationYearSliderMin.value)) {
            creationYearSliderMax.value = creationYearSliderMin.value;
        }
        updateCreationYearLabel();
    });

    firstAlbumSliderMin.addEventListener("input", () => {
        if (parseInt(firstAlbumSliderMin.value) > parseInt(firstAlbumSliderMax.value)) {
            firstAlbumSliderMin.value = firstAlbumSliderMax.value;
        }
        updateFirstAlbumLabel();
    });

    firstAlbumSliderMax.addEventListener("input", () => {
        if (parseInt(firstAlbumSliderMax.value) < parseInt(firstAlbumSliderMin.value)) {
            firstAlbumSliderMax.value = firstAlbumSliderMin.value;
        }
        updateFirstAlbumLabel();
    });
};
