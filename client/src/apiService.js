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

async function getMovies(page) {
    if (page === undefined) {
        return
    }
    let url = `http://localhost:8080/movies?page=${page}`;
    let res = await fetch(url, {
        method: "get",
        mode: 'cors',
        credentials: 'same-origin'
    });
    let body = await res.json();
    return await body;
}

async function movieAmount(){
    let url = `http://localhost:8080/movieamount`;
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
    getMovies,
    movieAmount
}
