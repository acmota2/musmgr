<script lang="ts">
  import "./button.css";

  let { class: className = "", href, children, ...props } = $props();
</script>

<a {href} class="button {className}" {...props}> {@render children()} </a>
