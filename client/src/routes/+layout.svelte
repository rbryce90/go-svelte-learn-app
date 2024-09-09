<script>
  import Nav from "$lib/Nav.svelte";
</script>

<Nav />

<main>
  <slot></slot>
</main>

<footer>this is a footer</footer>

<style>
  main {
    margin: 0; /* Remove margin around the main content */
    padding: 0; /* Remove padding around the main content */
    min-height: 100vh; /* Ensure main content takes up at least full viewport height */
  }
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }
</style>
