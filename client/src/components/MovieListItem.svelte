<script>
  export let movie;
  $: desc = truncate(movie.Description)

  function truncate(st) {
    return st.substr(0,160) + '&hellip;'
  }
</script>

<div class="card">
  <h1>{movie.Name}</h1>
  <p class="desc">{@html desc}</p>
  <p class="subTitle"><span>From: {movie.Year}</span> <span>Score: {movie.Score}</span></p>
</div>

<style>
  h1 {
    font-size: 1.3em;
    height:30px;
  }

  .subTitle {
    display: flex;
    justify-content: space-between;
    color: #7FC0FC;
  }

  .desc {
    height: 100px;
    text-overflow: ellipsis;
    overflow:hidden;
  }

  .card:hover {
    cursor: pointer;
  }

  .card {
    display:flex;
    flex-direction: column;
    align-content: space-between;
    padding: 10px;
    margin: 10px;
    background: #FFF;
    width: 300px;
    border-radius: 20px;
  }
</style>
