<script lang="ts">
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";
  import TextArea from "$lib/ui/form/TextArea.svelte";
  import TextInput from "$lib/ui/form/TextInput.svelte";

  function cancel() {
    history.back();
  }

  function confirm() {
    confirmDialog?.showModal();
  }

  const action = "?/create";
  let confirmDialog: HTMLDialogElement | null = $state(null);
</script>

<div class="container narrow">
  <FormField {action} method="POST" id="create-composer">
    <h1>About yourself</h1>
    <br>

    <label for="full_name">
      <h2>Full name:</h2>
    </label>
    <TextInput id="full_name" name="full_name" required />

    <label for="biography">
      <h2>Biography:</h2>
    </label>
    <TextArea id="biography" name="biography" required />
    <FormButtonSet {cancel} {confirm} />
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="create-composer" />
