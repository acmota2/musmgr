<script>
  import Button from "$lib/ui/Button.svelte";
  import TextArea from "$lib/ui/TextArea.svelte";
  import TextInput from "$lib/ui/TextInput.svelte";

  const composedAt = $state("");
  const description = $state("");
  const title = $state("");
</script>

<form action="POST">
  <label for="title">
    <h2>Title</h2>
  </label>
  <TextInput id="title" required />
  <label for="composed_at">
    <h2>Composed at</h2>
  </label>
  <TextInput id="composed_at" required placeholder="YYYY" />
  <label for="description">
    <h2>Description</h2>
  </label>
  <TextArea id="description" required />

  <Button>Submit</Button>
</form>
