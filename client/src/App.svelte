<script>
  import Form from './components/Form.svelte';
  import Movie from './components/Movie.svelte';

  let searching = false;
  let searchString;

  function handleSearch(e) {
    searchString = e.detail.id;
    searching = true;
  }

</script>

<main>
  <div class="container">
    <Form on:search={handleSearch}/>

    {#if searching}
      <Movie imdbid={searchString}/>
    {/if}
  </div>
</main>

<style>
  .container {
    flex-direction: column;
    height: 130px;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>
