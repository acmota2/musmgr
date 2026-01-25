<script lang="ts">
  let { composer } = $props();
</script>

{#if composer}
  <h1>Biography</h1>

  <p>{composer.biography}</p>
  <img src={`/composer/${composer.picture}`} alt="Portrait of {composer.fullName}" />
{:else}
  <h1>Composer Not Found</h1>
{/if}
