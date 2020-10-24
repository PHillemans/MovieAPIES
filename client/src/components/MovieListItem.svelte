<script>
  import {createEventDispatcher} from 'svelte';

  export let movie;

  let dispatch = createEventDispatcher();

  $: desc = truncate(movie.Description)
  $: nameStyle = movie.Name.length > 40;

  function truncate(st) {
    return st.substr(0,160) + '&hellip;'
  }

  function handleClick() {
    dispatch('select', {
      selected: movie
    });
  }
</script>

<div class="card" on:click={handleClick}>
  <span class="id">{movie.IMDBId}</span>
  <h1 class:smallerName={nameStyle}>{movie.Name}</h1>
  <p class="desc">{@html desc}</p>
  <p class="subTitle"><span>From: {movie.Year}</span> <span>Score: {movie.Score}</span></p>
</div>

<style>
  h1 {
    font-size: 1em;
    height:30px;
  }

  .smallerName {
    font-size: 0.90em;
  }

  .id {
    font-weight: lighter;
    font-size: 0.8em;
  }

  .subTitle {
    display: flex;
    font-size: 0.8em;
    justify-content: space-between;
    color: #7FC0FC;
  }

  .desc {
    font-size: 0.8em;
    color: #616c7d;
    height: 100px;
    text-overflow: ellipsis;
  }

  .card:hover {
    cursor: pointer;
  }

  .card {
    display:flex;
    flex-direction: column;
    align-content: space-between;
    padding: 20px;
    margin: 20px 10px;
    background-color: #FCFCFF;
    box-shadow: inset 0px 0px 45px -27px rgba(0,0,0,0.25);
    width: 180px;
    border-radius: 20px;
  }
</style>
