<script lang="ts">
  import EditablePicture from "$lib/ui/editable/EditablePicture.svelte";
  import UnderConstruction from "$lib/ui/UnderConstruction.svelte";

  let { data } = $props();
  const { composer } = $derived(data);
</script>

{#if composer}
  <h1>Biography</h1>

  <div class="biography-content">
    <article class="biography-text">
      <p>{composer.biography}</p>
    </article>

    <div class="biography-picture">
      <EditablePicture
        wrapperClass="biography-picture"
        pictureId={composer.pictureId}
        src="/api/composer/picture"
        alt="Portrait of {composer.fullName}"
      />
    </div>
  </div>
{:else}
  <UnderConstruction />
{/if}

<style>
  .biography-content {
    max-width: var(--layout-default);
    display: flex;
    flex-direction: row;
    height: 100%;

    & .biography-text {
      flex: 8;
      & p {
        margin: 0;
      }
    }
  }

  .biography-picture {
    /* TODO: review */
    flex: 7;
    margin: 0;
  }
</style>
