<script lang="ts">
  import type { ChangeEventHandler } from "svelte/elements";
  import { IS_ADMIN } from "$lib/app";
  import Button from "../Button.svelte";
  import ConfirmDialog from "./ConfirmDialog.svelte";

  const onclick: ChangeEventHandler<HTMLButtonElement> = () => {
    confirmDialog?.showModal();
  };

  let { class: className = "", children, id, ...props } = $props();
  let confirmDialog: HTMLDialogElement | null = $state(null);
</script>

{#if IS_ADMIN}
  <form class="admin-form {className}" {id} {...props}>
    {@render children()}

    <div id="bottom-submit">
      <p><em>Changes were made</em></p>
      <div class="admin-form-button-set">
        <Button type="submit" {onclick}>Submit</Button>
        <Button type="reset">Reset</Button>
      </div>
    </div>
  </form>

  <ConfirmDialog bind:confirmDialog formId={id} />
{:else}
  {@render children()}
{/if}

<style>
  .admin-form {
    display: flex;
    min-height: 0;
    width: 100%;
  }

  #bottom-submit {
    display: flex;
    border: none;

    background: linear-gradient(
      to bottom,
      rgba(255, 255, 255, 0) 0%,
      rgba(255, 255, 255, 0.1) 10%,
      rgba(255, 255, 255, 0.6) 40%,
      rgba(255, 255, 255, 0.8) 80%,
      rgba(255, 255, 255, 1) 100%
    );

    padding: var(--default-padding);
    backdrop-filter: blur(15px);
    justify-content: center;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    gap: 20px;
  }

  #bottom-submit > * {
    display: flex;
    flex-direction: row;
    gap: 8px;
    align-items: center;
  }

  .admin-form-button-set {
    display: flex;
    justify-content: space-between;
    gap: 10px;
  }
</style>
