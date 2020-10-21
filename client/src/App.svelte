<script>
  import Form from './components/Form.svelte';
  import Result from './components/Result.svelte';
  import Movies from './components/Movies.svelte';

  let searchString;

  function handleSearch(e) {
    searchString = e.detail.id;
    searching = true;
  }

</script>

<main class="main">
  <div class="container">
    <Form on:search={handleSearch}/>
    <Result imdbid={searchString}/>
    <Movies/>
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
