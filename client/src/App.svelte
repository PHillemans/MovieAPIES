<script>
  import Form from './components/Form.svelte';
  import Movie from './components/Movie.svelte';

  let search = false;
  let searchString;

  function handleSearch(e) {
    searchString = e.detail.id;
    search = true;
  }

</script>

<main>
  <div class="container">
    <Form on:search={handleSearch}/>

    {#if search}
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
