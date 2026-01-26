<script lang="ts">
  const { action = "", children, class: className = "" } = $props();
</script>

<form {action} class={`form-field ${className}`}>{@render children()}</form>

<style>
  .form-field {
    padding: var(--default-padding);
    border: var(--default-border) solid transparent;
  }
</style>
