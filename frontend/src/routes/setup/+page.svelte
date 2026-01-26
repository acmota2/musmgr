<script>
    import FormField from "$lib/ui/FormField.svelte";
    import TextArea from "$lib/ui/TextArea.svelte";
    import TextInput from "$lib/ui/TextInput.svelte";
</script>


<div class="container narrow">
  <FormField>
    <h1>About yourself</h1>
    <br/>
    <label for="full_name">
      <h2>Full name:</h2>
    </label>
    <TextInput id="full_name" required />
    <label for="biography">
      <h2>Biography:</h2>
    </label>
    <TextArea id="biography" required />
  </FormField>
</div>
