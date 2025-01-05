<svelte:options customElement="menu-tile" />

<script lang="ts">
  let {
    href,
    text,
    class: className = `
        xl:flex-row
        md:flex-column
        my-3
        sm:rounded-sm
        border-2
        border-black
        px-4
        pb-20
        pt-4
        text-xl
        font-bold`,
  } = $props();
</script>

<a {href} class="square-link">
  <div class={className}>
    <p class="square-text">{text}</p>
  </div>
</a>
