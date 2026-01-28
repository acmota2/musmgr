<script lang="ts">
  const { children, class: className = "", ...rest } = $props();
</script>

<form class={`form-field ${className}`} {...rest}>{@render children()}</form>

<style>
  .form-field {
    padding: var(--default-padding);
    border: var(--default-border) solid transparent;
  }
</style>
