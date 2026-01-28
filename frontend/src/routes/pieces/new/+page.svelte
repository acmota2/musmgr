<script lang="ts">
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";
  import TextArea from "$lib/ui/form/TextArea.svelte";
  import TextInput from "$lib/ui/form/TextInput.svelte";

  function confirm() {
    confirmDialog.showModal();
  }

  function cancel() {
    history.back();
  }

  let confirmDialog: HTMLDialogElement;
</script>

<div class="container narrow">
  <FormField method="POST" action="?/create" id="create-piece">
    <label for="title">
      <h2>Title</h2>
    </label>
    <TextInput name="title" required />
    <label for="composed_at">
      <h2>Composed at</h2>
    </label>
    <TextInput id="composed_at" required placeholder="YYYY" />
    <label for="description">
      <h2>Description</h2>
    </label>
    <TextArea name="description" required />

    <FormButtonSet {cancel} {confirm} />
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="create-piece" />
