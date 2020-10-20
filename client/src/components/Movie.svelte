<script>
  import MovieError from './MovieError.svelte';
  export let movie;
</script>

{#if !movie.error}

  <div>
    <p><strong>Name:</strong> {movie.Name}</p>
    <p><strong>Year:</strong> {movie.Year}</p>
    <p><strong>Score:</strong> {movie.Score}</p>
    <p><strong>Description:</strong> {movie.Description}</p>
  </div>

{:else}

  <MovieError error={movie.error}/>

{/if}
