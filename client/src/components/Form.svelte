<script>
  import { createEventDispatcher } from 'svelte';

  let imdbID

  const dispatch = createEventDispatcher();
  function emitSearch() {
    dispatch('search', {
      id: imdbID
    })
  }
</script>

<style>
  .form-title {
    padding: 20px;
  }

  .btn-search {
    background-color: #1200dd;
    color: white;
  }

  .form-container {
    flex-direction: column;
    height: 130px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  </style>

  <div class="form-container">
    <div class="form-title">
      Search for a movie by IMDBId:
    </div>
    <form class="formpje" on:submit|preventDefault={emitSearch}>
      <input bind:value={imdbID}>
      <button type="submit" class="btn-search">Search</button>
    </form>
  </div>
