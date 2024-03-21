function searchArtist() {
    var input = document.getElementById('search-bar').value.toLowerCase();
    var artistes = document.querySelectorAll('.artiste');

    artistes.forEach(function (artiste) {
        var artistName = artiste.getAttribute('data-artist').toLowerCase();
        if (artistName.includes(input)) {
            artiste.style.display = '';
        } else {
            artiste.style.display = 'none';
        }
    });
}