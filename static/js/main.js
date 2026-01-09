document.addEventListener('DOMContentLoaded', function() {
    const searchInput = document.getElementById('searchInput');
    const suggestions = document.getElementById('suggestions');

    if (!searchInput || !suggestions) return;

    searchInput.addEventListener('input', function() {
        const query = this.value.trim().toLowerCase();
        if (query === '') {
            suggestions.classList.remove('active');
            suggestions.innerHTML = '';
            return;
        }
        const results = search(query);
        display(results, query);
    });

    document.addEventListener('click', function(e) {
        if (!searchInput.contains(e.target) && !suggestions.contains(e.target)) {
            suggestions.classList.remove('active');
        }
    });
});

function search(query) {
    const results = [];

    artists.forEach(artist => {
        if (artist.name.toLowerCase().includes(query)) {
            results.push({type: 'artist', label: artist.name, id: artist.id});
        }
    });

    artists.forEach(artist => {
        let members = artist.members;
        if (typeof members === 'string') {
            try { members = JSON.parse(members); } catch(e) { members = []; }
        }
        members.forEach(member => {
            if (member.toLowerCase().includes(query)) {
                results.push({type: 'member', label: member, artistName: artist.name, id: artist.id});
            }
        });
    });

    artists.forEach(artist => {
        if (artist.creationDate.toString().includes(query)) {
            results.push({type: 'creation', label: artist.creationDate.toString(), artistName: artist.name, id: artist.id});
        }
    });

    const seen = new Set();
    const unique = results.filter(result => {
        const key = result.type + '-' + result.label + '-' + result.id;
        if (seen.has(key)) return false;
        seen.add(key);
        return true;
    });

    return unique.slice(0, 10);
}

function display(results, query) {
    const suggestions = document.getElementById('suggestions');

    if (results.length === 0) {
        suggestions.innerHTML = '<div class="suggestion-item" style="cursor: default;">Aucun résultat trouvé</div>';
        suggestions.classList.add('active');
        return;
    }

    let html = '';
    results.forEach(result => {
        let typeName = result.type;
        if (typeName === 'artist') typeName = 'Artiste';
        else if (typeName === 'member') typeName = 'Membre';
        else if (typeName === 'creation') typeName = 'Création';

        let text = result.label;
        if (result.type !== 'artist') {
            text = result.label + ' - ' + result.artistName;
        }

        const regex = new RegExp('(' + query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&') + ')', 'gi');
        text = text.replace(regex, '<strong style="color: #6c63ff;">$1</strong>');

        html += '<div class="suggestion-item" onclick="goToArtist(' + result.id + ')">';
        html += '<span class="suggestion-type">' + typeName + '</span>';
        html += '<span>' + text + '</span>';
        html += '</div>';
    });

    suggestions.innerHTML = html;
    suggestions.classList.add('active');
}

function goToArtist(id) {
    window.location.href = '/artist/' + id;
}
