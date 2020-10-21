async function getMovie(movie) {
    let url = `http://localhost:8080/movies/${movie}`;
    let res = await fetch(url, {
        method: "get",
        mode: 'cors',
        credentials: 'same-origin'
    });
    let body = await res.json();
    return await body;
}

async function getMovies() {
    let url = `http://localhost:8080/movies?wow=wow`;
    let res = await fetch(url, {
        method: "get",
        mode: 'cors',
        credentials: 'same-origin'
    });
    let body = await res.json();
    return await body;
}

module.exports = {
    getMovie,
    getMovies
}
