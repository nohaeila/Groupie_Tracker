document.addEventListener("DOMContentLoaded", function() {
    var locationElements = document.querySelectorAll(".location-container li");
    var map = L.map('map').setView([51.505, -0.09], 3);
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', { maxZoom: 19, attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>' }).addTo(map);
    var concertLocations = [];

    locationElements.forEach(function(locationElement) {
        // Récupération du texte de l'élément et suppression des espaces inutiles
        var locationText = locationElement.textContent.trim();
        concertLocations.push(locationText);

        var location = locationText.split(" le ")[0].replace('_', ' '); // Nom de l'emplacement
        var date = locationText.split(" le ")[1]; // Date du concert

        // Requête à l'API de géocodage
        fetch("https://nominatim.openstreetmap.org/search?q=" + encodeURIComponent(location) + "&format=json&limit=1")
            .then(response => response.json())
            .then(data => {
                if (data.length > 0) {
                    var coordinates = [parseFloat(data[0].lat), parseFloat(data[0].lon)]; // Coordonnées de l'emplacement
                    console.log("Coordonnées de", location, ":", coordinates);

                    // Ajout d'un marqueur à la carte
                    L.marker(coordinates).addTo(map).bindPopup("<b>" + location + "</b><br>Date: " + date);
                } else {
                    console.error("Impossible de trouver les coordonnées pour l'emplacement:", location);
                }
            })
            .catch(error => {
                console.error("Erreur lors de la récupération des coordonnées:", error);
            });
        });

    console.log(concertLocations);
});