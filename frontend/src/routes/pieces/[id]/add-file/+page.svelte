<script lang="ts">
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FileInput from "$lib/ui/form/FileInput.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";
  import Select from "$lib/ui/form/Select.svelte";
  import TextInput from "$lib/ui/form/TextInput.svelte";
  import { capitalize } from "$lib/utils/index.js";

  function cancel() {
    history.back();
  }

  function confirm() {
    confirmDialog?.showModal();
  }

  let { data } = $props();
  let confirmDialog: HTMLDialogElement | null = $state(null);
  let filename = $state("");
</script>

<div class="container narrow">
  <h1>Adding file to {data.piece.title}</h1>
  <FormField
    class="create-file"
    method="POST"
    action="?/createFile"
    id="create-file"
    enctype="multipart/form-data"
  >
    <input type="hidden" name="piece_id" value="{data.piece.id}">
    <label for="name">
      <h2>File name</h2>
      <TextInput
        name="name"
        value={filename}
        placeholder="This can be filled by choosing a file"
        required
      />
    </label>

    <FileInput name="file" bind:filename />

    <label for="classification">
      <h2>Classification</h2>
      <Select name="classification">
        <option value="Public">Public</option>
        <option value="Admin">Admin</option>
      </Select>
    </label>

    <label for="file_type">
      <h2>File type</h2>
      <Select name="file_type">
        {#each data.fileTypes as fileType}
          <option value="{fileType}">{capitalize(fileType)}</option>
        {/each}
      </Select>
    </label>
    <FormButtonSet {cancel} {confirm}></FormButtonSet>
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="create-file" />

<style>
  :global(.create-file) {
    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  h1 {
    padding: 12px 0;
  }
  label {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }
</style>
