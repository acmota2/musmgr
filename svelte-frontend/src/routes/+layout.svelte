<script lang="ts">
  import "../app.css";

  let { children } = $props();
</script>

<nav class="flex items-center bg-slate-400 px-10 py-5">
  <a href="/" class="text-lg font-bold">SongMGR</a>
</nav>

<div class="px-10 pt-4">
  {@render children()}
</div>

<footer class="flex justify-center pt-5">
  <p><b>SongMGR</b> 2025</p>
</footer>
