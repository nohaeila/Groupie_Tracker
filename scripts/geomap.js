document.addEventListener("DOMContentLoaded", function() {
    // Récupère tous les éléments de la classe "location-container"
    var locationElements = document.querySelectorAll(".location-container li");

    // Crée une carte Leaflet
    var map = L.map('map').setView([51.505, -0.09], 3); // Réglez la vue initiale de la carte selon vos besoins

    // Ajoute une couche de tuiles OpenStreetMap à la carte
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    // Tableau pour stocker les lieux des concerts
    var concertLocations = [];

    // Parcours tous les éléments et récupère leur contenu textuel
    locationElements.forEach(function(locationElement) {
        var locationText = locationElement.textContent.trim();
        concertLocations.push(locationText);

        var location = locationText.split(" le ")[0].replace('_', ' '); // Extrait le nom de l'emplacement et remplace les underscores par des espaces
        var date = locationText.split(" le ")[1]; // Extrait la date du concert

        // Utilise l'API de géocodage pour obtenir les coordonnées de l'emplacement
        fetch("https://nominatim.openstreetmap.org/search?q=" + encodeURIComponent(location) + "&format=json&limit=1")
            .then(response => response.json())
            .then(data => {
                if (data.length > 0) {
                    var coordinates = [parseFloat(data[0].lat), parseFloat(data[0].lon)]; // Coordonnées de l'emplacement

                    console.log("Coordonnées de", location, ":", coordinates);

                    // Ajoute un marqueur à la carte pour chaque emplacement de concert
                    L.marker(coordinates).addTo(map)
                        .bindPopup("<b>" + location + "</b><br>Date: " + date);
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
