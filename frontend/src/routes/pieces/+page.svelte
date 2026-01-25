<script lang="ts">
  let { pieces } = $props();
</script>

{#each pieces as piece}
  <h2>{piece.title}</h2>
  <p>Year: {piece.composedAt}</p>
  <p>{piece.instrumentation}</p>
{/each}
