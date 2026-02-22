<script lang="ts">
  let { children, class: className = "", ...props } = $props();
</script>

<div class="bottom-container {className}" {...props}>{@render children()}</div>

<style>
  .bottom-container {
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
</style>
