<svelte:options customElement="menu-tile" />

<script lang="ts">
  let {
    href,
    text,
    class: className = "\
      my-3 \
      sm:rounded-sm \
      border-2 \
      border-black \
      px-4 \
      pb-20 \
      pt-4 \
      text-xl \
      w-full \
      font-bold",
  } = $props();
</script>

<a {href} class="flex max-w-96 flex-row md:w-full">
  <div class={className}>
    <p class="square-text">{text}</p>
  </div>
</a>
