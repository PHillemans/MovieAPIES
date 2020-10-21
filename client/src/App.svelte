<script>
  import Form from './components/Form.svelte';
  import Result from './components/Result.svelte';
  import Movies from './components/Movies.svelte';

  let searching = false;
  let searchString;

  function handleSearch(e) {
    searchString = e.detail.id;
    searching = true;
  }

</script>

<main class="main">
  <div class="container">
    {#if searching}
      <Result imdbid={searchString}/>
    {/if}
    <Form on:search={handleSearch}/>
    <Movies/>
  </div>
</main>

<style>
  .main {
    margin: -8px;
    background-color: #FFFFD1;
    height: 100vh;
  }

  .container {
    padding: 20px;
    flex-direction: column;
    height: 130px;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>
