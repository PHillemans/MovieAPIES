<script>
  import Movie from './components/Movie.svelte';
  import Loader from './components/Loader.svelte';
  import Form from './components/Form.svelte';
  import apiService from './apiService'

  let movie;
  let loading = false;

  async function handleSearch(e) {
    loading = true;
    movie = await apiService.getMovie(e.detail.id);
    loading = false;
  }
</script>

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

<main>
  <div class="container">
    <Form on:search={handleSearch}/>

    {#if loading }
      <Loader/>
    {:else if movie}
      <Movie bind:movie/>
    {/if}
  </div>
</main>
