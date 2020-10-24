<script>
  import {movieAmount} from '../apiService';

  export let page;

  let pageAmount = 10;

  let movies = movieAmount()
  let pages;
  movies.then((res) => {pages = Math.ceil(res/pageAmount)})
</script>

<div>
  {#each Array(pages) as _,pageNumber}
    <button 
      class="page"
      class:selected={page == pageNumber+1}
      on:click={() => page = pageNumber+1}
    >{pageNumber+1}</button>
  {/each}
</div>

  <style>
  .page {
    background-color: #99CCFF;
    border: 0px;
    cursor: pointer;
    color: white;
    width: 30px;
    margin: 2px;
  }

  .selected {
    background-color: #55AACC;
  }
  </style>
