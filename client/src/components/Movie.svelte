<script>
  import apiService from '../apiService'
  import MovieError from './MovieError.svelte';
  import Loader from './Loader.svelte';

  export let imdbid;

  async function handleSearch(movieId) {
    search = await apiService.getMovie(movieId);
  }
</script>

<div>
    {#await handleSearch(imdbid)}
      <Loader/>
    {:then movie}
      {console.log(movie)}
      <div>somgn</div>
    {:catch error}
      <MovieError {error}/>
    {/await}
  </div>
