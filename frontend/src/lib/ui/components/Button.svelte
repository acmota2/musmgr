<script lang="ts">
  import "./button.css";

  const { class: className = "", children, ...rest } = $props();
</script>

<button class="button ${className}" {...rest}>{@render children()}</button>
