async function getMovie(movie) {
    // console.log(movie);
    let res = await fetch('https://swapi.dev/api/people/1')
    let body = await res.json()
    return await body;
}

module.exports = {
    getMovie
}
