<script>
  import apiService from '../apiService'
  import Movie from './Movie.svelte'
  import MovieError from './MovieError.svelte';
  import Loader from './Loader.svelte';

  export let imdbid;

  let search;

  $: imdbid, search = handleSearch(imdbid);

  async function handleSearch(movieId) {
    return await apiService.getMovie(movieId);
  }
</script>

<div>
  {#await search}
    <Loader/>
  {:then movie}
    <Movie {movie}/>
  {:catch error}
    <MovieError {error}/>
  {/await}
</div>
