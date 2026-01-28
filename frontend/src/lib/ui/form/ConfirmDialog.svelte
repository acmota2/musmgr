<script lang="ts">
  import Button from "../Button.svelte";

  let {
    class: className = "",
    confirmDialog = $bindable<HTMLDialogElement>(),
    formId,
    message = "Apply changes?",
  } = $props();
</script>

<dialog class={`confirm-dialog ${className}`} bind:this={confirmDialog}>
  <p>{message}</p>

  <div class="confirm-buttons">
    <form method="dialog">
      <Button>Cancel</Button>
    </form>

    <Button type="submit" form={formId}>Yes</Button>
  </div>
</dialog>

<style>
  .confirm-dialog {
    & p {
      justify-self: center;
    }
    border: solid 1px rgba(17, 12, 46, 0.15);
    border-radius: var(--default-radius);
    box-shadow: rgba(17, 12, 46, 0.15) 0px 48px 100px 0px;
    padding: 15px 30px;
  }

  .confirm-buttons {
    display: flex;
    gap: 20px;
    justify-content: space-between;
    padding: var(--default-padding);
    border: solid var(--default-border) transparent;
  }
</style>
