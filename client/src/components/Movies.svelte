<script>
  import Loader from './Loader.svelte';
  import MovieListItem from './MovieListItem.svelte';
  import Pagination from './Pagination.svelte';

  import {getMovies} from '../apiService';
  import {createEventDispatcher} from 'svelte';

  export let newResult;

  const dispatch = createEventDispatcher();

  let movies = getMovieHandler();
  
  let page = 1;

  //fetch on new movie
  $: newResult, movies = getMovieHandler(1);

  //fetch on new page
  $: page, movies = getMovieHandler(page);

  async function getMovieHandler(page) {
    return await getMovies(page);
  }

  function select(e) {
    dispatch("select", e.detail.selected)
  }
</script>

<h1 class="movieTitle">
  All movies
</h1>

<div class="container">
  {#await movies}
    <Loader/>
  {:then movies}
    {#each movies as movie}
      <MovieListItem on:select={select} {movie}/>
    {/each}
  {:catch error}
    {error}
  {/await}
</div>

<Pagination bind:page={page} />

<style>
  .movieTitle {
    width: 50vw;
    text-align:center;
    padding-top: 40px;
    border-top: 1px #CACACA solid;
    margin-top: 40px;
  }

  .container {
  display: flex;
  flex-wrap: wrap;
  }
</style>
