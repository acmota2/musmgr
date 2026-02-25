<script lang="ts">
  import "../components/button.css";

  let { class: className = "", filename = $bindable(), name, ...rest } = $props();

  function handleChange(e: Event) {
    const input = e.target as HTMLInputElement;
    filename = input.files?.[0]?.name ?? "";
  }
</script>

<label for="file" class="{className}">
  <input type="file" id="file" {name} {...rest} hidden onchange={handleChange}>
  <span class="button">Choose file</span>
  <span class="text">{filename || "No file selected"}</span>
</label>

<style>
  .text {
    font-size: var(--paragraph-font);
  }
</style>
