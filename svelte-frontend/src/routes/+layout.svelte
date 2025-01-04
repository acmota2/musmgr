<script lang="ts">
  import "../app.css";
  let { children } = $props();
</script>

<nav class="flex px-10 py-5">
  <a href="/" class="home">Home</a>
  <a href="/" class="addSong">+</a>
</nav>

<div class="body px-10">
  {@render children()}
</div>

<footer class="center pt-5">
  <p class="flex">SongMGR 2025</p>
</footer>
