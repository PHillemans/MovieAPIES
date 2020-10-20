<script>
  export let error;
</script>

<div>
  There was an error sending a request: <span class='error'> {error} </span>
</div>

<style>
  .error{
    color:red;
    font-weight: bold;
  }
</style>
