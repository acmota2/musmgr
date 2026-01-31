<script lang="ts">
  import { IS_ADMIN } from "$lib/app";
  import type { Composer } from "$lib/server/composer";
  import EditablePicture from "$lib/ui/editable/EditablePicture.svelte";
  import EditableTextArea from "$lib/ui/editable/EditableTextArea.svelte";
  import AdminForm from "$lib/ui/form/AdminForm.svelte";
  import UnderConstruction from "$lib/ui/UnderConstruction.svelte";

  let { data } = $props();
</script>

{#snippet page(composer: Composer)}
  <div class="biography-text">
    <article class="text-container">
      <h1 class="composer-name">{composer.fullName}</h1>
      <EditableTextArea id="biography" name="biography" text={composer.biography} />
    </article>
  </div>
  {#if IS_ADMIN || composer.pictureId !== null}
    <div class="biography-picture">
      <EditablePicture
        alt="Portrait of {composer.fullName}"
        name="picture"
        pictureId={composer.pictureId}
        src="/api/composer/picture"
      />
    </div>
  {/if}
{/snippet}

{#if data.composer}
  <div class="biography-content">
    {#if IS_ADMIN}
      <AdminForm
        id="biography-admin-form"
        method="POST"
        action="?/updateComposer"
        enctype="multipart/form-data"
      >
        {@render page(data.composer)}
      </AdminForm>
    {:else}
      {@render page(data.composer)}
    {/if}
  </div>
{:else}
  <UnderConstruction />
{/if}

<style>
  .composer-name {
    margin: 0;
    font-size: 100px;
  }

  .biography-content {
    display: flex;
    height: 100%;
  }

  .biography-text,
  .biography-picture {
    flex: 1;
    min-width: 0;
    min-height: 0;
  }

  .biography-picture {
    flex-shrink: 0;
    contain: layout paint size;
    aspect-ratio: auto;
  }

  .biography-text {
    flex-grow: 1;
  }

  :global(#biography) {
    overflow-y: auto;
  }

  .text-container {
    display: flex;
    flex-direction: column;
    padding: 125px;
    height: 100%;
    width: 100%;
  }

  :global(#biography) {
    margin: 100px 0;
  }
</style>
