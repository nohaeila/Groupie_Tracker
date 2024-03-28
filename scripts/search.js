// Fonction pour effectuer une recherche
function search() {
    // Récupération de la valeur de la barre de recherche et conversion en minuscules
    var input = document.getElementById('search-bar').value.toLowerCase();

    // Récupération de tous les éléments avec la classe 'artiste'
    var artistes = document.querySelectorAll('.artiste');

    // Pour chaque artiste...
    artistes.forEach(function (artiste) {
        // Récupération du nom de l'artiste et conversion en minuscules
        var artistName = artiste.getAttribute('data-artist').toLowerCase();

        // Si le nom de l'artiste contient la valeur de recherche...
        if (artistName.includes(input)) {
            // ... l'artiste est affiché
            artiste.style.display = '';
        } else {
            // ... sinon, l'artiste est caché
            artiste.style.display = 'none';
        }
    });
}