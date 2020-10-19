async function getMovie(movie) {
    let url = `http://localhost:8080/movies/${movie}`;
    console.log(url);
    let res = await fetch(url, {
        method: "get",
        mode: 'cors',
        credentials: 'same-origin'
    });
    let body = await res.json();
    console.log(body)
    return await body;
}

module.exports = {
    getMovie
}
