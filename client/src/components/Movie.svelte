<script>
  import MovieError from './MovieError.svelte';
  export let movie;
</script>

<div class="movieContainer">
{#if !movie.error && movie}

  <div >
    <div class="movie">
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
</div>

<style>
  .movieContainer {
    display:flex;
    justify-content: space-between;
    flex-direction: row;
    max-width: 70%;
  }

  .movie {
    background-color: #EAECFF;
    box-shadow: 10px 10px 45px -27px rgba(0,0,0,0.25);
    padding: 20px;
    border-radius: 5px;
  }

  .description {
    margin-top: 2px;
    max-width:90%;
    display:flex;
    font-size: 0.85em;
    color: #666666;
    font-weight: bold;
    padding:5px;
  }

  .image{
    margin-left: 5%;
    height: 400px;
    border: 2px solid blue;
    border-radius: 20px;
    background-color: blue;
  }
</style>
