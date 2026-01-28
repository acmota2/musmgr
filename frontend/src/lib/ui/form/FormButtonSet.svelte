<script lang="ts">
  import Button from "../Button.svelte";

  const {
    class: className = "",
    cancel,
    confirm,
    cancelMessage = "Cancel",
    confirmMessage = "Confirm",
  } = $props();
</script>

<div class={`actions ${className}`}>
  <Button type="button" onclick={cancel}>{cancelMessage}</Button>
  <Button type="button" onclick={confirm}>{confirmMessage}</Button>
</div>

<style>
  .actions {
    display: flex;
    justify-content: space-between;
    padding: var(--default-padding) 0;
    border: solid transparent var(--default-padding) 0;
    width: 100%;
  }
</style>
