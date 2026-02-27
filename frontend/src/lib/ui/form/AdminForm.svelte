<script lang="ts">
  import type { ChangeEventHandler } from "svelte/elements";
  import { IS_ADMIN } from "$lib/app";
  import BottomControlsContainer from "../components/BottomControlsContainer.svelte";
  import Button from "../components/Button.svelte";
  import ConfirmDialog from "./ConfirmDialog.svelte";

  const onclick: ChangeEventHandler<HTMLButtonElement> = () => {
    confirmDialog?.showModal();
  };

  const onreset: ChangeEventHandler<HTMLButtonElement> = () => {
    wasChanged = false;
  };

  const onchange: ChangeEventHandler<HTMLFormElement> = () => {
    wasChanged = true;
  };

  let { class: className = "", children, id, ...props } = $props();
  let confirmDialog: HTMLDialogElement | null = $state(null);
  let wasChanged = $state(false);
</script>

{#if IS_ADMIN}
  <form class="admin-form {className}" {onchange} {id} {...props}>
    {@render children()}

    {#if wasChanged}
      <BottomControlsContainer id="bottom-submit">
        <p><em>Changes were made</em></p>
        <div class="admin-form-button-set">
          <Button type="submit" {onclick}>Submit</Button>
          <Button type="reset" {onreset}>Clear content</Button>
        </div>
      </BottomControlsContainer>
    {/if}
  </form>

  <ConfirmDialog bind:confirmDialog formId={id} />
{:else}
  {@render children()}
{/if}

<style>
  .admin-form {
    position: relative;
    display: flex;
    min-height: 0;
    width: 100%;
  }

  :global(#bottom-submit > *) {
    display: flex;
    flex-direction: row;
    gap: 8px;
    align-items: center;
  }

  :global(#bottom-submit) {
    position: absolute;
    bottom: 0;
    left: 0;
  }

  .admin-form-button-set {
    display: flex;
    justify-content: space-between;
    gap: 10px;
  }
</style>
