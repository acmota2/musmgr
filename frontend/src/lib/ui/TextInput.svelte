<script lang="ts">
let {
  class: className = "",
  id = "",
  placeholder = "",
  required = false,
} = $props();
</script>

<input 
  class={`text-input ${className}`}
  {id} 
  {placeholder} 
  type="text" 
  {required} 
  min="" 
  max="" 
/>

<style>
  .text-input {
    border: solid 1px var(--border-color);
    border-radius: var(--default-radius);
    padding: var(--default-padding);
    width: 100%;
  }

  .text-input:focus {
    outline: none;
    border-color: var(--accent);
    box-shadow: 0 0 5px var(--accent);
  }
</style>
