<script>
  import apiService from '../apiService'
  import Movie from './Movie.svelte'
  import MovieError from './MovieError.svelte';
  import Loader from './Loader.svelte';

  export let newResult;
  export let imdbid;

  let search = false;

  $: imdbid, search = imdbid != undefined ? handleSearch(imdbid) : false;

  async function handleSearch(movieId) {
    let results = await apiService.getMovie(movieId);
    newResult = true;
    return await results;
  }
</script>

<div class="container">
  {#await search}
    <Loader/>
  {:then movie}
    <Movie {movie}/>
  {:catch error}
    <MovieError {error}/>
  {/await}
</div>

<style>
  .container {
    display:flex;
    justify-content: center;
  }
</style>
