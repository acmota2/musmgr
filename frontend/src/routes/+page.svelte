<script lang="ts">
import UnderConstruction from "$lib/ui/UnderConstruction.svelte";

let { composer } = $props();
</script>

{#if composer}
  <h1>Biography</h1>

  <p>{composer.biography}</p>
  <img src={`/composer/${composer.picture}`} alt="Portrait of {composer.fullName}" />
{:else}
  <UnderConstruction />
{/if}
