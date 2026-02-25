<script lang="ts">
  const { children, class: className = "", ...rest } = $props();
</script>

<form class="form-field {className}" {...rest}>{@render children()}</form>
