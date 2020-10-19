<script>
  import apiService from '../apiService'
  import MovieError from './MovieError.svelte';
  import Loader from './Loader.svelte';

  export let imdbid;

  function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  async function handleSearch(movieId) {
    await timeout(1000);
    return await apiService.getMovie(movieId);
  }
</script>

<div>
    {#await handleSearch(imdbid)}
      <Loader/>
    {:then movie}
      <div>
        <p><strong>Name:</strong> {movie.Name}</p>
        <p><strong>Year:</strong> {movie.Year}</p>
        <p><strong>Score:</strong> {movie.Score}</p>
        <p><strong>Description:</strong> {movie.Description}</p>
      </div>
    {:catch error}
      <MovieError {error}/>
    {/await}
  </div>
