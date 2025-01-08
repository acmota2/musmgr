<svelte:options customElement="header-1" />

<script lang="ts">
  let { children } = $props();
</script>

<h1 class="pb-4 text-3xl font-bold">
  {@render children()}
</h1>
