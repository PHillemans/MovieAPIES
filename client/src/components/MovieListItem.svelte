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
    font-size: 1em;
    height:30px;
  }

  .subTitle {
    display: flex;
    font-size: 0.8em;
    justify-content: space-between;
    color: #7FC0FC;
  }

  .desc {
    font-size: 0.9em;
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
    margin: 40px 10px;
    background: #FAFAFA;
    box-shadow: inset 0px 0px 45px -27px rgba(0,0,0,0.25);
    width: 200px;
    border-radius: 20px;
  }
</style>
