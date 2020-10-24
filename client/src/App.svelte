<script>
  import Form from './components/Form.svelte';
  import Result from './components/Result.svelte';
  import Movies from './components/Movies.svelte';
  import Movie from './components/Movie.svelte';

  let searchString;
  let newResult = false;  
  let movie = false;
  
  function handleSearch(e) {
    newResult = false;
    searchString = e.detail.id;
    movie = false;
  }

  function handleSelected(e) {
    newResult = false;
    movie = e.detail;
  }

</script>

<main class="main">
  <div class="container">
    <Form on:search={handleSearch}/>
    {#if movie}
      <Movie {movie}/>
    {:else}
      <Result bind:newResult={newResult} imdbid={searchString}/>
    {/if}
    <Movies on:select={handleSelected} bind:newResult={newResult}/>
  </div>
</main>

<style>
  .main {
    margin: -8px;
  }

  .container {
    height: 100vh;
    padding: 20px;
    flex-direction: column;
    height: 130px;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>
