<script>
  import Loader from './Loader.svelte';
  import MovieListItem from './MovieListItem.svelte';
  import {getMovies} from '../apiService';

  let movies = getMovieHandler();

  async function getMovieHandler() {
    return await getMovies();
  }
</script>


<div class="container">
  {#await movies}
    <Loader/>
  {:then movies}
    {#each movies as movie}
      <MovieListItem {movie}/>
    {/each}
  {:catch error}
    {console.log(error)}
  {/await}
</div>

<style>
  .container {
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
  }
</style>
