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

<div class="form-container">
  <h1 class="form-title">
    Search for a movie by IMDBId:
  </h1>

  <form class="formpje" on:submit|preventDefault={emitSearch}>
    <input bind:value={imdbID}>
    <button type="submit" class="btn-search">Search</button>
  </form>
</div>

<style>
  input {
    border: 1px solid #12AADD;
  }
  .form-title {
    padding: 20px;
  }

  .btn-search {
    margin-left:-10px;
    border: #12AADD;
    background-color: #12AADD;
    height: 32px;
    color: white;
  }

  .form-container {
    flex-direction: column;
    height: 130px;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 40px;
  }
  </style>
