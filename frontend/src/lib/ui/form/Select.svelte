<script lang="ts">
  let { class: className = "", children, ...props } = $props();
</script>

<select class="select-elem will-focus {className}" {...props}>{@render children()}</select>

<style>
  .select-elem {
    border-radius: var(--default-radius);
    font-size: 20px;
    background: white;
    padding: var(--default-padding);
  }
</style>
