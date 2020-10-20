<script>
  import {onDestroy} from 'svelte';

  let loader = "";
  
  let interval = setInterval(spin, 500)
  
  onDestroy(() => clearInterval(interval));
  
  function spin() {
    if (loader.length < 5) {
      loader += '.';
    } else {
      loader = "";
    }
  }
</script>

<div>
  loading{loader}
</div>
