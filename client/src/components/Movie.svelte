<script>
  import MovieError from './MovieError.svelte';
  export let movie;
</script>

{#if !movie.error && movie}

  <div class="movie">
    <div>
      <p><strong>Name:</strong> {movie.Name}</p>
      <p><strong>Year:</strong> {movie.Year}</p>
      <p><strong>Score:</strong> {movie.Score}</p>
      <div><strong>Description:</strong>
        <p class="description">{movie.Description}</p>
      </div>
    </div>
  </div>

  {#if movie.Poster}
    <img class="image" src={movie.Poster} alt="moviePoster" />
  {/if}

{:else if movie.error}

  <MovieError error={movie.error}/>

{/if}

<style>
  .movie {
    background-color: #FCFFFC;
    box-shadow: 10px black;
    max-width: 70%;
    padding: 10px;
    border-radius: 5px;
  }

  .description {
    margin-top: 2px;
    max-width:90%;
    display:flex;
    text-align: center;
    font-size: 0.9em;
  }

  .image{
    margin-left: 5%;
    width: 30%;
    height: auto;
    border: 2px solid blue;
    border-radius: 20px;
    background-color: blue;
  }
</style>
